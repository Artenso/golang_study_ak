package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id   int
		opts []UserOption
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "#1",
			args: args{
				id: 1,
				opts: []UserOption{
					WithUsername("testuser"),
					WithEmail("testuser@example.com"),
					WithRole("admin"),
				},
			},
			want: &User{
				ID:       1,
				Username: "testuser",
				Email:    "testuser@example.com",
				Role:     "admin",
			},
		},
		{
			name: "#2",
			args: args{
				id: 1,
				opts: []UserOption{
					WithUsername("testuser"),
					WithRole("admin"),
				},
			},
			want: &User{
				ID:       1,
				Username: "testuser",
				Role:     "admin",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.id, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	oldOut := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("failed create pipe: %s", err)
	}

	os.Stdout = w

	main()
	w.Close()

	os.Stdout = oldOut

	var out bytes.Buffer
	out.ReadFrom(r)
	expected := "User: &{ID:1 Username:testuser Email:testuser@example.com Role:admin}\n"
	if out.String() != expected {
		t.Errorf("want: %s\ngot: %s", expected, out.String())
	}
}
