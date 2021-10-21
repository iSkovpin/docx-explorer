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
		{
			name: "Extract false",
			fields: fields{
				extract:      false,
				update:       false,
				overwriteDir: false,
				removeDir:    false,
			},
			want: false,
		},
		{
			name: "Extract true",
			fields: fields{
				extract:      true,
				update:       false,
				overwriteDir: false,
				removeDir:    false,
			},
			want: true,
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
		{
			name: "OverwriteDir false",
			fields: fields{
				extract:      false,
				update:       false,
				overwriteDir: false,
				removeDir:    false,
			},
			want: false,
		},
		{
			name: "OverwriteDir true",
			fields: fields{
				extract:      false,
				update:       false,
				overwriteDir: true,
				removeDir:    false,
			},
			want: true,
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
		{
			name: "RemoveDir false",
			fields: fields{
				extract:      false,
				update:       false,
				overwriteDir: false,
				removeDir:    false,
			},
			want: false,
		},
		{
			name: "RemoveDir true",
			fields: fields{
				extract:      false,
				update:       false,
				overwriteDir: false,
				removeDir:    true,
			},
			want: true,
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
		{
			name: "Update false",
			fields: fields{
				extract:      false,
				update:       false,
				overwriteDir: false,
				removeDir:    false,
			},
			want: false,
		},
		{
			name: "Update true",
			fields: fields{
				extract:      false,
				update:       true,
				overwriteDir: false,
				removeDir:    false,
			},
			want: true,
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
		{
			name: "No changes 1",
			args: args{
				extract:      true,
				update:       false,
				overwriteDir: true,
				removeDir:    false,
			},
			want: Command{
				extract:      true,
				update:       false,
				overwriteDir: true,
				removeDir:    false,
			},
		},
		{
			name: "No changes 2",
			args: args{
				extract:      false,
				update:       true,
				overwriteDir: false,
				removeDir:    false,
			},
			want: Command{
				extract:      false,
				update:       true,
				overwriteDir: false,
				removeDir:    false,
			},
		},
		{
			name: "Update priority",
			args: args{
				extract:      true,
				update:       true,
				overwriteDir: true,
				removeDir:    true,
			},
			want: Command{
				extract:      false,
				update:       true,
				overwriteDir: false,
				removeDir:    true,
			},
		},
		{
			name: "Update priority",
			args: args{
				extract:      true,
				update:       false,
				overwriteDir: true,
				removeDir:    true,
			},
			want: Command{
				extract:      true,
				update:       false,
				overwriteDir: true,
				removeDir:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommand(tt.args.extract, tt.args.update, tt.args.overwriteDir, tt.args.removeDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
