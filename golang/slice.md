# slice(切片)

## slice的存储结构
Go中的slice依赖于数组，它的底层就是数组，所以数组具有的优点，slice都有。且slice支持可以通过append向slice中追加元素，长度不够时会动态扩展，通过再次slice切片，可以得到得到更小的slice结构，可以迭代、遍历等。

实际上slice是这样的结构：先创建一个有特定长度和数据类型的底层数组，然后从这个底层数组中选取一部分元素，返回这些元素组成的集合(或容器)，并将slice指向集合中的第一个元素。换句话说，slice自身维护了一个指针属性，指向它底层数组中的某些元素的集合。

```golang
my_slice := make([]int, 3, 5)
fmt.Println(my_slice)    // 输出：[0 0 0]
```
表示先声明一个长度为5、数据类型为int的底层数组，然后从这个底层数组中从前向后取3个元素(即index从0到2)作为slice的结构。

源码：go/src/runtime/slice.go
结构体定义
```golang
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
![](img/slice_01.png)

每一个slice结构都由3部分组成：容量(capacity)、长度(length)和指向底层数组某元素的指针，它们各占8字节(1个机器字长，64位机器上一个机器字长为64bit，共8字节大小，32位架构则是32bit，占用4字节)，所以任何一个slice都是24字节(3个机器字长)。

- Pointer：表示该slice结构从底层数组的哪一个元素开始，该指针指向该元素
- Capacity：即底层数组的长度，表示这个slice目前最多能扩展到这么长
- Length：表示slice当前的长度，如果追加元素，长度不够时会扩展，最大扩展到Capacity的长度(不完全准确，后面数组自动扩展时解释)，所以Length必须不能比Capacity更大，否则会报错

可以通过len()函数获取slice的长度，通过cap()函数获取slice的Capacity。
```go
my_slice := make([]int,3,5)
fmt.Println(len(my_slice))  // 3
fmt.Println(cap(my_slice))  // 5
```

还可以直接通过print()或println()函数去输出slice，它将得到这个slice结构的属性值，也就是length、capacity和pointer：
```go
my_slice := make([]int,3,5)
println(my_slice)      // [3/5]0xc42003df10
```

`[3/5]`表示length和capacity，`0xc42003df10`表示指向的底层数组元素的指针。
务必记住slice的本质是`[x/y]0xADDR`，记住它将在很多地方有助于理解slice的特性。另外，个人建议，虽然slice的本质不是指针，但仍然可以将它看作是一种包含了另外两种属性的不纯粹的指针，也就是说，直接认为它是指针。其实不仅slice如此，map也如此。

## 创建、初始化、访问slice
- make()：
```go
// 创建一个length和capacity都等于5的slice
slice := make([]int,5)
// length=3,capacity=5的slice
slice := make([]int,3,5)
```
make只能构建slice、map和channel这3种结构的数据对象，因为它们都指向底层数据结构，都需要先为底层数据结构分配好内存并初始化。
- new():
```golang
my_slice := new([]int)
fmt.Println(my_slice) // &[]
```
make()比new()函数多一些操作，new()函数只会进行内存分配并做默认的赋0初始化，而make()可以先为底层数组分配好内存，然后从这个底层数组中再额外生成一个slice并初始化。


- 直接赋值初始化：
```golang
// 创建长度和容量都为4的slice，并初始化赋值
color_slice := []string{"red","blue","black","green"}
// 创建长度和容量为100的slice，并为第100个元素赋值为3
slice := []int{99:3}
```
注意区分array和slice：
```golang
// 创建长度为3的int数组
array := [3]int{10, 20, 30}
// 创建长度和容量都为3的slice
slice := []int{10, 20, 30}
```
由于slice底层是数组，所以可以使用索引的方式访问slice，或修改slice中元素的值：
```golang
// 创建长度为5、容量为5的slice
my_slice := []int{11,22,33,44,55}
// 访问slice的第2个元素
print(my_slice[1])
// 修改slice的第3个元素的值
my_slice[2] = 333
```
lice能被访问的元素只有length范围内的元素，那些在length之外，但在capacity之内的元素暂时还不属于slice，只有在slice被扩展时(见下文append)，capacity中的元素才被纳入length，才能被访问。


## nil slice和空slice
https://www.cnblogs.com/f-ck-need-u/p/9854932.html