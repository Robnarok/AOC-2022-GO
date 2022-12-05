package main

import (
	"reflect"
	"testing"
)

func Test_runMove(t *testing.T) {
	type args struct {
		ship []string
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				ship: []string{"AB", "DEC"},
				from: 1,
				to:   0,
			},
			want: []string{"ABC", "DE"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runMove(tt.args.ship, tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{
			name: "readCommand1",
			args: args{
				command: "move 1 from 2 to 1",
			},
			want:  1,
			want1: 2,
			want2: 1,
		},
		{
			name: "readCommand2",
			args: args{
				command: "move 3 from 1 to 3",
			},
			want:  3,
			want1: 1,
			want2: 3,
		},
		{
			name: "readCommand3",
			args: args{
				command: "move 2 from 2 to 1",
			},
			want:  2,
			want1: 2,
			want2: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := readCommand(tt.args.command)
			if got != tt.want {
				t.Errorf("readCommand() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("readCommand() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("readCommand() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
