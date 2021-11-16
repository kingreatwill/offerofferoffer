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

func TestSlice03(t *testing.T) {
	var nil_slice []int
	println(nil_slice)
	t.Log(nil_slice) // []
	for s := range nil_slice {
		t.Log(s)
		t.Log("in slice")
	}

	empty_slice := make([]int, 0)
	println(empty_slice) // [0/0]0xc000049f48
}

func TestSlice04(t *testing.T) {
	_ = `
	SLICE[A:B]
	SLICE[A:B:C]
	SLICE[A:]  // 从A切到最尾部
	SLICE[:B]  // 从最开头切到B(不包含B)
	SLICE[:]   // 从头切到尾，等价于复制整个SLICE
	`
	my_slice := []int{11, 22, 33, 44, 66, 55, 77, 88, 99}
	println(my_slice) // [9/9]0xc0000a40f0
	t.Log(my_slice)   // [11 22 33 44 66 55 77 88 99]
	// 生成新的slice，从第二个元素取，切取的长度为2, 容量为8
	new_slice := my_slice[1:3]
	println(new_slice) // [2/8]0xc0000a40f8
	t.Log(new_slice)   // [22 33]
	new_slice[1] = 66
	t.Log(my_slice)  // [11 22 66 44 66 55 77 88 99]
	t.Log(new_slice) // [22 66]

	new_slice = my_slice[1:3:5]
	println(new_slice) // [2/4]0xc0000a40f8
	t.Log(new_slice)   // [22 66]

	new_slice = my_slice[:]
	println(new_slice) // [9/9]0xc0000a40f0
	t.Log(new_slice)   // [11 22 66 44 66 55 77 88 99]

	my_slice[1] = 23
	println(new_slice[1]) // 23

	// 扩容后指针变化
	my_slice = append(my_slice, 100)
	println(my_slice) // [10/18]0xc0000fa090
	my_slice[1] = 101
	println(new_slice[1]) // 23

}
