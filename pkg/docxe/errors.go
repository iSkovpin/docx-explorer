package docxe

import "fmt"

type UnzipDirExistsError struct {
	Err string
	Dir string
}

func (e *UnzipDirExistsError) Error() string {
	return fmt.Sprintf(e.Err, e.Dir)
}
