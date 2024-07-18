package main

import (
	"reflect"
	"testing"
)

func Test_getConfigFromYAML(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "#1",
			args: args{
				data: []byte(`server:
  port: "8080"
db:
  host: localhost
  port: "5432"
  user: admin
  password: password123
`,
				),
			},
			want: Config{
				Server: Server{
					Port: "8080",
				},
				Db: Db{
					Host:     "localhost",
					Port:     "5432",
					User:     "admin",
					Password: "password123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConfigFromYAML(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConfigFromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfigFromYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}
