# Go语言入门

## 1.hello world

### 1)代码解释

    // 要求开发一个hello.go程序，可以输出hello world！

    package main

    import "fmt"

    func main() {

        fmt.Println("hello world!")

    }

* (1)go文件的后缀是.go
* (2)package main

    表示该hello.go文件所在的包是main,在go中，每个文件都必须归属于一个包

* (3)import "fmt"

    表示引入一个包，包名fmt，引入该包后，就可以使用fmt包的函数，比如：fmt.Println

* (4)func main(){

     }
     func是一个关键字，表示一个函数。
     main是函数名，是一个主函数，即我们程序的入口

* (5)fmt. Println("hello")

     表示调用fmt包的函数Println输出“hello world!”

### 2)运行

    go run hello.go

* 速度慢，占内存小

### 3)编译&运行

    go build .\hello.go
    .\hello.exe

    go build -o myhello.exe .\hello.go
    .\myhello.exe

* 速度快，占内存多
* 可将执行文件拷贝到没有go环境的机器上运行

## 2. Go程序开发注意事項

* (1)Go源文件以"go"为扩展名。
* (2)Go应用程序的执行入口是main()方法。
* (3)Go语言严格区分大小写。
* (4)Go方法由一条条语句构成，每个语句后不需要分好（Go语言会在每行后自动加分号），这也体现出Golang的简洁性。
* (5)Go编辑器是一行行进行编译的，因此我们一行就写一条语句，不能把多条语句写在同一个，否则报错。
* (6)Go语言定义的变量或者import的包如果没有使用到，代码不能编译通过。
* (7)大括号都是成对出现的，缺一不可。

## 3. Go语言转义字符

* \t : 一个制表位，实现对齐的功能
* \n : 换行符
* \ \ : 一个\
* \" : 一个"
* \r : 一个回车fmt. Println("天龙八部雪山飞狐\r张飞")

## 4. VS Code快捷键

### (1)编辑器与窗口管理

* 新建文件:   Ctrl+N
* 文件之间切换:   Ctrl+Tab
* 打开一个新的VS Code编辑器:    Ctrl+Shift+N
* 关闭当前窗口:   Ctrl+W
* 关闭当前的VS Code编辑器:   Ctrl+Shift+W
* 切出一个新的编辑器窗口（最多3个):   Ctrl+\
* 切换左中右3个编辑器窗口的快捷键:   Ctrl+1  Ctrl+2  Ctrl+3

### (2)代码编辑

* 代码行向左或向右缩进:   Ctrl+[ 、 Ctrl+]
* 代码格式化:   Shift+Alt+F
* 向上或向下移动一行:   Alt+Up 或 Alt+Down
* 向上或向下复制一行:   Shift+Alt+Up 或 Shift+Alt+Down
* 在当前行下方插入一行:   Ctrl+Enter
* 在当前行上方插入一行:   Ctrl+Shift+Enter

### (3)显示相关

* 全屏显示(再次按则恢复):   F11
* 放大或缩小(以编辑器左上角为基准):   Ctrl +/-
* 侧边栏显示或隐藏： Ctrl+B
* 显示 Debug:    Ctrl+Shift+D
* 显示搜索(光标切到侧边栏中才有效):   Ctrl+Shift+F
* 显示 Output:    Ctrl+Shift+U

# Golang变量

## 1. 变量使用的基本步骤

* 声明变量
* 赋值
* 使用

## 2. 变量快速入门

### (1)定义单个变量

    package main

    import "fmt"

    func main()  {

        // Golang的变量使用方式1
        // 第一种：指定变量类型，声明后若不赋值，使用默认值
        // int 的默认值是0
        var i int
        fmt. Println("i=", i)

        // 第二种：根据值自行判定变量类型（类型推导）
        var num = 10.11
        fmt.Println("num=", num)

        // 第三种：省略var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
        // 下面的方式等价 var name string name = "tom"
        name := "tom"
        fmt.Println("name=", name)

    }

