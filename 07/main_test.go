package main

import (
	"reflect"
	"testing"
)

func Test_changeDir(t *testing.T) {
	type args struct {
		current string
		dir     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ChangeIntoA",
			args: args{
				current: "/foo",
				dir:     "A",
			},
			want: "/foo/A",
		},
		{
			name: "ChangeRootIntoA",
			args: args{
				current: "/",
				dir:     "A",
			},
			want: "/A",
		},
		{
			name: "ChangeRootBackwards",
			args: args{
				current: "/",
				dir:     "..",
			},
			want: "/",
		},
		{
			name: "ChangeFooABackwards",
			args: args{
				current: "/Foo/A",
				dir:     "..",
			},
			want: "/Foo",
		},
		{
			name: "ChangeFooBackwards",
			args: args{
				current: "/Foo",
				dir:     "..",
			},
			want: "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeDir(tt.args.current, tt.args.dir); got != tt.want {
				t.Errorf("changeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFolders(t *testing.T) {
	type args struct {
		dir      string
		command  string
		computer map[string]folder
	}
	tests := []struct {
		name string
		args args
		want map[string]folder
	}{
		{
			name: "AddFooAtoFoo",
			args: args{
				dir:     "/Foo",
				command: "A",
				computer: map[string]folder{
					"/Foo": {
						dirPath: "/Foo",
						folders: []string{},
						files:   []int{500},
					},
				},
			},
			want: map[string]folder{
				"/Foo": {
					dirPath: "/Foo",
					folders: []string{"/Foo/A"},
					files:   []int{500},
				},
				"/Foo/A": {
					dirPath: "/Foo/A",
					folders: []string{},
					files:   []int{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFolders(tt.args.dir, tt.args.command, tt.args.computer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFolders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFiles(t *testing.T) {
	type args struct {
		dir      string
		command  string
		computer map[string]folder
	}
	tests := []struct {
		name string
		args args
		want map[string]folder
	}{
		{
			name: "AddFooAtoFoo",
			args: args{
				dir:     "/Foo",
				command: "500",
				computer: map[string]folder{
					"/Foo": {
						dirPath: "/Foo",
						folders: []string{},
						files:   []int{},
					},
				},
			},
			want: map[string]folder{
				"/Foo": {
					dirPath: "/Foo",
					folders: []string{},
					files:   []int{500},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFiles(tt.args.dir, tt.args.command, tt.args.computer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_folder_sum(t *testing.T) {
	type fields struct {
		dirPath string
		folders []string
		files   []int
	}
	type args struct {
		computer map[string]folder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "sumTest1",
			fields: fields{
				dirPath: "/Foo",
				folders: []string{"/Foo/A"},
				files:   []int{500},
			},
			args: args{
				computer: map[string]folder{
					"/Foo": {
						dirPath: "/Foo",
						folders: []string{"/Foo/A"},
						files:   []int{500},
					},
					"/Foo/A": {
						dirPath: "/Foo/A",
						folders: []string{},
						files:   []int{300},
					},
				},
			},
			want: 800,
		},
		{
			name: "sumTest2",
			fields: fields{
				dirPath: "/Foo",
				folders: []string{"/Foo/A"},
				files:   []int{},
			},
			args: args{
				computer: map[string]folder{
					"/Foo": {
						dirPath: "/Foo",
						folders: []string{"/Foo/A"},
						files:   []int{},
					},
					"/Foo/A": {
						dirPath: "/Foo/A",
						folders: []string{},
						files:   []int{300},
					},
				},
			},
			want: 300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folder := folder{
				dirPath: tt.fields.dirPath,
				folders: tt.fields.folders,
				files:   tt.fields.files,
			}
			if got := folder.sum(tt.args.computer); got != tt.want {
				t.Errorf("folder.sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runLine(t *testing.T) {
	type args struct {
		dir      string
		command  string
		computer map[string]folder
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 map[string]folder
	}{
		// First Case
		{
			name: "completeRun",
			args: args{
				dir:     "/Foo",
				command: "$ cd /",
				computer: map[string]folder{
					"/Foo": {
						dirPath: "/Foo",
						folders: []string{},
						files:   []int{},
					},
					"/": {
						dirPath: "/",
						folders: []string{"/Foo"},
						files:   []int{},
					},
				},
			},
			want: "/",
			want1: map[string]folder{
				"/Foo": {
					dirPath: "/Foo",
					folders: []string{},
					files:   []int{},
				},
				"/": {
					dirPath: "/",
					folders: []string{"/Foo"},
					files:   []int{},
				},
			},
		},
		// Second Case
		{
			name: "completeRun",
			args: args{
				dir:     "/Foo",
				command: "$ cd ..",
				computer: map[string]folder{
					"/Foo": {
						dirPath: "/Foo",
						folders: []string{},
						files:   []int{},
					},
					"/": {
						dirPath: "/",
						folders: []string{"/Foo"},
						files:   []int{},
					},
				},
			},
			want: "/",
			want1: map[string]folder{
				"/Foo": {
					dirPath: "/Foo",
					folders: []string{},
					files:   []int{},
				},
				"/": {
					dirPath: "/",
					folders: []string{"/Foo"},
					files:   []int{},
				},
			},
		},
		// third Run
		{
			name: "completeRunAddFile100",
			args: args{
				dir:     "/Foo",
				command: "100 fo",
				computer: map[string]folder{
					"/Foo": {
						dirPath: "/Foo",
						folders: []string{},
						files:   []int{},
					},
					"/": {
						dirPath: "/",
						folders: []string{"/Foo"},
						files:   []int{},
					},
				},
			},
			want: "/Foo",
			want1: map[string]folder{
				"/Foo": {
					dirPath: "/Foo",
					folders: []string{},
					files:   []int{100},
				},
				"/": {
					dirPath: "/",
					folders: []string{"/Foo"},
					files:   []int{},
				},
			},
		},
		// fourth Run
		{
			name: "completeRunAddFile100",
			args: args{
				dir:     "/Foo",
				command: "dir A",
				computer: map[string]folder{
					"/Foo": {
						dirPath: "/Foo",
						folders: []string{},
						files:   []int{},
					},
					"/": {
						dirPath: "/",
						folders: []string{"/Foo"},
						files:   []int{},
					},
				},
			},
			want: "/Foo",
			want1: map[string]folder{
				"/Foo/A": {
					dirPath: "/Foo/A",
					folders: []string{},
					files:   []int{},
				},
				"/Foo": {
					dirPath: "/Foo",
					folders: []string{"/Foo/A"},
					files:   []int{},
				},
				"/": {
					dirPath: "/",
					folders: []string{"/Foo"},
					files:   []int{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := runLine(tt.args.dir, tt.args.command, tt.args.computer)
			if got != tt.want {
				t.Errorf("runLine() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("runLine() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
