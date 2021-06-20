package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

	a1, b1, c1 := 1, 2, "3"
	fmt.Println(a1, b1, c1)

	const LENGTH int = 100
	const WIDTH = 200

	var area int
	area = LENGTH * WIDTH
	fmt.Printf("area:%s", area)

	const (
		Unkonwn = 0
		Female  = 1
		Male    = 2
		str     = "213213asasd"
		length  = len(str)
		size    = unsafe.Sizeof(str)
	)

	fmt.Println(Unkonwn, Female, Male, length, size)

	const (
		a = iota
		d
		e
	)
	fmt.Println(a, b1, c1, d, e)

	const (
		q = 1 << iota
		w = 4 << iota
		r
		t
	)
	fmt.Println(q, w, r, t)
	if q == 1 {
		fmt.Println("等于1")
	}
	switch q {
	case 1:
		fmt.Println("1asdasdasdsd")
	default:
		fmt.Println("deadasdas")

	}

	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	fmt.Println(max(1, 3))
	fmt.Println(swap("1", "2"))
	fmt.Println(swap1("1", "2", "c"))

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
	fmt.Println(test)

	var n [10]int

	for i := 0; i < 10; i++ {
		n[i] = i
	}
	for i := 0; i < len(n); i++ {
		fmt.Println(i)
	}

	var i, j, k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	var aaaa int = 232 /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */
	ip = &aaaa         /* 指针变量的存储地址 */

	fmt.Println(&aaaa)
	fmt.Println(ip)  /* 指针变量的存储地址 */
	fmt.Println(*ip) /* 使用指针访问值 */

	swapInt := func(x, y *int) {
		var n int
		n = *x
		*x = *y
		*y = n

	}
	qq, ww := 12, 34
	swapInt(&qq, &ww)
	fmt.Println(qq, ww)

	type Books struct {
		title   string
		author  string
		subject string
		bookId  int
	}

	fmt.Println(Books{"go语言", "张三", "语文", 1232322222222})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookId: 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	var book1 Books
	book1.bookId = 213
	book1.author = "asdasdsdsdsad"
	book1.title = "dsafdaf"
	book1.subject = "adsfdsfadsf"
	book1.title = "wyg"

	fmt.Printf("Book 1 title : %s\n", book1.title)
	fmt.Printf("Book 1 author : %s\n", book1.author)
	fmt.Printf("Book 1 subject : %s\n", book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", book1.bookId)

	printBook := func(book *Books) {
		fmt.Printf("Book title : %s\n", book.title)
		fmt.Printf("Book author : %s\n", book.author)
		fmt.Printf("Book subject : %s\n", book.subject)
		fmt.Printf("Book book_id : %d\n", book.bookId)
	}
	printBook(&book1)
	// 切片
	var numbers1 = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers1), cap(numbers1), numbers1)

	printSlice := func(x []int) {
		fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	}

	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers11 := make([]int, 0, 5)
	printSlice(numbers11)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Println("打开数据库失败")
		return
	}
	fmt.Println("链接成功")

	var user User
	rows, e1 := db.Query("select * from hand")
	if e1 == nil {
		errors.New("查询错误")
	}
	for rows.Next() {
		e := rows.Scan(&user.xm, &user.sfzh)
		fmt.Println(user.xm, user.sfzh)
		if e == nil {
			//fmt.Println(json.Marshal(user))
		}
	}
	rows.Close()

}

type User struct {
	xm   string `db:"xm"`
	sfzh string `db:"sfzh"`
}

var test int = 1999

func max(num1, num2 int) int {
	var rst int
	if num1 > num2 {
		rst = num1
	} else {
		rst = num2
	}
	return rst
}

func swap(x, y string) (string, string) {
	return y, x
}

func swap1(x, y, z string) (string, string, string) {
	return y, x, z
}
