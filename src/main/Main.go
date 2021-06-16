package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("hello world!")

	var date = "2020-12-31"
	var url = "date=%s"
	var target_url = fmt.Sprintf(url, date)
	fmt.Println(target_url)

	var s string = "this is a test"
	fmt.Println(s)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var aa int
	fmt.Println(aa)

	var bb bool
	fmt.Println(bb)

	a1, b1, c1 := 1 ,2, "3"
	fmt.Println(a1, b1, c1)

	const LENGTH int = 100
	const WIDTH = 200

	var area int
	area = LENGTH * WIDTH
	fmt.Printf("area:%s", area)

	const (
		Unkonwn = 0
		Female = 1
		Male = 2
		str = "213213asasd"
		length = len(str)
		size = unsafe.Sizeof(str)
	)

	fmt.Println(Unkonwn, Female, Male, length, size)

	const (
		a = iota
		d
		e
	)
	fmt.Println(a,b1,c1,d,e)

	const (
		q = 1 << iota
		w = 4 << iota
		r
		t
	)
	fmt.Println(q,w,r,t)
	if (q == 1) {
		fmt.Println("等于1")
	}
	switch q {
		case 1: fmt.Println("1asdasdasdsd")
		default:
			fmt.Println("deadasdas")

	}

	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	fmt.Println(max(1,3))
	fmt.Println(swap("1", "2"))
	fmt.Println(swap1("1", "2", "c"))

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
	fmt.Println(test)

	var n[10] int

	for i := 0; i < 10; i++ {
		n[i] = i
	}
	for i := 0; i < len(n); i++ {
		fmt.Println(i)
	}

	var i,j,k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i] )
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j] )
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1:2.0,3:7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k] )
	}


	var aaaa int = 232/* 声明实际变量 */
	var ip *int/* 声明指针变量 */
	ip = &aaaa/* 指针变量的存储地址 */

	fmt.Println(&aaaa)
	fmt.Println(ip)/* 指针变量的存储地址 */
	fmt.Println(*ip)/* 使用指针访问值 */

	swapInt := func(x , y *int) {
		var n int
		n = *x
		*x = *y
		*y = n

	}
	qq, ww := 12, 34
	swapInt(&qq, &ww)
	fmt.Println(qq, ww)

}
var test int = 1999
func max(num1, num2 int) int  {
	var rst int
	if num1 > num2 {
		 rst = num1
	}else {
		rst = num2
	}
	return  rst
}

func swap(x, y string) (string, string) {
	return y, x
}

func swap1(x,y,z string) (string, string, string) {
	return y, x, z
}