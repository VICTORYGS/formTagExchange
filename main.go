package main

import (
	"bufio"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func main() {
	println("\n-----------------  表单元素 转换说明  ------------------\n")
	println("□与=表示input checkbox标签")
	println("_表示input text标签")
	println("如文字中包含上面三种符号请在前面加上\\")
	println("元素在不同的标签,请用空格隔开,反之如2个元素要放在一个标签里面, 两者之间请不要有空格")

	println("内容示例：意识状态 =清楚 =意识模糊 =嗜睡 =昏迷_")
	println("\n-------------  请在下面粘贴内容然后回车  --------------")

	input:=bufio.NewScanner(os.Stdin)
	for input.Scan(){
		res:=input.Text()
		println("--------------------结果---------------------")
		r:=exchange(res)
		clipboard.WriteAll(r)
		println(r)
		println("------------以上结果已复制到剪切面板------------")
	}
}

func exchange(str string)(res string ) {
	//占位符
	z:="qcvg97llgllkdama5rd4ad5a4d4s0145as56da504d"//□与=
	z1:="k5f4sdjfsfllfksdkljfklsjfppw445auuqad577dg"//_
	arr:=strings.Split(str," ")
	arr1:=[]string{}

	for _,s:=range arr{
		res1:=s
		res1=strings.Replace(res1,`\□`,z,-1)
		res1=strings.Replace(res1,`\=`,z,-1)
		res1=strings.Replace(res1,`\_`,z1,-1)
		res1=strings.Replace(res1,`=`,`<input type="checkbox" />`,-1)
		res1=strings.Replace(res1,`□`,`<input type="checkbox" />`,-1)
		res1=strings.Replace(res1,`_`,`<input type="text" />`,-1)

		res1=strings.Replace(res1,z,`□`,-1)
		res1=strings.Replace(res1,z,`=`,-1)
		res1=strings.Replace(res1,z1,`_`,-1)
		//println(res1)
		arr1=append(arr1,res1)
	}

	for _,v:=range arr1{
		res+="<span>"+v+"</span>\n"
	}
	return
}
