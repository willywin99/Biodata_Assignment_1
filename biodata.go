package main

import (
	"fmt"
	"os"
	"strings"
)

type Friends struct {
	number  int
	name    string
	address string
	job     string
	reason  string
}

var friendData = []Friends{
	{1, "Willy", "Apt. 794 Jl. MH. Thamrin No. 65, Batu, JK 06310", "Software Engineer", "Golang is easy to learn"},
	{2, "Roland", "Jl. Gatot Soebroto No. 21, Kotamobagu, PP 83797", "Quality Assurance", "Golang has great performance"},
	{3, "Yulia", "Apt. 478 Jl. Juanda No. 38, Sumenep, SG 15738", "Quality Assurance", "Golang is open source"},
	{4, "Juan", "Jl. MH. Thamrin No. 62, Jayawijaya, JA 74257", "Quality Assurance", "Golang has standard library"},
	{5, "Esa", "Apt. 661 Jl. Kartini No. 82, Pangkal Pinang, JA 65988", "Quality Assurance", "Golang has garbage collector"},
	{6, "Bill", "Apt. 519 Jl. Kartini No. 33, Belitung, YO 96209", "Quality Assurance", "Golang is supported by Google Developers"},
}

// Print data with student number params
func showData(studentNum int) {
	for _, friend := range friendData {
		if friend.number == studentNum {
			fmt.Println("Student Number:", studentNum)
			fmt.Println(strings.Repeat("-", 100))
			fmt.Println("Name\t\t\t\t\t:", friend.name)
			fmt.Println("Address\t\t\t\t\t:", friend.address)
			fmt.Println("Job\t\t\t\t\t:", friend.job)
			fmt.Println("What is the reason of picking Golang\t:", friend.reason)
			return
		}
	}
	fmt.Println("Student with number", studentNum, "is not found!!!")
}

func main() {
	// Check whether the argument is available
	if len(os.Args) < 2 {
		fmt.Println("Hint: go run main.go [student number]")
		return
	}

	// Get student number from comman line arguments
	studentNumber := os.Args[1]

	// Casting student number from string to integer
	var studNumInt int
	_, err := fmt.Sscanf(studentNumber, "%d", &studNumInt)
	if err != nil {
		fmt.Println("Student number must be an integer")
		return
	}

	showData(studNumInt)
}
