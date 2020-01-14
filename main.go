package main

import "fmt"

func main()  {
	fmt.Println("Hello Worlds")
	res := obaid("hi")
	fmt.Println(res)

	z:=abc(2,3)
	fmt.Println(z)

	x:=plus(1,2,3,4 )
	fmt.Println("this is function return ",x)


}

func plus(a int ,b int,c int , d int )int{
	return a+b+c+d

}
func obaid(a string) string {

	fmt.Println("obaid uh are in func now")
	return a + "Obaid "


}

func abc(a int,b int )int{
	c:=  a+b
	return c
}

