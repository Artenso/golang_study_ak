package main

import (
	"reflect"
	"testing"
	"time"
)

func TestHashMap_SetGet(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name      string
		hm        HashMaper
		args      args
		wantVal   any
		wantCheck bool
	}{
		{
			name: "crc64",
			hm:   NewHashMap(16, WithHashCRC64()),
			args: args{
				key:   "test",
				value: "test value",
			},
			wantVal:   "test value",
			wantCheck: true,
		},
		{
			name: "crc32",
			hm:   NewHashMap(16, WithHashCRC32()),
			args: args{
				key:   "test",
				value: "test value",
			},
			wantVal:   "test value",
			wantCheck: true,
		},
		{
			name: "crc16",
			hm:   NewHashMap(16, WithHashCRC16()),
			args: args{
				key:   "test",
				value: "test value",
			},
			wantVal:   "test value",
			wantCheck: true,
		},
		{
			name: "crc8",
			hm:   NewHashMap(16, WithHashCRC8()),
			args: args{
				key:   "test",
				value: "test value",
			},
			wantVal:   "test value",
			wantCheck: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hm.Set(tt.args.key, tt.args.value)
			value, check := tt.hm.Get(tt.args.key)
			if !reflect.DeepEqual(value, tt.wantVal) {
				t.Errorf("wrong value: %v, want %v", value, tt.wantVal)
			}
			if !reflect.DeepEqual(check, tt.wantCheck) {
				t.Errorf("wrong check: %v, want %v", check, tt.wantCheck)
			}
		})
	}
}

func TestMeassureTime(t *testing.T) {
	type args struct {
		f func()
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "#1",
			args: args{func() { time.Sleep(time.Second) }},
			want: time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MeassureTime(tt.args.f); got < tt.want {
				t.Errorf("MeassureTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
