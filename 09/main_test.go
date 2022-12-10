package main

import (
	"reflect"
	"testing"
)

func Test_checkDetached(t *testing.T) {
	type args struct {
		head position
		tail position
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CheckXOffset",
			args: args{
				head: position{0, 0},
				tail: position{-2, 0},
			},
			want: true,
		},
		{
			name: "CheckXOffsetDiagonal",
			args: args{
				head: position{0, 1},
				tail: position{-2, 0},
			},
			want: true,
		},
		{
			name: "CheckYOffset",
			args: args{
				head: position{0, 2},
				tail: position{0, 0},
			},
			want: true,
		},
		{
			name: "CheckYOffset2",
			args: args{
				head: position{0, 2},
				tail: position{0, 4},
			},
			want: true,
		},
		{
			name: "CheckNoOffset",
			args: args{
				head: position{0, 2},
				tail: position{0, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDetached(tt.args.head, tt.args.tail); got != tt.want {
				t.Errorf("checkDetached() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_move(t *testing.T) {
	type args struct {
		head    position
		tail    position
		command string
	}
	tests := []struct {
		name  string
		args  args
		want  position
		want1 position
	}{
		{
			name: "Run1",
			args: args{
				head:    position{0, 1},
				tail:    position{0, 0},
				command: "U",
			},
			want:  position{0, 2},
			want1: position{0, 1},
		},
		{
			name: "Run2",
			args: args{
				head:    position{0, 1},
				tail:    position{0, 0},
				command: "R",
			},
			want:  position{1, 1},
			want1: position{0, 0},
		},
		{
			name: "Run3",
			args: args{
				head:    position{1, 1},
				tail:    position{0, 0},
				command: "R",
			},
			want:  position{2, 1},
			want1: position{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := move(tt.args.head, tt.args.tail, tt.args.command)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("move() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("move() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
