package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	ifFile1()
	ifFile2()
	//println(switchEval(1, 2, "-1"))
	print(switch2(101))

}

func ifFile1(){
	const filename string  ="abc.txt"
	content ,err :=ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n", content)
	}
}
//简化写法
func ifFile2(){
	const filename="abc.txt"
	if c,er :=ioutil.ReadFile(filename);er!=nil{
		fmt.Println(er)
	}else{
		fmt.Printf("%q\n",c)
	}
}

//switch 普通用法
func switchEval(a,b int,op string) int{
	var res int
	switch op {
	case "+":
		res=a+b
	case "-":
		res=a-b
	case "*":
		res=a*b
	case "/":
		res=a/b
	default:
		panic("unsupport opt"+op)
	}

	return res
}
func switch2(score int) string {
	g:=""
	switch  {
	case score<0||score>100:
		panic(fmt.Sprintf("wrong score %d",score))
	case score<60:
		g="F"
	case score <80:
		g="c"
	case score <90:
		g="b"
	case score <=100:
		g="a"

	}
	return g
}