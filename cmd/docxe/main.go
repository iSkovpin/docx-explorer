package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)
import flag "github.com/spf13/pflag"

var (
	filename   string
	fExtract   bool
	fUpdate    bool
	fRemove    bool
	fOverwrite bool
)

func init() {
	flag.StringVarP(&filename, "file", "f", "", "Document  filename (*.docx)")
	flag.BoolVarP(&fExtract, "extract", "e", true, "Extract docx file")
	flag.BoolVarP(&fUpdate, "update", "u", false, "Update docx file and keep inner XML data")
	flag.BoolVarP(&fRemove, "remove", "r", false, "Remove extracted inner XML data")
	flag.BoolVarP(&fOverwrite, "overwrite", "o", false, "Overwrite already extracted inner XML data")
}

func main() {
	flag.Parse()
	if len(filename) == 0 && len(flag.Args()) == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(filename) == 0 {
		filename = flag.Arg(0)
	}

	fInfo, err := getDocxFileInfo(filename)
	if err != nil {
		panic(err.Error())
	}

	var unzipDirExists, docxFileExists bool
	unzipDirExists, err = fInfo.UnzipDirExists()
	if err != nil {
		log.Fatal(err.Error())
	}
	docxFileExists, err = fInfo.DocxFileExists()
	if err != nil {
		log.Fatal(err.Error())
	}

	// The update operation has more priority
	if fUpdate {
		fExtract = false
		fOverwrite = false
	}

	if fExtract {
		fUpdate = false
		fRemove = false
	}

	if fExtract {
		if unzipDirExists && !fOverwrite {
			log.Fatalf("Directory %s already exists. Use --overwrite (-o) flag to overwrite\n", fInfo.unzipDir)
		} else if unzipDirExists && fOverwrite {
			err := os.RemoveAll(fInfo.unzipDir)
			if err != nil {
				log.Fatal(err.Error())
			}
		}

		if err := Unzip(fInfo.docxPath, fInfo.unzipDir); err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("Inner XML data of '%s' extracted into '%s'\n", fInfo.srcFilename, fInfo.unzipDir)
	}

	if fUpdate {
		if !unzipDirExists {
			log.Fatalf("Directory '%s' is not exist. There are no sources for updating\n", fInfo.unzipDir)
		}

		if err := RecursiveZip(fInfo.unzipDir, fInfo.tempDocxPath); err != nil {
			log.Fatal(err.Error())
		}
		defer func() {
			err := os.RemoveAll(fInfo.tempDocxPath)
			if err != nil {
				log.Fatal(err.Error())
			}
		}()

		if docxFileExists {
			err := os.Rename(fInfo.docxPath, fInfo.docxBkpPath)
			if err != nil {
				log.Fatal(err.Error())
			}
			err = os.Rename(fInfo.tempDocxPath, fInfo.docxPath)
			if err != nil {
				log.Fatal(err.Error())
			}

			err = os.RemoveAll(fInfo.docxBkpPath)
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			err = os.Rename(fInfo.tempDocxPath, fInfo.docxPath)
			if err != nil {
				log.Fatal(err.Error())
			}
		}

		fmt.Printf("File '%s' updated from the source dirictory: '%s'\n", fInfo.srcFilename, fInfo.unzipDir)
	}

	if fRemove {
		if err := os.RemoveAll(fInfo.unzipDir); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Source directory '%s' has been removed\n", fInfo.unzipDir)
	}
}

type DocxFileInfo struct {
	srcFilename  string
	baseDir      string
	docxPath     string
	docxBkpPath  string
	unzipDir     string
	tempDocxPath string
}

func (f *DocxFileInfo) UnzipDirExists() (bool, error) {
	return f.exists(f.unzipDir)
}

func (f *DocxFileInfo) DocxFileExists() (bool, error) {
	return f.exists(f.docxPath)
}

func (f *DocxFileInfo) exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func getDocxFileInfo(filename string) (DocxFileInfo, error) {
	var err error
	var fInfo = DocxFileInfo{srcFilename: filename}

	fileExt := filepath.Ext(filename)
	if fileExt != ".docx" {
		err = errors.New(fmt.Sprintf("wrong file extension: '.docx' expected, got '%s'", fileExt))
		return fInfo, err
	}

	fInfo.docxPath, err = filepath.Abs(filename)
	if err != nil {
		return fInfo, err
	}

	fInfo.baseDir = filepath.Dir(fInfo.docxPath)
	fInfo.unzipDir = fInfo.docxPath[:len(fInfo.docxPath)-len(fileExt)]
	fInfo.tempDocxPath = fInfo.baseDir + "/" + RandomString(20) + ".docx"
	fInfo.docxBkpPath = fInfo.unzipDir + "_bkp_" + RandomString(20) + ".docx"

	return fInfo, err
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				return err
			}

			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func RecursiveZip(pathToZip, destinationPath string) error {
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	myZip := zip.NewWriter(destinationFile)
	err = filepath.WalkDir(pathToZip, func(filePath string, d os.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		relPath := strings.TrimPrefix(filePath, pathToZip+"/")
		zipFile, err := myZip.Create(relPath)
		if err != nil {
			return err
		}
		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = myZip.Close()
	if err != nil {
		return err
	}
	return nil
}

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
