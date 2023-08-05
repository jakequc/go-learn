package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// fmt.Printf("My weight on the surface of Mars is %v lbs", 149.0*0.3783);
	// fmt.Printf(" and I would be %45vv years old. \n", 41*365/687);

	// fmt.Printf("My weight on the surface of %v is %-15v lbs.\n", "Earth",149.0)

	// var distance = 5600000;
	// var speed = 100080;

	// var (
	// 	dis1=50000
	// 	spee = 1088800
	// )

	// const housPeryDay, minutesPerHour = 24,60
	// fmt.Println("333",housPeryDay,minutesPerHour,dis1,spee,distance,speed)

	/// demo2

	// var num = rand.Intn(10) + 1
	// fmt.Println(num)

	// num = rand.Intn(10) + 1
	// fmt.Println(num)

	// fmt.Println("need speed is: ", 56000000.0/28.0/24)

	/// demo3 bool
	// fmt.Println("You find yourself in a dimly lit cavern.")
	// var command = "walk outside"

	// var exit = strings.Contains(command, "outside")

	// fmt.Println("You leave the cave: ", exit)

	/// demo4 if
	// var command = "go east"

	// if command == "go east" {
	// 	fmt.Println("You head further up the mountain.")
	// } else if command == "go inside" {
	// 	fmt.Print("You enter the cave where you live out the rest of your life.")
	// } else {
	// 	fmt.Println("Didn't quite get that.")
	// }

	// fmt.Println("The year is 2100, should you leap?")

	// var year = 2100
	// var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	// if leap {
	// 	fmt.Println("Look before you leap!")
	// } else {
	// 	fmt.Println("Keep your feet on the groud.")
	// }

	/// demo5 switch
	// fmt.Println("There is a cavern entrace here and a path to east.")
	// var command = "go inside"

	// switch command {
	// case "go east":
	// 	fmt.Println("You head further up the mountain")
	// case "enter cave", "go inside":
	// 	fmt.Println("You find yourself in a dimly lit cavern.")
	// case "read sign":
	// 	fmt.Println("the sign reads 'No minors'.")
	// default:
	// 	fmt.Println("Didn't quite get that.")
	// }

	// var room = "lake"
	// switch {
	// case room == "cave":
	// 	fmt.Println("You find yourself in a dimly lit cavern.")
	// case room == "lake":
	// 	fmt.Println("The ice seems solid enough.")
	// 	fallthrough // fall througn 将会指向当前的 case 和下一个 case
	// case room == "underwater":
	// 	fmt.Println("The water is freezing cold.")
	// case room == "3":
	// 	fmt.Println("3.")
	// }

	/// demo6 for
	// var count = 10

	// for count > 0 {
	// 	// for 中可以提前 break
	// 	fmt.Println(count)
	// 	time.Sleep(time.Second)
	// 	count--
	// }
	// fmt.Println("Liftoff!")

	// for { // 这样将会无限循环
	// 	println("loop Inifinited")
	// }

	/// demo7 猜数字

	// var age int
	// var answer = rand.Intn(10)*100 + 1
	// // fmt.Println("answer>>", answer)

	// for {
	// 	fmt.Println("Please input a number: ")
	// 	fmt.Scanf("%d", &age)
	// 	if age > answer {
	// 		fmt.Println("To big")
	// 	} else if age < answer {
	// 		fmt.Println("To small.")
	// 	} else {
	// 		fmt.Println("You win")
	// 		break
	// 	}
	// }

	/// demo8 var 和 短声明, go 语言中同一作用域不能有同名的声明，但是短声明不可以在 package 声明
	// var variable = 10 // variable := 10 一样效果
	// shortVar := 10    // 有的地方不支持 var 声明(比如 for, if,switch 循环中），可以使用短声明来代替

	// fmt.Println(variable, shortVar)

	// var count = 0
	// for count = 5; count > 0; count-- {
	// 	fmt.Println(count)
	// }

	// fmt.Println()
	// // 在 for 循环中的使用
	// for count := 5; count > 0; count-- {
	// 	fmt.Println(count)
	// }

	// // 短声明在 if 中的使用
	// if num := rand.Intn(3); num == 0 {
	// 	fmt.Println("Space Adventures")
	// } else if num == 1 {
	// 	fmt.Println("num == 1 is true")
	// } else {
	// 	fmt.Println("virgin Galactic")
	// }

	// // 在switch 语句中使用 switch， 声明 num 后对 num  switch
	// switch num := rand.Intn(10); num {
	// case 0:
	// 	fmt.Println("Zero")
	// case 1:
	// 	fmt.Println("One")
	// default:
	// 	fmt.Println("other")
	// }

	// // 生成随机年份，闰年2月为29天，非闰年2月为28天，使用 for 循环来生成和展示 10 个日期
	// print10Year()

	// /// log 浮点数类型
	// third := 1.0 / 3
	// fmt.Println(third)

	// fmt.Printf("%v\n", third)
	// fmt.Printf("%f\n", third)

	// fmt.Printf("%.3f\n", third)
	// fmt.Printf("%04.3f\n", 1.32)
	// // 4.2 表示宽度为4，小数位置为2 %[0width.precision]f 类型
	// fmt.Printf("%4.2f\n", third)

	// /// 浮点类型的精度问题： 浮点类型不适用于金融类计算机，为了尽量最小化舍入错误，建议先做乘法，再做除法

	// fmt.Println(third + third + third) // 1

	// sum := 0.1
	// sum += 0.2
	// fmt.Println(sum) // 0.30000000000000004

	// /// 作业题： 随机
	// /*
	// 	编写一个程序：
	// 	随机地将五分镍币（0.05美元）、一角硬币（0.10美元）和 25 美分硬币（0.25美元）放入一个空的储蓄罐，直到里面至少有20美元。
	// 	每次存款后显示存钱罐的余额
	// 	并以适当的宽度和精度格式化。
	// */
	// money := [3]float64{0.05, 0.10, 0.25}
	// sum = 0.00
	// var count = 0

	// for {
	// 	idx := rand.Intn(3)
	// 	fmt.Printf("%v th put %4.2f money", count, money[idx])
	// 	fmt.Println()
	// 	count++

	// 	if isEqual(sum, 20.00) {
	// 		break
	// 	}
	// 	sum += money[idx]
	// }

	// fmt.Println("sum = ", sum)

	/// 常用类型， int 是有符号类型的， uint 是无符号整数类型的
	// int 和 unit 是针对目标设备优化的类型，在比较老的移动设备上，int 和 uint 都是 32 位的，在比较新的计算机上，int 和 unit 都是 64 位的
	// 如果在较老的 32 位设备上，使用了超过 20 亿的整数，并且代码还能运行，那么最好使用 int64 和 unit64 来代替 int 和 uint
	// year := 2023
	// // %T 是获取对应变量的类型
	// fmt.Printf("Type %T for %v\n", year, year)

	// a := "text"
	// fmt.Printf("Type %T for %[1]v\n", a)

	// b := 42
	// fmt.Printf("Type %T for %[1]v\n", b)

	// c := 3.14
	// fmt.Printf("Type %T for %[1]v\n", c)

	// d := true
	// fmt.Printf("Type %T for %[1]v\n", d)

	// /// uint8 可以用来表示 8 位颜色值（如 红绿蓝 0-255)
	// var red, green, blue uint8 = 0, 141, 213
	// // 十六进制表示法
	// var red0x, green0x, blue0x uint8 = 0x00, 0x8d, 0xd5
	// // 使用 %x 格式化动词
	// fmt.Printf("%x %x %x\n", red, green, blue)

	// // 也可以指定最小宽度和填充
	// // 0 表示不足宽度是的填充值，2表示要打印的宽度
	// fmt.Printf("color: #%02x%02x%02x\n", red, green, blue)
	// fmt.Printf("color: #%02x%02x%02x\n", red0x, green0x, blue0x)

	// 为什么颜色不能使用 int 呢，因为 uint8 的取值范围正好合适，而 int 则都出来几十亿不合理的数
	// 如果很多颜色数据连续存储，例如未被压缩的图片，那么使用 uint8 可以节省很多内存

	// intMoneyStore()

	/// demo9 大整数
	// const lightSpeed = 299792
	// const secondsPerDay = 86400

	// var distance int64 = 41.3e12
	// fmt.Println("Alpha Centauri is", distance, "km away")

	// days := distance / lightSpeed / secondsPerDay
	// fmt.Println("That is", days, "days of travel at light speed.")

	// lightSpeed2 := big.NewInt(30e10) // 如果 big.NewInt 传递的数字太大，会报错
	// secondsPerDay2 := big.NewInt((86400))
	// fmt.Println(lightSpeed2, secondsPerDay2)

	// distance2 := new(big.Int) // 创建一个 Big.Int 实例
	// distance2.SetString("2400000000000000", 10)
	// fmt.Println(distance2)

	// rune 的类型是为了表示 unicode code point （即每个字符分配了相应的数值），run 和 int32 是别名关系，因此两者可以互相使用
	// byte 是 uint8 类型的别名，目的是用于二进制数据， byte 可以表示由 ASCII 定义的英语自负，它是 Unicode 的一个子集（共 128 个字符）

	// fmt.Println("peace be upon you \n upon you be peace")
	// fmt.Println(`strings can span multiple lines with the \n escape sequence`)

	// fmt.Println(`peace be upon you
	// upon you be peace`)

	// var pi rune = 960 // rune 是 int 的别名
	// var alpha rune = 940
	// var omega rune = 969
	// var bang byte = 33 // byte 是 uint 的别名

	// // %v 将会打印 code point 的值
	// fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)

	// fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)

	// fmt.Printf("%v %v \n", '*', 'é')

	// print("shalom")
	// loopEncode()

	// c := 'g'

	// c = c - 'a' + 'A'

	// fmt.Println(c)

	// message := "uv vagreangvbany fcnpr fgngvba"
	// for i := 0; i < len(message); i++ {
	// 	c := message[i]
	// 	if c >= 'a' && c <= 'z' {
	// 		c = c + 13
	// 		if c > 'z' {
	// 			c = c - 26
	// 		}
	// 	}

	// 	fmt.Printf("%c", c)
	// }

	// question := "?Como estasd j"
	// fmt.Println(len(question), "bytes")

	// fmt.Println(utf8.RuneCountInString(question), "runes")

	// // go 语言中函数支持返回多个值
	// c, size := utf8.DecodeRuneInString(question)
	// fmt.Printf("First rune: %c %v bytes", c, size)

	// fmt.Println("--------- split -------")

	// for i, c := range question {
	// 	fmt.Printf("%v %c\n", i, c)
	// }

	/// c
	// str := "L fdph, L vdz, L frqtxhuhg"

	// for i := 0; i < len(str); i++ {
	// 	c := str[i]
	// 	c -= 3
	// }

	////////********************** demo test ****************

	/*
	   把西班牙语 “Hola Estación Espacial Internacional” 用 ROT13 进行加密
	   使用 range 关键字
	   带重音符号的字母要保留
	*/
	// plaintext := "Hola Estación Espacial Internacional"

	// ciphertext := rot13Encrypt(plaintext)
	// fmt.Println("加密后的结果为： ", ciphertext)

	// /*
	//   L fdph, L vdz, L frqtxhuhg，每个字母向前移动 3个位置，能得到什么字符串？
	// */
	// plaintext = "L fdph, L vdz, L frqtxhuhg"
	// ciphertext = caesarEncrypt(plaintext, 3)
	// fmt.Println("加密后的结果: ", ciphertext)

	// v := 42
	// if v >= 0 && v <= math.MaxUint8 {
	// 	v8 := uint8(v)
	// 	fmt.Println("converted: ", v8)
	// }

	// // strconv 包 的 itoa 函数可以将 int 转化为 string， Itoa 的意思是 Integer to ASCII 的意思
	// str := "Launch in T minus " + strconv.Itoa(10) + " seconds."
	// fmt.Println(str)

	// // 字符串转换的方式， strconv.xxx or fmt.Sprintf(x1,x2)
	// countdown := 9
	// str = fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	// fmt.Println(str)

	// count, err := strconv.Atoi(("10ds"))

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(count)

	// fmt.Println(bool2Int(true))
	// fmt.Println(bool2Int(false))

	// inputs := []string{"true", "yes", "1", "false", "no", "0", "other"}

	// for _, input := range inputs {
	// 	result, err := str2bool(input)
	// 	if err != nil {
	// 		fmt.Print(err)
	// 	} else {
	// 		fmt.Printf("str \"%s\" convert boolean: %v\n", input, result)
	// 	}
	// }

	// fmt.Println()
	// res := vigenereEncrypt("ISOITEUIWUIZNSROCNKFD", "GOLANG")
	// fmt.Println(res)

	/// demo 函数
	// kelvin := 233.0
	// celsius := kelvinToCelsius((kelvin))
	// fmt.Println(kelvin, "K is ", celsius, " C")

	// fmt.Println(celsiusToFahrenheit(23.0))
	// fmt.Println(kelvinToFahrenheit(0.0))

	/// demo 类型别名 type aliastype othertype, 类型别名是新的类型，不能和原来的类型混用
	// type celsius float64

	// const degrees = 20

	// var temperature celsius = degrees
	// temperature += 10

	// fmt.Println(temperature)

	// var k keivin = 294.0
	// var c celsius

	// c = kelvin2Celsius(k)

	// // 方法的调用
	// c = k.celsius()
	// fmt.Println(c)

	// // 变量也接受函数
	// print := fn
	// print(1, 2, 3, 4)

	/// 匿名函数&闭包
	// f()

	// /*
	// 	匿名函数就是没有名字的函数，也叫做函数字面值
	// 	函数字面值需要保留外部作用域变量的引用，所以形成了闭包，即函数字面值是可闭包的
	// */
	// f1 := func(message string) {
	// 	fmt.Println(message)
	// }

	// f1("I love you...")

	// // 立即执行函数
	// func() {
	// 	fmt.Println("Immediately running...")
	// }()

	// // add 是 addWapper 返回的一个方法
	// addOne := addWapper(1)
	// fmt.Println(addOne(2)) // print 3

	/// 温度表作业
	//  第一个表格： 摄氏度转华氏度
	// fmt.Println("摄氏度转华氏度table: ")
	// drawTable(celsiusToFahrenheitw)

	// //  第2个表格： 华氏度转摄氏度
	// fmt.Println("华氏度转摄氏度table: ")
	// drawTable(fahrenheitToCelsius)

	// 数组的复制
	// 不论数组赋值给新的变量还是将它传递给函数，都会产生一个完整的数组副本
	// defaultCloneDeepArr()

	// 因为数组传递给函数参数或者是赋值给变量的时候是深克隆（完整的副本），所以在传递数组给函数参数的时候效率就非常低
	// terrformDemo()

	// twoDimensionalArray()
	// twoDimensionalList()

	/*
		切分数组就是切片，它不会修改数组，他只是创建了指向数组的一个窗口或视图，这种视图就是 slice 类型
		切片是的方式是 arr[start:end] ,切的范围是前开后必的，即 [start,end), 打给你 start 没有的时候默认是 0， end没有的时候默认是 arr 的长度
	*/
	// sliceDemo()
	// sliceStr()

	// sliceDeclaration()

	// slicePowerful()
	// stringSliceSort()

	// newSliceTerraformTask()

	// appendSliceEle()

	// lenAndCapacityDemo()

	// sliceCount()
	sliceSpecifiedCapacity()
}