### (2)定义多个变量和定义全局变量

    package main

    import "fmt"

    // 定义全局变量
    var n1 = 100
    var n2 = 200
    var name = "jack"
    // 上面的声明方式，也可以改成一次性声明
    var (
        n3 = 300
        n4 = 900
        name2 = "mary"
    )

    func main()  {

        // Golang如何一次性声明多个变量
        // var n1, n2, n3 int
        // fmt. Println("n1=", n1, "n2=", n2, "n3=", n3)

        

        // 一次性声明多个变量的方式2
        // var n1, name, n2 = 100, "tom", 888
        // fmt. Println("n1=", n1, "name=", name, "n2=", n2)

        

        // 一次性声明多个变量的方式3，同样可以使用类型推导
        // n1, name, n3 := 100, "tom", 888
        // fmt. Println("n1=", n1, "name=", name, "n3=", n3)
        fmt. Println("n3=", n3, "name=", name, "n4=", n4)

        

    }

## 3. 变量使用的注意事项

    package main

    import "fmt"

    // 变量的注意事项

    func main()  {

        

        // 该区域的数据值可以在同一类型范围内不断变化
        var i int = 10
        i = 30
        i = 50
        fmt. Println("i", i)
        // i = 1.2 //int, 原因是不能改变数据类型

        // 变量在同一个作用域（在一个函数或者代码块）内不能重名
        // var i int = 59
        // i := 99

    }

* 该区域的数据值可以在同一类型范围内不断变化
* 变量在同一个作用域内不能重名
* 变量 = 变量名 + 值 + 数据类型
* Golang的变量如果没有赋初始值，编译器会使用默认值，比如int默认值0、string默认值为空

## 4. 程序中+号的使用

    package main

    import "fmt"

    func main()  {

        var i = 1
        var j = 2
        var r = i + j //做加法运算
        fmt. Println("r=", r)

        var str1 = "hello "
        var str2 = "world"
        var res = str1 + str2  //做拼接操作
        fmt. Println("res=", res)

    }

* 当左右两边都是数值型时，则做法加法运算
* 当左右两边都是字符串，则做字符串拼接

## 5. 数据类型

### （1）整数类型

    package main

    // import "fmt"
    // import "unsafe"

    import (
        "fmt"
        "unsafe"
    )

    // 演示golang中整数类型使用

    func main()  {

        var i int = 1
        fmt. Println("i=", i)

        

        // 测试一下in8的范围 -128~127
        // 其他的 int16, int32, int64, 类推。。。
        var j int8 = 127
        fmt. Println("j=", j)

        // 测试一下unit8的范围， 其他的uint16, uint32, unit64一样
        var k uint8 = 255
        fmt.Println("k=", k)

        // int, unint, rune, byte的使用
        var a int = 8900
        fmt.Println("a=", a)
        var b uint = 1
        var c byte = 255
        fmt.Println("b=", b, "c=", c)

        // 整型的使用细节
        var n1 = 100 // ? n1 是什么类型
        // 在这里我们介绍一下如何查看某个变量的数据类型
        // fmt.Print()可以用于做格式化输出
        fmt.Printf("n1 的类型%T", n1)

        // 如何在程序查看某个变量的占用字节大小和数据类型（使用较多）
        var n2 int64 = 10
        // unsage.Sizeof(n2) 是unsafe包的一个函数，可以返回n1变量占用的字节数
        fmt.Printf("n2 的类型%T, n2占用的字节数是%d \n", n2, unsafe.Sizeof(n2))

        // Golang程序中整型变量在使用时，遵守保小不保大的原则，即：在保证程序正确运行下，尽量使用占用空间小的数据类型[如：年龄]
        var age byte = 90
        fmt. Printf("age= %d", age)

    }

* Golang个整数类型分：有符号和无符号，int uint的大小和系统有关
* Golang的整型默认声明为int型
* 如何在程序查看摸个变量的字节大小和数据类型
* Golang程序中整型变量在使用时，遵守保小不保大的原则，即：在保证程序正确运行下，尽量使用占用空间小的数据类型[如：年龄]
* bit：计算机中的最小存储单位，byte：计算机中基本存储单元[二进制在详细说] 1 byte = 8 bit

### （2）浮点类型

    package main

    import (
        "fmt"
    )

    // 演示golang中小数类型使用

    func main()  {

        var price float32 = 89.12
        fmt. Println("price=", price)
        var num1 float32 = -0.00089
        var num2 float32 = -7809656.09
        fmt. Println("num1=", num1, "num2=", num2)

        // 尾数部分可能丢失，造成精度损失。-123.0000901
        var num3 float32 = -123.0000901
        var num4 float64 = -123.0000901
        fmt.Println("num3=", num3, "num4=", num4)

        // Golang的浮点型默认声明为float64类型
        var num5 = 1.1
        fmt.Printf("num5的数据类型是%T \n", num5)

        // 十进制数形式，如： 5.12    .512（必须有小数点）
        num6 := 5.12
        num7 := .123 // => 0.123
        fmt. Println("num6=", num6, "num7=", num7)

        

        // 科学计数法形式
        num8 := 5.1234e2 // 5.1234 * 10的2次方
        num9 := 5.1234E2 // 5.1234 * 10的2次方
        num10 := 5.1234E-2 // 5.1234 * 10的-2次方
        fmt. Println("num8=", num8, "num9=", num9, "num10=", num10)

        

    }

