package sort

import (
	"reflect"
	"testing"
)

func Test_bubblesort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"testO1", args{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"testO2", args{[]int{37, 8, 9, 0, 2, 3}}, []int{0, 2, 3, 8, 9, 37}},
		{"testO3", args{[]int{1, 2, 3, 4, 5, 8, 7}}, []int{1, 2, 3, 4, 5, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubblesort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubblesort() = %v, want %v", got, tt.want)
			}
		})
	}
}
