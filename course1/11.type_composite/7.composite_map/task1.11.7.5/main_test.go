package main

import (
	"testing"
)

func Test_filterSentenceReplace(t *testing.T) {
	type args struct {
		sentence string
		filter   map[string]bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "#1",
			args: args{
				sentence: "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum",
				filter:   map[string]bool{"ipsum": true, "elit": true},
			},
			want: "Lorem dolor sit amet consectetur adipiscing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterSentenceReplace(tt.args.sentence, tt.args.filter); got != tt.want {
				t.Errorf("filterSentenceReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterSentence(t *testing.T) {
	type args struct {
		sentence string
		filter   map[string]bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "#1",
			args: args{
				sentence: "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum",
				filter:   map[string]bool{"ipsum": true, "elit": true},
			},
			want: "Lorem dolor sit amet consectetur adipiscing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterSentence(tt.args.sentence, tt.args.filter); got != tt.want {
				t.Errorf("filterSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