### （3）字符类型：char

    package main

    import (
        "fmt"
    )

    // 演示golang中字符类型使用

    func main() {

        var c1 byte = 'a'
        var c2 byte = '0' // 字符的0

        // 当我们直接输出byte值，就是输出了对应字符的码值
        fmt.Println("c1=", c1)
        fmt.Println("c2=", c2)
        // 如果我们希望输出对应字符串，需要使用格式化输出
        fmt.Printf("c1=%c c2=%c \n", c1, c2)

        // var c3 byte = '北' // overflow溢出
        var c3 int = '北' // overflow溢出
        fmt.Printf("c3=%c c3对应码值=%d \n", c3, c3)

        // 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出改数字对应的unicode字符
        var c4 int = 120 // 22269 -> "国" 120 -> "x"
        fmt.Printf("c4=%c \n", c4)

        // 字符类型是可以进行运算的，相当于一个整数，因为他都有对应的Unicode码
        var n1 = 10 + 'a' // 10 + 97 = 107
        fmt. Println("n1=", n1)

    }

代码说明：

* 如果我们保存的字符在ASCII表中， 比如[0-1, a-z, A-Z]直接可以保存到byte
* 如果我们保存的字符对应码值大于255，这时我们可以考虑使用int类型保存
* 如果我们需要安装字符的方式输出，这时我们需要格式化输出，即fmt. Printf("%c", c1)

字符类型使用细节：

* 字符常量是使用单引号('')括起来的单个字符。例如：var c1 byte = 'a' var c2 int = '中' var c3 byte = '9'
* Go中允许使用转义字符'\'来将其后的字符转换为特殊字符型常量。例如，var c3 char = '\n' // '\n'表示换行符
* Go语言的字符使用UTF-8编码

    英文字母-1个字节 汉字-3个字节

* 在GO中，字符的本质是一个整数，直接输出是，是该字符对应的UTF-8编码的码值
* 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出改数字对应的unicode字符
* 字符类型是可以进行运算的，相当于一个整数，因为他都有对应的Unicode码

字符类型本质探讨：

* 字符类型存储到计算机中， 需要将字符类型对应的码值找出来（整数）找出来

    存储： 字符 ---> 对应码值 ---> 二进制 ----> 存储
    读取： 二进制 ---> 码值 ---> 字符 ---> 读取

* 字符和码值的对应关系是通过字符编码表决定的（是规定好）
* Go语言的编码都统一成了utf-8，非常的方便，很统一，再也没有编码乱码的困扰

### （4）布尔类型：bool

    package main

    import (
        "fmt"
        "unsafe"
    )

    func main()  {

        var b = false
        fmt. Println("b=", b)
        // 注意事项
        // 1.bool类型占用存储空间是1个字节
        fmt. Printf("b占用存储空间是：%d", unsafe. Sizeof(b))
        // 2.bool类型只能取值true或者false

    }

* 布尔类型也叫bool类型，bool类型数据只允许取值true和false
* bool类型占1个字节
* bool类型适用于逻辑运算，一般用于程序流程控制[注：这个后面会详细介绍]

    1> if 条件控制语句:
    2> for 循环控制语句:

### （5）字符串类型：string

    package main

    import "fmt"

    func main()  {

        

        // string的基本使用
        var address string = "北京长城 110 hello world"
        fmt. Println("address=", address)

        // 字符穿一旦赋值了，字符串就不能修改了，在Go语言中字符串是不可变的
        // var str string = 'hello' // invalid character literal (more than one character)go
        // str[0] = "a" // cannot assign to str[0]go
        // fmt.Println("str=", str)

        // 字符串的两种表示形式
        // 1> 双引号， 会识别转义字符
        str2 := "abc\nabc"
        fmt.Println(str2)
        // 2> 反引号， 以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
        str3 := `* Go语言的字符串使用UTF-8编码标识Unicode文本，这样Golang统一使用UTF-8编码

        * 字符穿一旦赋值了，字符串就不能修改了，在Go语言中字符串是不可变的
        * 字符串的两种表示形式`

        fmt.Println(str3)

        var str = "hello " + "world"
        str += " hahaha"
        fmt. Println(str)
        // 当一个拼接的操作很长时，怎么办，可以分行写， 需要将'+'保留在上一行
        var str4 = "hello " + "world" + "hello " + 
        "world" + "hello " + "world" + "hello " + 
        "world" + "hello " + "world"
        fmt. Println(str4)

    }