func terraformDemo(prefix string,worlds, ...string)[]string {
	newWorlds := make([]string,len(worlds));f
}

func sliceSpecifiedCapacity() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune",
	}

	// 指定 terrestrial 的 startIdx=0， endIdx = 4, capacity is 4
	terrestrial := planets[0:4:4]

	worlds := append(terrestrial, "Ceres")

	dump("planets", planets)
	dump("terrestrial", terrestrial)
	dump("worlds", worlds)
}

func sliceCount() {
	dwarfs1 := []string{"Ceres", "Puto", "Haumea", "Makemake", "Eris"}
	dwarfs1[1] = "333"
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia", "Auaoar", "Sedna")

	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)
}

func lenAndCapacityDemo() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])
}

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

// add elements to slice
func appendSliceEle() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dwarfs = append(dwarfs, "Orcus")
	dwarfs = append(dwarfs, "Orcus2", "Orcus3")

	fmt.Print(dwarfs)
}

type planets []string

// associate terraform methods for planets types
func (s planets) terraform() {
	for idx := range s {
		s[idx] = "New" + s[idx]
	}
}

func newSliceTerraformTask() {
	p := planets{"A", "C", "B"}

	fmt.Println("before: ", p)

	p.terraform()

	fmt.Println("after: ", p)
}

