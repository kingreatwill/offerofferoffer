package golang_offer_test

import (
	"testing"
)

func TestSlice(t *testing.T) {
	my_slice := make([]int, 3, 5)
	t.Log(my_slice)      // [0 0 0]
	t.Log(len(my_slice)) // 3;
	t.Log(cap(my_slice)) // 5;
	println(my_slice)    // [3/5]0xc00000a480

}