* Go语言的字符串使用UTF-8编码标识Unicode文本，这样Golang统一使用UTF-8编码
* 字符穿一旦赋值了，字符串就不能修改了，在Go语言中字符串是不可变的
* 字符串的两种表示形式

    1> 双引号， 会识别转义字符
    2> 反引号， 以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果

* 字符串拼接方式
* 当一个拼接的操作很长时，怎么办，可以分行写

### （5）基本数据类型相互转换

    package main

    import "fmt"

    func main()  {

        var i int32 = 100
        // 希望将 i --> float
        var n1 float32 = float32(i)
        fmt. Println(n1)
        var n2 int8 = int8(i)
        fmt. Println(n2)
        var n3 int64 = int64(i) // 低精度 -> 高精度
        fmt. Println(n3)

        // 在转换中，比如将 int8 【-128 --- 127】， 编译时不会报错
        // 只是转换的结果是按溢出处理， 和我们希望的结果不一样
        var num1 int64 = 999999
        var num2 int8 = int8(num1)
        fmt. Println(num2)

    }

### （5）基本数据类型和string的转换

方式1：fmt. Sprintf("%参数", 表达式)

* 参数需要和表达式的数据类型相匹配
* fmt. Sprintf() 会返回转换后的字符串

方式2：使用strconv包的函数

案例演示:

    package main

    import (
        "fmt"
        "strconv"
    )

    // 演示golang中基本数据联系转成string使用

    func main()  {

        var num1 int = 99
        var num2 float64 = 23.456
        var b bool = true
        var myChar byte = 'h'
        var str string // 空的str

        

        // 使用第一种方式来转换fmt. Sprintf方法
        str = fmt. Sprintf("%d", num1)
        fmt. Printf("str type %T str=%q \n", str, str)

        

        str = fmt. Sprintf("%f", num2)
        fmt. Printf("str type %T str=%q \n", str, str)

        str = fmt.Sprintf("%t", b)
        fmt.Printf("str type %T str=%q \n", str, str)

        str = fmt.Sprintf("%c", myChar)
        fmt.Printf("str type %T str=%q \n", str, str)

        // 第二种方式 strconv 包的函数
        var num3 int = 99
        var num4 float64 = 23.456
        var b2 bool = true
        str = strconv. FormatInt(int64(num3), 10)
        fmt. Printf("str type %T str=%q \n", str, str)
        // strconv. FormatFloat(num4, 'f', 10, 64)
        // 说明： f：格式 10：表示小数位保留10位 64：表示这个小数是float64
        str = strconv. FormatFloat(num4, 'f', 10, 64)
        fmt. Printf("str type %T str=%q \n", str, str)

        

        str = strconv. FormatBool(b2)
        fmt. Printf("str type %T str=%q \n", str, str)

        // strconv包中有一个函数Itoa
        var num5 int64 = 4567
        str = strconv. Itoa(int(num5))
        fmt. Printf("str type %T str=%q \n", str, str)

    }

### （6）string转基本数据类型的

    package main

    import (
        "fmt"
        "strconv"
    )

    // 演示go中string转成基本数据类型
    func main(){
        str := "true" 
        var b bool
        // b, _ = strconv. ParseBool(str)
        // 说明
        // 1.strconv. ParseBool(str) 函数会返回两个值(value bool, err error)
        // 因为只想获取到value bool, 不想获取err 所以使用 _ 忽略
        b, _ = strconv. ParseBool(str)
        fmt. Printf("b type %T b = %v \n", b, b)

        var str2 string = "123456789"
        var n1 int64
        n1, _ = strconv. ParseInt(str2, 10, 64)
        n2 := int(n1)
        fmt. Printf("n1 type %T n1=%v \n", n1, n1)
        fmt. Printf("n2 type %T n2=%v \n", n2, n2)

        

        var str3 string = "123.456"
        var f1 float64
        f1, _ = strconv. ParseFloat(str3, 64)
        f2 := float32(f1)
        fmt. Printf("f1 type %T f1=%v \n", f1, f1)
        fmt. Printf("f2 type %T f2=%v \n", f2, f2)

        // 注意：
        var str4 string = "hello"
        var n3 int64 = 11
        n3, _ = strconv. ParseInt(str4, 10, 64)
        fmt. Printf("n3 type %T n3=%v", n3, n3)

    }

