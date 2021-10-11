package docxe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCommand_Extract(t *testing.T) {
	type fields struct {
		extract      bool
		update       bool
		overwriteDir bool
		removeDir    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				extract:      tt.fields.extract,
				update:       tt.fields.update,
				overwriteDir: tt.fields.overwriteDir,
				removeDir:    tt.fields.removeDir,
			}
			if got := c.Extract(); got != tt.want {
				t.Errorf("Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_OverwriteDir(t *testing.T) {
	type fields struct {
		extract      bool
		update       bool
		overwriteDir bool
		removeDir    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				extract:      tt.fields.extract,
				update:       tt.fields.update,
				overwriteDir: tt.fields.overwriteDir,
				removeDir:    tt.fields.removeDir,
			}
			if got := c.OverwriteDir(); got != tt.want {
				t.Errorf("OverwriteDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_RemoveDir(t *testing.T) {
	type fields struct {
		extract      bool
		update       bool
		overwriteDir bool
		removeDir    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				extract:      tt.fields.extract,
				update:       tt.fields.update,
				overwriteDir: tt.fields.overwriteDir,
				removeDir:    tt.fields.removeDir,
			}
			if got := c.RemoveDir(); got != tt.want {
				t.Errorf("RemoveDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_Update(t *testing.T) {
	type fields struct {
		extract      bool
		update       bool
		overwriteDir bool
		removeDir    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				extract:      tt.fields.extract,
				update:       tt.fields.update,
				overwriteDir: tt.fields.overwriteDir,
				removeDir:    tt.fields.removeDir,
			}
			if got := c.Update(); got != tt.want {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_normalize(t *testing.T) {
	type fields struct {
		extract      bool
		update       bool
		overwriteDir bool
		removeDir    bool
	}
	tests := []struct {
		name       string
		fields     fields
		wantFields fields
	}{
		{
			name: "Update priority",
			fields: fields{
				extract:      true,
				update:       true,
				overwriteDir: true,
				removeDir:    true,
			},
			wantFields: fields{
				extract:      false,
				update:       true,
				overwriteDir: false,
				removeDir:    true,
			},
		},
		{
			name: "Extract priority",
			fields: fields{
				extract:      true,
				update:       false,
				overwriteDir: true,
				removeDir:    true,
			},
			wantFields: fields{
				extract:      true,
				update:       false,
				overwriteDir: true,
				removeDir:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				extract:      tt.fields.extract,
				update:       tt.fields.update,
				overwriteDir: tt.fields.overwriteDir,
				removeDir:    tt.fields.removeDir,
			}
			c.normalize()
			r := fields{
				extract:      c.extract,
				update:       c.update,
				overwriteDir: c.overwriteDir,
				removeDir:    c.removeDir,
			}

			assert.Equal(t, tt.wantFields, r)
		})
	}
}

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

func TestNewCommand(t *testing.T) {
	type args struct {
		extract      bool
		update       bool
		overwriteDir bool
		removeDir    bool
	}
	tests := []struct {
		name string
		args args
		want Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommand(tt.args.extract, tt.args.update, tt.args.overwriteDir, tt.args.removeDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommand() = %v, want %v", got, tt.want)
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

func TestUnzipDirExistsError_Error(t *testing.T) {
	type fields struct {
		Err string
		Dir string
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
			e := &UnzipDirExistsError{
				Err: tt.fields.Err,
				Dir: tt.fields.Dir,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
