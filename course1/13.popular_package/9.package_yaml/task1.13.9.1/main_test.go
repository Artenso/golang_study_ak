package main

import "testing"

func Test_getYAML(t *testing.T) {
	type args struct {
		cfgs []Config
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "#1",
			args: args{
				cfgs: []Config{
					{
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
					{
						Server: Server{
							Port: "8080",
						},
						Db: Db{
							Host:     "localhost",
							Port:     "5432",
							User:     "user",
							Password: "password456",
						},
					},
				},
			},
			want: `- server:
    port: "8080"
  db:
    host: localhost
    port: "5432"
    user: admin
    password: password123
- server:
    port: "8080"
  db:
    host: localhost
    port: "5432"
    user: user
    password: password456
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getYAML(tt.args.cfgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("getYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}