# 指针

## 1. 基本介绍

    package main

    import (
        "fmt"
    )

    // 演示golang中指针的类型

    func main() {

        // 基本数据类型在内存布局
        var i int = 10
        // i 的地址是什么，&i
        fmt. Println("i的地址=", &i)

        // 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
        // var ptr *int = &num
        // 1.ptr 是一个指针变量
        // 2.ptr的类型 *int
        // 3.ptr本身的值&i
        var ptr *int = &i
        fmt.Printf("ptr=%v \n", ptr)
        fmt.Printf("ptr的地址是%v \n", &ptr)
        fmt.Printf("ptr的指向是%v \n", *ptr)
        fmt.Printf("ptr的指向是%v \n", *ptr)

    }

* 基本数据类型，变量存的就是值，也叫值类型
* 获取变量的地址，用&，比如：var num int，获取num的地址：&num
* 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值

    比如： var ptr *int = &num

* 获取指针类型所指向的值，使用 `*` ，比如： `var *ptr int` , 使用 *ptr获取p指向的值
* 图解：

    0xc000006030 ---------> 0xc0000120b8 ---------> ptr ---------> 10 ---------> i

## 2. 指针细节说明

* 值类型，都有对应的指针类型，形式为 `*数据类型` ，比如int的对应的指针就是 `*int` , float32对应的指针类型就是 `*float32` ，依次类推。
* 值类型包括：基本数据类型 `int系列，float系列，bool、数组和结构体struct`
# 值类型和引用类型

* 值类型：基本数据类型 `int系列，float系列，bool，string、数组和结构体struct`
* 引用类型： `指针、slice切片、map、管道chan、interface` 等都是引用类型

## 值类型和引用类型使用特点

1）值类型：变量直接存储值，内存通常在栈中分配
2）引用类型：变量存储的是一个地址，在这个地址对应的空间才真正存储数据(值)，内存通常在堆上分配，当没有任何变量引用这个地址时，改地址对应的数据空间就成为一个垃圾，由GC来回收

# 标识符的命名规范

* 包名：保持package的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，不要和标准库有冲突
* 变量名、函数名、常量名：采用驼峰法
* 如果变量名、函数名、常量名首字母大写，则可以被其他的报访问；如果首字母小写，则只能在本包中使用（ `注` ：可以简单的理解成，首字母大写是共有的，首字母小写是私有的）

案例演示：

工具包：

    package model

    var Heroname string = "宋江"

引用包：

    package main

    import (
        "fmt"
        "go_code/project02/model"
    )

    func main() {

        fmt. Println(model. Heroname)

    }

# Golang的运算符

## 1. 算数运算符

    package main

    import "fmt"

    func main() {

        // 重点讲解 /  %
        // 说明，如果运算的数都是整数，运算的结果是小数，去掉小数部分，保留整数部分
        fmt. Println(10 / 4)

        

        var n1 float32 = 10 / 4
        fmt. Println(n1)

        

        // 说明，如果需要保留小数部分
        var n2 float32 = 10.0 / 4
        fmt. Println(n2)

        

        // 演示%
        // 看一个公式：a % b = a - a / b * b
        fmt. Println("10%3=", 10 % 3)
        fmt. Println("-10%3=", -10 % 3)
        fmt. Println("10%-3=", 10 % -3)
        fmt. Println("-10%-3=", -10 % -3)

    }

* `+` --------> `正号、加号、拼接` ------> `+3; 5+5; "he" + "llo"`
* `-` --------> `负号、减号` ------> `-4; 6-4`
* `*` --------> `乘` ------> `4*3`
* `/` --------> `除` ------> `5/5`
* `%` --------> `取余` ------> `7%5`
* `++` --------> `自增` ------> `a=2 a++`
* `--` --------> `自减` ------> `a=2 a--`
细节说明

    package main

    import (
        "fmt"
    )

    func main() {

        var i = 8
        var a int
        // a = i++ // 错误，i++只能独立使用
        // a = i-- // 错误，i--只能独立使用
        i++
        // ++i // 错误，在golang中没有++i
        // --i // 错误，在golang中没有--i

    }

