package main

import (
	"reflect"
	"testing"
)

func Test_unmarshal(t *testing.T) {
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    interface{}
	}{
		{
			name: "JSON",
			args: args{
				data: []byte(`{"name":"John","age":30}`),
				v:    &Person{},
			},
			wantErr: false,
			want: Person{
				Name: "John",
				Age:  30,
			},
		},
		{
			name: "YAML",
			args: args{
				data: []byte("name: John\nage: 30\n"),
				v:    &Person{},
			},
			wantErr: false,
			want: Person{
				Name: "John",
				Age:  30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := unmarshal(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.want, *tt.args.v.(*Person)) {
				t.Errorf("Person: %v, want: %v", *tt.args.v.(*Person), tt.want)
			}
		})
	}
}