func stringSliceSort() {
	planets := []string{"B", "A", "C", "E", "D", "F"}

	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
}

func slicePowerful() {
	planets := []string{"A ", " Earth", "Mars  "}

	fmt.Println("before: ", planets)

	hyperspace(planets)

	fmt.Println(strings.Join(planets, "-"))
	fmt.Println("after: ", planets)
}

// slice can accept any length string slice
// slice can mutation value of passed
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

// There are two ways to declare slice
func sliceDeclaration() {
	arr := [...]string{"A", "B", "C"}

	sliceDeclaration1 := arr[:]

	// 因为 切片的长度时变化的，所以可以不用给 [] 添加长度
	sliceDeclaration2 := []string{"A", "B"}

	fmt.Println(sliceDeclaration1, sliceDeclaration2)
	fmt.Printf("%v is %T type\n", arr, arr)
	fmt.Printf("%v is %T type\n", sliceDeclaration1, sliceDeclaration1)
}

func sliceStr() {
	// 切分数组的语法也适用于切分字符串
	// 切分字符串时，索引代表的是字节 而非 rune 数
	str := "abcd"

	c := str[0:1]
	fmt.Printf("%v is %T type", c, c)
}

func sliceDemo() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[:4] // 等于 planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:] // equal to planets[6: len(panets)]

	allPlanets := planets[:] // 等与 planets[0:len(planets)]

	fmt.Println(terrestrial, gasGiants, iceGiants, allPlanets)
}

