package data

import (
	"reflect"
	"testing"
)

func TestGetType(t *testing.T) {
	type args struct {
		d interface{}
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.

		{name: "int", args: args{12}, want: reflect.Int},
		{name: "string", args: args{"hello,wolrd"}, want: "string"},
		{name: "float", args: args{4.155}, want: "float64"},
		{name: "double", args: args{4.1551232131231312312313123113131313131313131231212313321}, want: "double"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetType(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}