## 2. 关系运算符(比较运算符)

    package main

    import "fmt"

    func main() {

        // 演示关系运算符
        var n1 int = 9
        var n2 int = 8
        fmt. Println(n1 == n2)
        fmt. Println(n1 != n2)
        fmt. Println(n1 > n2)
        fmt. Println(n1 < n2)
        fmt. Println(n1 >= n2)
        fmt. Println(n1 <= n2)
        flag := n1 > n2
        fmt. Printf("flag 的type 是 %T flag:%s", flag, flag)

    }

* `==` --------> `相等于` ------> `4==3` -------> `false`
* `!=` --------> `不等于` ------> `4!=3` -------> `true`
* `<` ---------> `小于` ------> `4<3` -------> `false`
* `>` ---------> `大于` ------> `5>5` -------> `false`
* `<=` --------> `小于等于` ------> `7<=5` -------> `false`
* `>=` --------> `大于等于` ------> `7>=5` -------> `true`

## 2. 逻辑运算符

* `&&` 与
* `||` 或
* `!` 非

注意事项和细节说明

* && 也叫短路与：如果第一个条件为false，则第二个条件不会判断，最终结果为false

* || 也叫短路或，如果第一个条件为true， 则第二个条件不会判断，最终结果为true

    package main

    import "fmt"

    func test() bool {
        fmt. Println("test...")
        return true

    }

    func main() {

        var i int = 10
        if i < 9 && test() {
            fmt. Print("ok")

        }

        if i > 9 || test() {
            fmt. Println("hello")

        }

    }

## 3. 赋值运算符

* `=` --------> `简单的赋值运算` ------> `c=a+b`
* `+=` --------> `相加后在赋值` ------> `c+=b`
* `-=` --------> `相减后再赋值` ------> `c-=b`
* `*=` --------> `相乘后在赋值` ------> `c*=b`
* `/=` --------> `相除后再赋值` ------> `c/=b`
* `<<=` --------> `左移后赋值` ------> `c%=b`
* `>>=` --------> `右移后赋值` ------> `c%=b`
* `&=` --------> `换位与后赋值` ------> `c%=b`
* `^=` --------> `换位异或后赋值` ------> `c%=b`
* `|=` --------> `换位或后赋值` ------> `c%=b`
案例演示

    package main

    import "fmt"

    func main() {

        // 赋值运算符的使用演示
        // var i int
        // i = 10 // 基本赋值

        // 有两个变量， a和b， 要求将其进行交换，最终打印结果
        // a = 9, b = 2 ====> a = 2  b = 9
        a := 9
        b := 2
        fmt.Printf("交换前的情况是 a = %v, b = %v \n", a, b)
        // 定义一个临时变量
        t := a
        a = b 
        b = t
        fmt.Printf("交换后的情况是 a = %v, b = %v \n", a, b)

        a += 17
        fmt. Println("a=", a)

    }

## 4. 三元运算符

    package main

    import "fmt"

    func main() {

        // 三元运算符
        var n int
        var i int = 10
        var j int = 10

        // 传统的三元运算
        // n = i > j ? i : j
        if i > j {
            n = i
        } else {
            n = j

        }

        fmt. Println("n=", n)

    }

特别说明：Golang不支持三元运算符

# 键盘输入语句

## (1)fmt. Scanln函数

    package main

    import "fmt"

    func main() {

        // 要求：可以从控制台接收用户信息 
        // 方式1  fmt. Scanln
        // 先声明需要的变量
        var name string
        var age byte
        var sal float32
        var isPass bool
        fmt. Println("请输入姓名 ")
        fmt. Scanln(&name)
        fmt. Println("请输入年龄 ")
        fmt. Scanln(&age)
        fmt. Println("请输入薪资 ")
        fmt. Scanln(&sal)
        fmt. Println("请输入是否通过考试")
        fmt. Scanln(&isPass)

        fmt. Printf("名字是 %v \n年龄是 %v \n薪水是 %v \n是否通过考试 %v \n", name, age, sal, isPass)

    }

## (2)Scanf函数

    package main

    import "fmt"

    func main() {

        // 方式2  fmt.Scanf
        fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开")
        fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
        fmt.Printf("名字是 %v \n年龄是 %v \n薪水是 %v \n是否通过考试 %v \n", name, age, sal, isPass)

    }