func twoDimensionalList() {
	var board [8][8]string
	chess := "kqrbnpKQRBNP"

	for outIdx, element := range board {
		for innerIdx := range element {
			currChar := chess[rand.Intn(12)]
			board[outIdx][innerIdx] = string(currChar)
		}
	}

	print2DimensionalList(board)

}

func print2DimensionalList(twoDimensionalList [8][8]string) {
	fmt.Println(strings.Repeat("-", 40))
	for _, outEle := range twoDimensionalList {
		fmt.Print("|")
		for _, currEle := range outEle {
			fmt.Printf("%2s | ", currEle)
		}
		fmt.Println()
		fmt.Println(strings.Repeat("-", 40))
	}
}

// 二维数组 demo
func twoDimensionalArray() {
	var board [8][8]string

	board[0][0] = "r"
	board[0][0] = "r"

	for column := range board[1] {
		board[1][column] = "p"
	}

	fmt.Println(board)
}

func terrformDemo() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terraform(planets)
	fmt.Println("after:", planets)
}

func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New-" + planets[i]
	}
	fmt.Println("in-terraform", planets)
}

func defaultCloneDeepArr() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	planetsMarkII := planets

	planets[2] = "whoops"
	fmt.Println(planets)
	// 不影响 planets 的值，因为 go 中的数组默认是产生一个完整的副本（即深克隆）
	fmt.Println(planetsMarkII)

}

