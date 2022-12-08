package main

import (
	"testing"
)

func Test_checkVisibilityRight(t *testing.T) {
	type args struct {
		forest [][]int
		posx   int
		posy   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				forest: [][]int{
					{0, 0, 1, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 0,
				posx: 2,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				forest: [][]int{
					{0, 0, 0, 1},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 0,
				posx: 3,
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				forest: [][]int{
					{0, 0, 0, 1},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 1,
				posx: 3,
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				forest: [][]int{
					{0, 0, 0, 1},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 0,
				posx: 2,
			},
			want: false,
		},
		{
			name: "case5",
			args: args{
				forest: [][]int{
					{0, 1, 0, 2},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 0,
				posx: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkVisibilityRight(tt.args.forest, tt.args.posx, tt.args.posy); got != tt.want {
				t.Errorf("checkVisibilityRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkVisibilitLeft(t *testing.T) {
	type args struct {
		forest [][]int
		posx   int
		posy   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				forest: [][]int{
					{0, 0, 1, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 0,
				posx: 2,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				forest: [][]int{
					{0, 2, 1, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 0,
				posx: 2,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				forest: [][]int{
					{2, 0, 2, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 0,
				posx: 2,
			},
			want: false,
		},
		{
			name: "case4",
			args: args{
				forest: [][]int{
					{1, 1, 2, 3},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 0,
				posx: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkVisibilityLeft(tt.args.forest, tt.args.posx, tt.args.posy); got != tt.want {
				t.Errorf("checkVisibilityLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkVisibilityUp(t *testing.T) {
	type args struct {
		forest [][]int
		posx   int
		posy   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				forest: [][]int{
					{0, 0, 0, 0},
					{1, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 1,
				posx: 0,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				forest: [][]int{
					{2, 0, 0, 0},
					{1, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 1,
				posx: 0,
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				forest: [][]int{
					{2, 0, 0, 0},
					{1, 0, 0, 0},
					{3, 0, 0, 0},
				},
				posy: 1,
				posx: 0,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				forest: [][]int{
					{2, 0, 0, 0},
					{1, 0, 0, 0},
					{3, 0, 0, 0},
				},
				posy: 0,
				posx: 0,
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				forest: [][]int{
					{2, 0, 0, 0},
					{1, 0, 0, 0},
					{3, 0, 0, 0},
				},
				posy: 2,
				posx: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkVisibilityUp(tt.args.forest, tt.args.posx, tt.args.posy); got != tt.want {
				t.Errorf("checkVisibilityUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkVisibilityDown(t *testing.T) {
	type args struct {
		forest [][]int
		posx   int
		posy   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				forest: [][]int{
					{2, 0, 0, 0},
					{1, 0, 0, 0},
					{3, 0, 0, 0},
				},
				posy: 2,
				posx: 0,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				forest: [][]int{
					{2, 0, 0, 0},
					{1, 0, 0, 0},
					{3, 0, 0, 0},
				},
				posy: 1,
				posx: 0,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				forest: [][]int{
					{2, 0, 0, 0},
					{1, 0, 0, 0},
					{3, 0, 0, 0},
				},
				posy: 0,
				posx: 0,
			},
			want: false,
		},
		{
			name: "case4",
			args: args{
				forest: [][]int{
					{2, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				posy: 0,
				posx: 0,
			},
			want: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkVisibilityDown(tt.args.forest, tt.args.posx, tt.args.posy); got != tt.want {
				t.Errorf("checkVisibilityDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkVisibility(t *testing.T) {
	type args struct {
		forest [][]int
		posx   int
		posy   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				forest: [][]int{
					{3, 0, 3, 7, 3},
					{2, 5, 5, 1, 2},
					{6, 5, 3, 3, 2},
					{3, 3, 5, 4, 9},
					{3, 5, 3, 9, 0},
				},
				posy: 1,
				posx: 3,
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				forest: [][]int{
					{3, 0, 3, 7, 3},
					{2, 5, 5, 1, 2},
					{6, 5, 3, 3, 2},
					{3, 3, 5, 4, 9},
					{3, 5, 3, 9, 0},
				},
				posy: 2,
				posx: 2,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				forest: [][]int{
					{3, 0, 3, 7, 3},
					{2, 5, 5, 1, 2},
					{6, 5, 3, 3, 2},
					{3, 3, 5, 4, 9},
					{3, 5, 3, 9, 0},
				},
				posy: 3,
				posx: 1,
			},
			want: false,
		},
		{
			name: "case4",
			args: args{
				forest: [][]int{
					{3, 0, 3, 7, 3},
					{2, 5, 5, 1, 2},
					{6, 5, 3, 3, 2},
					{3, 3, 5, 4, 9},
					{3, 5, 3, 9, 0},
				},
				posy: 3,
				posx: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkVisibility(tt.args.forest, tt.args.posx, tt.args.posy); got != tt.want {
				t.Errorf("checkVisibility() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countVisible(t *testing.T) {
	type args struct {
		forest [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case4",
			args: args{
				forest: [][]int{
					{3, 0, 3, 7, 3},
					{2, 5, 5, 1, 2},
					{6, 5, 3, 3, 2},
					{3, 3, 5, 4, 9},
					{3, 5, 3, 9, 0},
				},
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVisible(tt.args.forest); got != tt.want {
				t.Errorf("countVisible() = %v, want %v", got, tt.want)
			}
		})
	}
}
