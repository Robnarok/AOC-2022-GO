package main

import (
	"reflect"
	"testing"
)

func Test_parseRange(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "parse1-5",
			args: args{
				input: "1-5",
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "parse2-6",
			args: args{
				input: "2-6",
			},
			want: []int{2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseRange(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkSubspace(t *testing.T) {
	type args struct {
		elve1 []int
		elve2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case1",
			args: args{
				elve1: []int{2, 3, 4},
				elve2: []int{1, 2, 3, 4, 5, 6},
			},
			want: true,
		},
		{
			name: "Case2",
			args: args{
				elve1: []int{1, 2, 3, 4, 5, 6},
				elve2: []int{2, 3, 4},
			},
			want: true,
		},
		{
			name: "Case3",
			args: args{
				elve1: []int{1, 2, 3, 4, 5, 6},
				elve2: []int{5, 6, 7, 8, 9},
			},
			want: false,
		},
		{
			name: "Case4",
			args: args{
				elve1: []int{4, 5, 6, 7, 8},
				elve2: []int{1, 2, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubspace(tt.args.elve1, tt.args.elve2); got != tt.want {
				t.Errorf("checkSubspace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_challangeOne(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2-8,3-7",
			args: args{"2-8,3-7"},
			want: true,
		},
		{
			name: "6-6,4-6",
			args: args{"6-6,4-6"},
			want: true,
		},
		{
			name: "2-4,6-8",
			args: args{"2-4,6-8"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challangeOne(tt.args.input); got != tt.want {
				t.Errorf("challangeOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
