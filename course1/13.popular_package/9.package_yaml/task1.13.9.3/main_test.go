package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func Test_writeYAML(t *testing.T) {
	type args struct {
		filePath string
		data     []User
	}
	data := []User{
		{
			Name: gofakeit.Name(),
			Age:  gofakeit.Number(18, 60),
			Comments: []Comment{
				{
					Text: gofakeit.Generate("????? ??????? ?????????"),
				},
			},
		},
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr error
	}{
		{
			name: "#1",
			args: args{
				filePath: "./test/test.YAML",
				data:     data,
			},
			want: []byte(
				fmt.Sprintf(`- name: %s
  age: %d
  comments:
  - text: %s
`,
					data[0].Name,
					data[0].Age,
					data[0].Comments[0].Text,
				),
			),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeYAML(tt.args.filePath, tt.args.data); err != nil {
				t.Errorf("err = %s, want = %s", err.Error(), tt.wantErr)
			}

			file, err := os.OpenFile(tt.args.filePath, os.O_RDONLY, 0644)
			if err != nil {
				t.Errorf("failed to open test file: %s", err.Error())
			}
			defer os.RemoveAll(filepath.Dir(tt.args.filePath))

			data, err := io.ReadAll(file)
			if err != nil {
				t.Errorf("failed to read test file: %s", err.Error())
			}

			file.Close()

			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("file contains: %s,\nwant: %s", data, tt.want)
			}
		})
	}
}
