package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CheckName(Name string) bool {
	for _, r := range Name {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println("Welcome to my test!")
	fmt.Printf("Please type your name:")
	var Name string
	for {
		reader := bufio.NewReader(os.Stdin)
		Name, _ = reader.ReadString('\n')
		Name = strings.TrimSpace(Name)
		checked := CheckName(Name)
		if checked == false {
			fmt.Printf("Please don't use numbers, type in English and no spaces.")
		} else {

			break
		}
	}

	fmt.Printf("Welcome %v :), hope you enjoy this test!\n", Name)
	time.Sleep(2 * time.Second)
	fmt.Printf("Enter your age:")
	var (
		age    string
		ageOld int
	)
	score := 0
	num_questions := 3
	reader := bufio.NewReader(os.Stdin)
	age, _ = reader.ReadString('\n')
	age = strings.TrimSpace(age)
	ageOld, _ = strconv.Atoi(age)
	if ageOld < 17 {
		fmt.Printf("Can't do this test yet!\n")
		return
	} else {
		fmt.Println("Time for your test....")
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Write the options given as an answer")
	time.Sleep(2 * time.Second)
	fmt.Println("Question 1 : PC or Consoles?")
	var answer string
	answer, _ = reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	matched, _ := regexp.MatchString(`(?i)Consoles`, answer)
	matched2, _ := regexp.MatchString(`(?i)PC`, answer)
	if matched {
		fmt.Println("Lol wrong :)")
	} else if matched2 {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Didn't write one of the choices at all!")
	}
	fmt.Println("Question 2: Diablo 2 or Diablo 3?")
	answer, _ = reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	matched3, _ := regexp.MatchString(`(?i)Diablo\s?2`, answer)
	matched4, _ := regexp.MatchString(`(?i)Diablo\s?3`, answer)
	if matched3 {
		fmt.Println("Stay a while and listen!")
		score++
	} else if matched4 {
		fmt.Println("Wrong.")
	} else {
		fmt.Println("Did not type the options given!")
	}
	fmt.Println("Question 3: 9 + 10?")
	answer, _ = reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	matched5, _ := regexp.MatchString(`19`, answer)
	matched6, _ := regexp.MatchString(`21`, answer)
	if matched5 {
		fmt.Println("Excellent Mathematical skills.")
		score++
	} else if matched6 {
		fmt.Println("You lack crucial mathematic knowledge...........")
	} else {
		fmt.Println("You lack crucial mathematic knowledge...........")
	}
	fmt.Printf("You scored %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	respercent := math.Round(percent)
	fmt.Printf("You Scored: %v%%.\n", respercent)
}
