package main

import (
	"testing"
)

func Test_parseItems(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name:  "Case1",
			args:  args{"vJrwpWtwJgWrhcsFMMfFFhFp"},
			want:  "vJrwpWtwJgWr",
			want1: "hcsFMMfFFhFp",
		},
		{
			name:  "Case2",
			args:  args{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
			want:  "jqHRNqRjqzjGDLGL",
			want1: "rsFMfFZSrLrFZsSL",
		},
		{
			name:  "Case3",
			args:  args{"PmmdzqPrVvPwwTWBwg"},
			want:  "PmmdzqPrV",
			want1: "vPwwTWBwg",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseItems(tt.args.input)
			if got != tt.want {
				t.Errorf("parseItems() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseItems() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getCommonItem(t *testing.T) {
	type args struct {
		input1 string
		input2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case1",
			args: args{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
			want: "p",
		},
		{
			name: "Case2",
			args: args{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
			want: "L",
		},
		{
			name: "Case3",
			args: args{"PmmdzqPrV", "vPwwTWBwg"},
			want: "P",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommonItem(tt.args.input1, tt.args.input2); got != tt.want {
				t.Errorf("getCommonItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValue(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Casea",
			args: args{"a"},
			want: 1,
		},
		{
			name: "Casez",
			args: args{"z"},
			want: 26,
		},
		{
			name: "CaseA",
			args: args{"A"},
			want: 27,
		},
		{
			name: "Casea",
			args: args{"Z"},
			want: 52,
		},
		{
			name: "Cases",
			args: args{"s"},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValue(tt.args.input); got != tt.want {
				t.Errorf("getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
