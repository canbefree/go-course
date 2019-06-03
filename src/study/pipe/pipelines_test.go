//Package pipe 管道

//@auther neo

package pipe

import (
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		num  int
		next NextClosure
	}
	tests := []struct {
		name string
		args args
		want NextClosure
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.num, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