func arraydemo() {

	/// 数组： 固定长度且有序的元素集合
	// defined: var arrayName [len]type

	// 复合字面值初始化数组
	// arr := [len]type{v1,v2,v3,...,vlen-1}
	// arr := [...]type{v1,v2,...,vn} ... 会根据后边大括号设置的初始值来去计算长度
	var planets [8]string

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	fmt.Println(earth)

	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	// 第一种遍历方式
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}

	// 第二种遍历方式
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}
}

// 定义温度转化函数： 摄氏度转化华氏度
func celsiusToFahrenheitw(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// 华氏度转摄氏度
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// 自定义填充值函数
func drawValue(value float64) {
	fmt.Printf("| %11.2f", value)
}

// 定义重复字符串函数
func repeatStr(str string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += str
	}

	return result
}

// 定义划线函数
func drawLine(length int) {
	fmt.Println("+", repeatStr("-", length), "+")
}

func drawTable(converter func(float64) float64) {
	drawLine(23)
	fmt.Println("|   摄氏度   |   华氏度   |")
	drawLine(23)

	for celsius := -40.0; celsius <= 100.0; celsius += 5.0 {
		fahrenheit := converter(celsius)
		drawValue((celsius))
		drawValue(fahrenheit)
		fmt.Println("|")
	}

	drawLine(23)
}

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 1
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

// 定义返回函数的类型
type addFn func(num float64) float64

