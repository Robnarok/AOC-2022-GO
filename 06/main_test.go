package main

import (
	"testing"
)

func Test_collectSequenceFromPos(t *testing.T) {
	type args struct {
		input string
		place int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ABCDEF1",
			args: args{
				input: "ABCDEF",
				place: 3,
			},
			want: "ABCD",
		},
		{
			name: "ABCDEF2",
			args: args{
				input: "ABCDEF",
				place: 5,
			},
			want: "CDEF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collectSequenceFromPos(tt.args.input, tt.args.place); got != tt.want {
				t.Errorf("collectSequenceFromPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSignal(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isNoSignal1",
			args: args{
				input: "AABC",
			},
			want: false,
		},
		{
			name: "isNoSignal2",
			args: args{
				input: "ABAC",
			},
			want: false,
		},
		{
			name: "isNoSignal3",
			args: args{
				input: "ABAA",
			},
			want: false,
		},
		{
			name: "isSignal1",
			args: args{
				input: "ABCD",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSignal(tt.args.input); got != tt.want {
				t.Errorf("isSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchSequence(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{
				input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			want: 7,
		},
		{
			name: "Case2",
			args: args{
				input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 5,
		},
		{
			name: "Case5",
			args: args{
				input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchSequence(tt.args.input); got != tt.want {
				t.Errorf("searchSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
