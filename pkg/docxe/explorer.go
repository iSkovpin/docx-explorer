package docxe

import (
	"fmt"
	"github.com/iSkovpin/docx-explorer/pkg/compress"
	"github.com/iSkovpin/docx-explorer/pkg/file"
	"github.com/iSkovpin/docx-explorer/pkg/random"
	"log"
	"os"
	"path/filepath"
)

type Explorer struct {
	srcFilename  string
	baseDir      string
	docxPath     string
	docxBkpPath  string
	unzipDir     string
	tempDocxPath string
}

func (e *Explorer) RunCommand(c Command) error {
	if c.Extract() {
		if err := e.Extract(c.OverwriteDir()); err != nil {
			return err
		}
	}

	if c.Update() {
		if err := e.UpdateDocx(); err != nil {
			return err
		}
	}

	if c.RemoveDir() {
		if err := e.RemoveExtracted(); err != nil {
			return err
		}
	}

	return nil
}

func (e *Explorer) Extract(overwrite bool) error {
	unzipDirExists, err := e.UnzipDirExists()
	if err != nil {
		return err
	}

	if unzipDirExists && !overwrite {
		return &UnzipDirExistsError{Err: "Directory %s already exists. Use overwrite=true to overwrite", Dir: e.unzipDir}
	} else if unzipDirExists && overwrite {
		if err := e.RemoveExtracted(); err != nil {
			return err
		}
	}

	if err := compress.Unzip(e.docxPath, e.unzipDir); err != nil {
		return err
	}
	return nil
}

func (e *Explorer) UpdateDocx() error {
	var unzipDirExists, docxFileExists bool
	var err error

	if unzipDirExists, err = e.UnzipDirExists(); err != nil {
		return err
	}
	if docxFileExists, err = e.DocxFileExists(); err != nil {
		return err
	}

	if !unzipDirExists {
		return &UnzipDirExistsError{Err: "Directory '%s' is not exist. There are no sources for updating", Dir: e.unzipDir}
	}

	if err := compress.RecursiveZip(e.unzipDir, e.tempDocxPath); err != nil {
		return err
	}
	defer func() {
		err := os.RemoveAll(e.tempDocxPath)
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	if docxFileExists {
		if err := os.Rename(e.docxPath, e.docxBkpPath); err != nil {
			return err
		}
		if err := os.Rename(e.tempDocxPath, e.docxPath); err != nil {
			return err
		}
		if err := os.RemoveAll(e.docxBkpPath); err != nil {
			return err
		}
	} else if err := os.Rename(e.tempDocxPath, e.docxPath); err != nil {
		return err
	}

	return nil
}

func (e *Explorer) RemoveExtracted() error {
	err := os.RemoveAll(e.unzipDir)
	if err != nil {
		return err
	}
	return nil
}

func (e *Explorer) UnzipDirExists() (bool, error) {
	return file.Exists(e.unzipDir)
}

func (e *Explorer) DocxFileExists() (bool, error) {
	return file.Exists(e.docxPath)
}

func NewExplorer(filename string) (Explorer, error) {
	var err error
	var d = Explorer{srcFilename: filename}

	fileExt := filepath.Ext(filename)
	if fileExt != ".docx" {
		err = fmt.Errorf("wrong file extension: '.docx' expected, got '%s'", fileExt)
		return d, err
	}

	d.docxPath, err = filepath.Abs(filename)
	if err != nil {
		return d, err
	}

	d.baseDir = filepath.Dir(d.docxPath)
	d.unzipDir = d.docxPath[:len(d.docxPath)-len(fileExt)]
	d.tempDocxPath = d.baseDir + string(os.PathSeparator) + random.String(20) + ".docx"
	d.docxBkpPath = d.unzipDir + "_bkp_" + random.String(20) + ".docx"

	return d, err
}

func (e *Explorer) GetUnzipDir() string {
	return e.unzipDir
}

func (e *Explorer) GetSrcFilename() string {
	return e.srcFilename
}
