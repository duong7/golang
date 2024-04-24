package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")
	//var name string = "Tim"
	//var name int = -4
	//var name float64 =2.144444
	//var name bool = false
	//name = "hello"
	//fmt.Println(name)

	//name := "Tim"
	//age := 21
	//fmt.Printf("Hello %v, you are %v", name, age)
	fmt.Printf("Enter\n your name:")
	// var age int
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v welcome to the game!\n", name)
	fmt.Printf("Enter your age:")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You cannot play!")
		return
	}
	score := 0
	num_questions := 2
	fmt.Println("What is the best means of transportation to work far away?, the car or the subway? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	if answer+" "+answer2 == "the subway" || answer+" "+answer2 == "THE SUBWAY" || answer+" "+answer2 == "the SUBWAY" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Incorrect!")
	}
	fmt.Println("2+2= ?")
	var as int
	fmt.Scan(&as)
	if as == 4 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")

	}
	fmt.Printf("You scored %v out of %v\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored:%v%%", percent)
}
