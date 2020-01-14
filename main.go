package main

import "fmt"

func main()  {
	fmt.Println("Hello Worlds")
	res := obaid("hi")
	fmt.Println(res)

	z:=abc(2,3)
	fmt.Println(z)

}

func obaid(a string) string {

	fmt.Println("obaid uh are in func now")
	return a + "Obaid "


}

func abc(a int,b int )int{
	c:=  a+b
	return c
}

