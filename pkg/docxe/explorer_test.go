package docxe

import (
	"reflect"
	"testing"
)

func TestExplorer_DocxFileExists(t *testing.T) {
	type fields struct {
		srcFilename  string
		baseDir      string
		docxPath     string
		docxBkpPath  string
		unzipDir     string
		tempDocxPath string
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Explorer{
				srcFilename:  tt.fields.srcFilename,
				baseDir:      tt.fields.baseDir,
				docxPath:     tt.fields.docxPath,
				docxBkpPath:  tt.fields.docxBkpPath,
				unzipDir:     tt.fields.unzipDir,
				tempDocxPath: tt.fields.tempDocxPath,
			}
			got, err := e.DocxFileExists()
			if (err != nil) != tt.wantErr {
				t.Errorf("DocxFileExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DocxFileExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExplorer_Extract(t *testing.T) {
	type fields struct {
		srcFilename  string
		baseDir      string
		docxPath     string
		docxBkpPath  string
		unzipDir     string
		tempDocxPath string
	}
	type args struct {
		overwrite bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Explorer{
				srcFilename:  tt.fields.srcFilename,
				baseDir:      tt.fields.baseDir,
				docxPath:     tt.fields.docxPath,
				docxBkpPath:  tt.fields.docxBkpPath,
				unzipDir:     tt.fields.unzipDir,
				tempDocxPath: tt.fields.tempDocxPath,
			}
			if err := e.Extract(tt.args.overwrite); (err != nil) != tt.wantErr {
				t.Errorf("Extract() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExplorer_GetSrcFilename(t *testing.T) {
	type fields struct {
		srcFilename  string
		baseDir      string
		docxPath     string
		docxBkpPath  string
		unzipDir     string
		tempDocxPath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Explorer{
				srcFilename:  tt.fields.srcFilename,
				baseDir:      tt.fields.baseDir,
				docxPath:     tt.fields.docxPath,
				docxBkpPath:  tt.fields.docxBkpPath,
				unzipDir:     tt.fields.unzipDir,
				tempDocxPath: tt.fields.tempDocxPath,
			}
			if got := e.GetSrcFilename(); got != tt.want {
				t.Errorf("GetSrcFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExplorer_GetUnzipDir(t *testing.T) {
	type fields struct {
		srcFilename  string
		baseDir      string
		docxPath     string
		docxBkpPath  string
		unzipDir     string
		tempDocxPath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Explorer{
				srcFilename:  tt.fields.srcFilename,
				baseDir:      tt.fields.baseDir,
				docxPath:     tt.fields.docxPath,
				docxBkpPath:  tt.fields.docxBkpPath,
				unzipDir:     tt.fields.unzipDir,
				tempDocxPath: tt.fields.tempDocxPath,
			}
			if got := e.GetUnzipDir(); got != tt.want {
				t.Errorf("GetUnzipDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExplorer_RemoveExtracted(t *testing.T) {
	type fields struct {
		srcFilename  string
		baseDir      string
		docxPath     string
		docxBkpPath  string
		unzipDir     string
		tempDocxPath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Explorer{
				srcFilename:  tt.fields.srcFilename,
				baseDir:      tt.fields.baseDir,
				docxPath:     tt.fields.docxPath,
				docxBkpPath:  tt.fields.docxBkpPath,
				unzipDir:     tt.fields.unzipDir,
				tempDocxPath: tt.fields.tempDocxPath,
			}
			if err := e.RemoveExtracted(); (err != nil) != tt.wantErr {
				t.Errorf("RemoveExtracted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExplorer_RunCommand(t *testing.T) {
	type fields struct {
		srcFilename  string
		baseDir      string
		docxPath     string
		docxBkpPath  string
		unzipDir     string
		tempDocxPath string
	}
	type args struct {
		c Command
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Explorer{
				srcFilename:  tt.fields.srcFilename,
				baseDir:      tt.fields.baseDir,
				docxPath:     tt.fields.docxPath,
				docxBkpPath:  tt.fields.docxBkpPath,
				unzipDir:     tt.fields.unzipDir,
				tempDocxPath: tt.fields.tempDocxPath,
			}
			if err := e.RunCommand(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("RunCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExplorer_UnzipDirExists(t *testing.T) {
	type fields struct {
		srcFilename  string
		baseDir      string
		docxPath     string
		docxBkpPath  string
		unzipDir     string
		tempDocxPath string
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Explorer{
				srcFilename:  tt.fields.srcFilename,
				baseDir:      tt.fields.baseDir,
				docxPath:     tt.fields.docxPath,
				docxBkpPath:  tt.fields.docxBkpPath,
				unzipDir:     tt.fields.unzipDir,
				tempDocxPath: tt.fields.tempDocxPath,
			}
			got, err := e.UnzipDirExists()
			if (err != nil) != tt.wantErr {
				t.Errorf("UnzipDirExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnzipDirExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExplorer_UpdateDocx(t *testing.T) {
	type fields struct {
		srcFilename  string
		baseDir      string
		docxPath     string
		docxBkpPath  string
		unzipDir     string
		tempDocxPath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Explorer{
				srcFilename:  tt.fields.srcFilename,
				baseDir:      tt.fields.baseDir,
				docxPath:     tt.fields.docxPath,
				docxBkpPath:  tt.fields.docxBkpPath,
				unzipDir:     tt.fields.unzipDir,
				tempDocxPath: tt.fields.tempDocxPath,
			}
			if err := e.UpdateDocx(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDocx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewExplorer(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Explorer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewExplorer(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewExplorer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExplorer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
