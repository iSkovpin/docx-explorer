package docxe

import "testing"

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
		{
			name: "Error message",
			fields: fields{
				Err: "Directory %s already exists",
				Dir: "/path/to/dir",
			},
			want: "Directory /path/to/dir already exists",
		},
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