# 分支控制

## 1.if分支

### (1)单分支

    跳过

### (2)多分支

    跳过

### (3)单分支

    跳过

## 2.switch分支结构

### (1)基本介绍

* Go 编程语言中 switch 语句的语法如下：

    switch var1 {
        case val1:
            ...
        case val2:
            ...
        default:
            ...

    }

* 变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
* switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
* switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
* switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。
* switch 可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。

    switch{
        case 1, 2, 3, 4:
        default:

    }

### (2)快速入门案例

代码：

    package main

    import "fmt"

    func main() {

    /* 定义局部变量 */
    var grade string = "B"
    var marks int = 90

    switch marks {

        case 90: grade = "A"
        case 80: grade = "B"
        case 50,60,70 : grade = "C"
        default: grade = "D"  

    }

    switch {

        case grade == "A" :
            fmt.Printf("优秀!\n" )    
        case grade == "B", grade == "C" :
            fmt.Printf("良好\n" )      
        case grade == "D" :
            fmt.Printf("及格\n" )      
        case grade == "F":
            fmt.Printf("不及格\n" )
        default:
            fmt.Printf("差\n" );

    }

    fmt. Printf("你的等级是 %s\n", grade ); 

    }

以上代码执行结果为：

    优秀!
    你的等级是 A

### (3)细节讨论

    package main

    import "fmt"

    func main() {

        var a byte
        fmt. Println("请输入字符：")
        fmt. Scanf("%c", &a)
        switch a {
            case 'a':
                fmt. Println("周一")
            case 'b':
                fmt. Println("周二")
            case 'c':
                fmt. Println("周三")
            case 'd', 'e':
                fmt. Println("周四")
            default:
                fmt. Println("有误")

        }

        switch age := 10; {
            case age == 10:
                fmt.Println("age == 10")
            case age == 20:
                fmt.Println("age == 20")
            case age == 30:
                fmt.Println("age == 30")
            default:
                fmt.Println("没有匹配到")

        }

        switch {
            case false:
                    fmt. Println("1、case 条件语句为 false")
                    fallthrough
            case true:
                    fmt. Println("2、case 条件语句为 true")
                    fallthrough
            case false:
                    fmt. Println("3、case 条件语句为 false")
                    fallthrough
            case true:
                    fmt. Println("4、case 条件语句为 true")
            case false:
                    fmt. Println("5、case 条件语句为 false")
                    fallthrough
            default:
                    fmt. Println("6、默认 case")

        }

    }

* switch后也可以不带表达式，类似if -- else分支来使用
* switch后也可以直接声明定义一个变量分号结束， 【不推荐】

# Go语言函数

函数是基本的代码块，用于执行一个任务。

Go 语言最少有个 main() 函数。

你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。

函数声明告诉了编译器函数的名称，返回类型，和参数。

Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接收不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

## 1. 函数定义

Go 语言函数定义格式如下：

    func function_name( [parameter list] ) [return_types] {
        函数体

    }

函数定义解析：

* func：函数由 func 开始声明
* function_name：函数名称，函数名和参数列表一起构成了函数签名。
* parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
* return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
* 函数体：函数定义的代码集合

## 2. 实例

以下实例为 max() 函数的代码，该函数传入两个整型参数 num1 和 num2，并返回这两个参数的最大值：

    /* 函数返回两个数的最大值 */
    func max(num1, num2 int) int {
        /* 声明局部变量 */
        var result int

        if (num1 > num2) {
            result = num1
        } else {
            result = num2

        }

        return result

    }

## 3. 函数调用

当创建函数时，你定义了函数需要做什么，通过调用该函数来执行指定任务。

调用函数，向函数传递参数，并返回值，例如：

    package main

    import "fmt"

    func main() {

        /* 定义局部变量 */
        var a int = 100
        var b int = 200
        var ret int

        /* 调用函数并返回最大值 */
        ret = max(a, b)

        fmt. Printf( "最大值是 : %d\n", ret )

    }

    /* 函数返回两个数的最大值 */
    func max(num1, num2 int) int {
        /* 定义局部变量 */
        var result int

        if (num1 > num2) {
            result = num1
        } else {
            result = num2

        }

        return result

    }

以上实例在 main() 函数中调用 max（）函数，执行结果为：

    最大值是 : 200

## 4. 函数返回多个值

