package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")	
	}

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday: 
			fmt.Println("It is weekend")

		default:
			fmt.Println("It is weekday")
	}

	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Before noon")

		default:
			fmt.Println("After noon")
	}

	whatAmI := func (i interface{}) {
		switch t := i.(type) {
			case bool: 
				fmt.Println("I am bool")
			case int:
				fmt.Println("I am int")
			default:
				fmt.Printf("don't know type %T \n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hi!")
}
