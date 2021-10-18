package main

import "fmt"

func main()  {
	var res = []string{}

	for i := 1; i <= 15; i++ {
		var temp string
		if i % 3 == 0 && i % 5 == 0 {
			temp = "FizzBuzz"
		} else if i % 3 == 0 {
			temp = "Fizz"
		} else if i % 5 == 0 {
			temp = "Buzz"
		} else {
			temp = fmt.Sprintf("%+v", i)
		}
		res = append(res, temp)
	}

	fmt.Println(res)
}