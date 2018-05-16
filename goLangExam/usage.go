package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var num1 int = 32
	var num2 float32 = 0.1
	var s1 string = "Hello, world!\n"
	var b1 bool = true
	const x, y int = 30, 50	// 병렬 할당

	const (
		Sunday	= iota // 0부터 순서대로 생성
		Monday
	)

	i := 10

	if i>=5 {
		fmt.Println("5 이상")
	}

	for i := 0; i<5;i++ { ///// 같은 라인에 중괄호
		fmt.Println(i)
	}

	fmt.Println(num1, num2, s1, b1);
	// comment in one line
	
	/* comment
	   comment
	   comment */
	
	Loop:
		for i:=0;i<3;i++ {
			for j:=0;j<3;j++ {
				if j==2 {
					break Loop	// or continue Loop
				}
			}
		}
		fmt.Println("inside Loop")

	switch i {
		case 0:
			fmt.Println(num1)
		case 1:
			fmt.Println(num2)
	}
}
