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
func CheckName(Name string) bool {
	for _, r := range Name {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// Introductory phase that asks for your name. Loops if you type it wrong. Change the DB connection settings for your own machine.
func main() {
	db, err := sql.Open("mysql", "DuckUser:Test123456@tcp(127.0.0.1:3306)/DuckDB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Welcome to my test!")
	fmt.Printf("Please type your name:")
	reader := bufio.NewReader(os.Stdin)
	var Name string
	for {
		Name, _ = reader.ReadString('\n')
		Name = strings.TrimSpace(Name)
		checked := CheckName(Name)
		if checked == false {
			fmt.Printf("Please don't use numbers, type in English and no spaces/empty spaces.")
		} else if len(Name) == 0 {
			fmt.Printf("Please don't use numbers, type in English and no spaces/empty spaces.")
		} else {
			break
		}
	}
	fmt.Printf("Welcome %v :), hope you enjoy this test!\n", Name)
	time.Sleep(2 * time.Second)
	fmt.Printf("Enter your age:")
	var (
		age           string
		ageOld        int
		num_questions int
	)
	//Score check and a variable with the amount of questions. Throughout the questions, it will keep a count of correct answers. Tracks how many questions there is as well.
	score := 0
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
	// Age check. If age check fails, exit program. If age has white space or letters, loop.
	for {
		age, _ = reader.ReadString('\n')
		age = strings.TrimSpace(age)
		IsLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(age)
		if len(age) == 0 {
			fmt.Printf("Please type your age! (No spaces, letters)")
		} else if IsLetter == true {
			fmt.Printf("Please type your age! (No spaces, letters)")
		} else {
			break
		}
	}
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
	// In this section, we use the DB to select the questions the user has set in the DB file, with their own correct answers. DB and program have a regexp check that allows whitespaces and case insensitive answers
	var (
		question  string
		answer    string
		coranswer string
		matched   bool
	)
	rows, err = db.Query("select Question, `Correct Answer` from Questions")
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
	// At last, the program counts how many questions were correct, and the percentages the user got. Adds it to a separate result table for tracking
	fmt.Printf("You scored %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	respercent := math.Round(percent)
	_, err = db.Exec("INSERT INTO Result(`Name`, `Number of Correct Answers`, `Correct in percent`) VALUES(?,?,?)", Name, score, respercent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You Scored: %v%%.\n", respercent)
}
