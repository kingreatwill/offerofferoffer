package golang_offer_test

import (
	"testing"
)

func TestSlice01(t *testing.T) {
	my_slice := make([]int, 3, 5)
	t.Log(my_slice)      // [0 0 0]
	t.Log(len(my_slice)) // 3;
	t.Log(cap(my_slice)) // 5;
	println(my_slice)    // [3/5]0xc00000a480

}

func TestSlice02(t *testing.T) {
	my_slice := new([]int)
	t.Log(my_slice) // &[]
	my_slice_2 := []int{5: 3}
	t.Log(my_slice_2) // [0 0 0 0 0 3]
	my_slice_3 := []string{5: "3"}
	t.Log(my_slice_3) // [     3]
}