Go 函数可以返回多个值，例如：

    package main

    import "fmt"

    func swap(x, y string) (string, string) {
        return y, x

    }

    func main() {

        a, b := swap("Google", "Runoob")
        fmt. Println(a, b)

    }

以上实例执行结果为：

    Runoob Google

## 5. 函数参数

函数如果使用参数，该变量可称为函数的形参。
形参就像定义在函数体内的局部变量。
调用函数，可以通过两种方式来传递参数：

### (1)值传递

值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。

    package main

    import "fmt"

    func main() {

    /* 定义局部变量 */
    var a int = 100
    var b int = 200

    fmt.Printf("交换前 a 的值为 : %d\n", a )
    fmt.Printf("交换前 b 的值为 : %d\n", b )

    /* 通过调用函数来交换值 */
    swap(a, b)

    fmt. Printf("交换后 a 的值 : %d\n", a )
    fmt. Printf("交换后 b 的值 : %d\n", b )

    }

    /* 定义相互交换值的函数 */
    func swap(x, y int) int {
    var temp int

    temp = x /* 保存 x 的值 */
    x = y    /* 将 y 值赋给 x */
    y = temp /* 将 temp 值赋给 y*/

    return temp; 

    }

以下代码执行结果为：

    交换前 a 的值为 : 100
    交换前 b 的值为 : 200
    交换后 a 的值 : 100
    交换后 b 的值 : 200

### (2)引用传递

引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

    package main

    import "fmt"

    func main() {

    /* 定义局部变量 */
    var a int = 100
    var b int= 200

    fmt.Printf("交换前，a 的值 : %d\n", a )
    fmt.Printf("交换前，b 的值 : %d\n", b )

    /* 调用 swap() 函数

    - &a 指向 a 指针，a 变量的地址
    - &b 指向 b 指针，b 变量的地址

    */
    swap(&a, &b)

    fmt. Printf("交换后，a 的值 : %d\n", a )
    fmt. Printf("交换后，b 的值 : %d\n", b )

    }

    func swap(x *int, y *int) {
    var temp int
    temp = *x    /* 保存 x 地址上的值 */
    *x = *y      /* 将 y 值赋给 x */
    *y = temp    /* 将 temp 值赋给 y */

    }

以上代码执行结果为：

    交换前，a 的值 : 100
    交换前，b 的值 : 200
    交换后，a 的值 : 200
    交换后，b 的值 : 100

`注意` ：默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

## 6. 函数用法

### (1)函数作为另外一个函数的实参

函数定义后可作为另外一个函数的实参数传入
Go 语言可以很灵活的创建函数，并作为另外一个函数的实参。以下实例中我们在定义的函数中初始化一个变量，该函数仅仅是为了使用内置函数 math.sqrt()，实例为：

    package main

    import (
        "fmt"
        "math"
    )

    func main(){
        /* 声明函数变量 */
        getSquareRoot := func(x float64) float64 {

            return math. Sqrt(x)

        }

        /* 使用函数 */
        fmt. Println(getSquareRoot(9))

    }

### (2)闭包

闭包是匿名函数，可在动态编程中使用

Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量，代码如下：

    package main

    import "fmt"

    func getSequence() func() int {
        i:=0
        return func() int {
            i+=1
            return i  

        }

    }

    func main(){
        /* nextNumber 为一个函数，函数 i 为 0 */
        nextNumber := getSequence()  

        /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
        fmt. Println(nextNumber())
        fmt. Println(nextNumber())
        fmt. Println(nextNumber())

        

        /* 创建新的函数 nextNumber1，并查看结果 */
        nextNumber1 := getSequence()  
        fmt. Println(nextNumber1())
        fmt. Println(nextNumber1())

    }

### (3)方法

方法就是一个包含了接受者的函数

Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：

    func (variable_name variable_data_type) function_name() [return_type]{
        /* 函数体*/

    }

实例:

    package main

    import (
        "fmt"  
    )

    /* 定义结构体 */
    type Circle struct {
     radius float64

    }

    func main() {

        var c1 Circle
        c1.radius = 10.00
        fmt. Println("圆的面积 = ", c1.getArea())

    }

    //该 method 属于 Circle 类型对象中的方法
    func (c Circle) getArea() float64 {
        //c.radius 即为 Circle 类型对象中的属性
        return 3.14 * c.radius * c.radius

    }

以上代码执行结果为：

    圆的面积 =  314
