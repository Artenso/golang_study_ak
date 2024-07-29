package main

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func Test_getJSON(t *testing.T) {
	type args struct {
		data []User
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
		want    string
		wantErr bool
	}{
		{
			name: "#1",
			args: args{data: data},
			want: fmt.Sprintf(
				"[{\"name\":\"%s\",\"age\":%d,\"comments\":[{\"text\":\"%s\"}]}]",
				data[0].Name,
				data[0].Age,
				data[0].Comments[0].Text,
			),
		},
		{
			name: "#2",
			args: args{},
			want: "null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("getJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