// addWapper 返回一个函数
func addWapper(num1 float64) addFn {
	return func(num2 float64) float64 {
		// 函数内部引用了外部的 num1， 形成了闭包
		return num1 + num2
	}
}

// 把匿名函数赋值给 f
var f = func() {
	fn(1, 2, 3, 4)
}

func fn(a ...interface{}) {
	fmt.Println(a)
}

/// 通过方法添加行为， go 中将某个类型和方法相关联（组合）
// 可以将方法与同包中声明的任何类型相关联，但不可以是 int，float64 等预声明的类型进行关联

type celsius float64
type keivin float64

// 这是一个方法，因为关联了 keivin 类型
func kelvin2Celsius(k keivin) celsius {
	return celsius(k - 273.15)
}

func (K keivin) celsius() celsius {
	return celsius(K - 273.15)
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9.0/5.0 + 32.0
}

func kelvinToFahrenheit(kelvin float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(kelvin))
}

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

/*
维吉尼亚加密是针对大写的 明文
*/
func vigenereEncrypt(plaintext, key string) string {
	plaintext = strings.ToUpper(plaintext)
	key = strings.ToUpper(key)

	var ciphertext string
	keyIndex := 0

	for i := 0; i < len(plaintext); i++ {
		if plaintext[i] >= 'A' && plaintext[i] <= 'Z' {
			// Calculate the shift for this character
			shift := int(key[keyIndex] - 'A')

			// Encrypt the character
			encryptedChar := byte((int(plaintext[i]-'A')+shift)%26 + 'A')
			ciphertext += string(encryptedChar)

			// Move to the next character in the key (loop around if needed)
			keyIndex = (keyIndex + 1) % len(key)
		} else {
			// Non-alphabetic characters remain unchanged
			ciphertext += string(plaintext[i])
		}
	}

	return ciphertext
}

func str2bool(str string) (bool, error) {
	switch strings.ToLower(str) {
	case "true", "yes", "1":
		return true, nil
	case "false", "no", "0":
		return false, nil
	default:
		return false, fmt.Errorf("invalid string: %s", str)
	}

}

func bool2Int(flag bool) int {
	if flag {
		return 1
	}

	return 0
}

func caesarEncrypt(plaintext string, shift int) string {
	var result string

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			// 对大写字母进行位移加密
			result += string(rune((int(char-'A')+26-shift)%26 + 'A'))
		} else if char >= 'a' && char <= 'z' {
			// 对小写字母进行位移加密
			result += string(rune((int(char-'a')+26-shift)%26 + 'a'))
		} else {
			// 非字母字符保持不变
			result += string(char)
		}
	}

	return result
}

func rot13Encrypt(plaintext string) string {
	var result string

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			// 对大写字母进行 ROT13 加密
			// char-'A'+13 是计算出当前的位置 后移动 13 位
			result += string((char-'A'+13)%26 + 'a')
		} else if char >= 'a' && char <= 'z' {
			// 对小写字母进行 ROT13 加密
			result += string((char-'a'+13)%26 + 'a')
		} else {
			// 非字母字符保持不变
			result += string(char)
		}
	}

	return result
}

func loopEncode() {
	c := 'a'
	c = c + 3
	fmt.Printf("%c", c)

	if c > 'z' {
		c -= 26
	}

	fmt.Printf("%c\n", c)
}

func print(str string) {
	for i := 0; i < len(strings.Split(str, "")); i++ {
		fmt.Printf("%c\n", str[i])
	}
}

func intMoneyStore() {
	moneyTypes := [3]int{5, 10, 25}
	sum := 0
	for {
		if isEqualInt(sum, 25) {
			break
		}

		sum += moneyTypes[rand.Intn(3)]
		fmt.Printf("current remainder %f $\n", float64(sum/10.0))
	}

	fmt.Println("sum money", sum)
}

func isEqual(num1 float64, num2 float64) bool {
	return math.Abs(num1-num2) <= 0.00001
}

func isEqualInt(num1 int, num2 int) bool {
	return math.Abs(float64(num1/10.0-num2)) <= 0.0001
}

func isLeap(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func print10Year() {
	for y := 0; y < 10; y++ {
		var randomYear = rand.Intn(9999)
		println("%v year is leap year: ", randomYear, isLeap((randomYear)))
	}
}
