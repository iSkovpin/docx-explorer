package main

import (
	"errors"
	"fmt"
	"github.com/iSkovpin/docx-explorer/pkg/docxe"
	"log"
	"os"
)
import flag "github.com/spf13/pflag"

var (
	filename  string
	extract   bool
	update    bool
	remove    bool
	overwrite bool
)

func init() {
	flag.StringVarP(&filename, "file", "f", "", "Document  filename (*.docx)")
	flag.BoolVarP(&extract, "extract", "e", true, "Extract docx file")
	flag.BoolVarP(&update, "update", "u", false, "Update docx file and keep inner XML data")
	flag.BoolVarP(&remove, "remove", "r", false, "Remove extracted inner XML data")
	flag.BoolVarP(&overwrite, "overwrite", "o", false, "Overwrite already extracted inner XML data")
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	flag.Parse()
	if len(filename) == 0 && len(flag.Args()) == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(filename) == 0 {
		filename = flag.Arg(0)
	}

	var (
		err      error
		command  docxe.Command
		explorer docxe.Explorer
	)

	command = docxe.NewCommand(extract, update, overwrite, remove)
	explorer, err = docxe.NewExplorer(filename)
	if err != nil {
		return err
	}

	if command.Extract() {
		if err := explorer.Extract(command.OverwriteDir()); err != nil {
			dirExistsError := &docxe.UnzipDirExistsError{}
			if errors.As(err, &dirExistsError) {
				return errors.New(fmt.Sprintf("Directory %s already exists. Use --overwrite (-o) flag to overwrite\n", explorer.GetUnzipDir()))
			}
			return err
		}
		fmt.Printf("Inner XML data of '%s' extracted into '%s'\n", explorer.GetSrcFilename(), explorer.GetUnzipDir())
	}

	if command.Update() {
		if err := explorer.UpdateDocx(); err != nil {
			dirExistsError := &docxe.UnzipDirExistsError{}
			if errors.As(err, &dirExistsError) {
				return errors.New(fmt.Sprintf("Directory '%s' is not exist. There are no sources for updating\n", explorer.GetUnzipDir()))
			}
			return err
		}
		fmt.Printf("File '%s' updated from the source dirictory: '%s'\n", explorer.GetSrcFilename(), explorer.GetUnzipDir())
	}

	if command.RemoveDir() {
		if err := explorer.RemoveExtracted(); err != nil {
			return err
		}
		fmt.Printf("Source directory '%s' has been removed\n", explorer.GetUnzipDir())
	}

	return nil
}
