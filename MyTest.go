package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Function that checks if the name matches the parameters set. It has to be English, no spaces, no numbers, no empty spaces.

func Userinput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func Checkname(input string) bool {
	for _, r := range input {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// Function that reads your name and checks it with the checkname function
func Uname(name string) string {
	checked := Checkname(name)
	if checked == false {
		return ""
	} else if len(name) == 0 {
		return ""
	} else {
		return name
	}
}

// Function that checks if age has numbers or something else. If age has white space or letters, loop.
func LetterAge(age string) string {
	IsLetter := regexp.MustCompile(`[a-zA-Z_\s\W]+`).MatchString(age)
	if len(age) == 0 {
		return ""
	} else if IsLetter == true {
		return ""
	} else {
		return age
	}
}

// Function that age checks. If age check fails, exit program.
func Uage(age string) bool {
	var ageOld int
	ageOld, _ = strconv.Atoi(age)
	if ageOld < 17 {
		return false
	} else {
		return true
	}
}

// Score check and a variable with the amount of questions. Throughout the questions, it will keep a count of correct answers. Tracks how many questions there is as well.
func Unum_questions() int {
	var num_questions int
	db, err := sql.Open("mysql", "DuckUser:Test123456@tcp(127.0.0.1:3306)/DuckDB") // Change the DB connection settings for your own machine.
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT COUNT(*) FROM Questions")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&num_questions)
		if err != nil {
			log.Fatal(err)
		}
	}
	return num_questions
}

// In this section, we use the DB to select the questions the user has set in the DB file, with their own correct answers. DB and program have a regexp check that allows whitespaces and case insensitive answers
func Test() int {
	var (
		question  string
		answer    string
		coranswer string
		matched   bool
		score     int
	)
	reader := bufio.NewReader(os.Stdin)
	db, err := sql.Open("mysql", "DuckUser:Test123456@tcp(127.0.0.1:3306)/DuckDB") // Change the DB connection settings for your own machine.
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select Question, `Correct Answer` from Questions")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&question, &coranswer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(question)
		answer, _ = reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		matched, _ = regexp.MatchString(`(?i)`+coranswer, answer)
		if matched {
			fmt.Println("Correct")
			score++
		} else {
			fmt.Println("Wrong")
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return score
}

// At last, the program counts how many questions were correct, and the percentages the user got. Adds it to a separate result table for tracking
func Result(score int, num_questions int, name string) {
	db, err := sql.Open("mysql", "DuckUser:Test123456@tcp(127.0.0.1:3306)/DuckDB") // Change the DB connection settings for your own machine.
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf("You scored %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	respercent := math.Round(percent)
	fmt.Printf("You Scored: %v%%.\n", respercent)
	_, err = db.Exec("INSERT INTO Result(`Name`, `Number of Correct Answers`, `Correct in percent`) VALUES(?,?,?)", name, score, respercent)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Welcome to my test!")
	fmt.Printf("Please type your name:")
	var myName string
	for { //Username response and check
		name := Userinput()
		myName = Uname(name)
		if myName != "" {
			break
		} else {
			fmt.Printf("Please don't use numbers, type in English and no spaces/empty spaces.")
		}
	}
	fmt.Printf("Welcome %v :), hope you enjoy this test!\n", myName)
	time.Sleep(2 * time.Second)
	fmt.Printf("Enter your age:")
	var (
		myAge bool
		age   string
	)
	for { // Age of user response and check.
		age = Userinput()
		lAge := LetterAge(age)
		if lAge != "" {
			break
		} else {
			fmt.Printf("Please type your age! (No spaces, letters)")
		}
	}
	myAge = Uage(age)
	if myAge == false {
		fmt.Printf("Can't do this test yet!\n")
		return
	} else {
		fmt.Println("Time for your test....")
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Write the options given as an answer")
	time.Sleep(2 * time.Second)
	unum_questions := Unum_questions()
	utest_score := Test()
	Result(utest_score, unum_questions, myName)
}
