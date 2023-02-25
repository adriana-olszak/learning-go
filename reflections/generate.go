package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	// Create the file name with the current date in the format of "YYYY-MM-DD"
	// Important, after the doc https://pkg.go.dev/time#Time.Format
	// The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
	// There are some wrinkles illustrated below.
	dateStr := time.Now().Format("2006-01-02")
	fileName := dateStr + "-reflection.md"

	templatePath := "./template.md"
	templateData, err := ioutil.ReadFile(templatePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", templatePath, err)
		os.Exit(1)
	}

	// Replace the $DATE$ placeholder with the current date
	fileData := strings.ReplaceAll(string(templateData), "$DATE$", dateStr)

	// Create the new file with the updated content
	err = ioutil.WriteFile(fileName, []byte(fileData), 0644)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", fileName, err)
		os.Exit(1)
	}

	fmt.Printf("Created new file %s with today's date %s.\n", fileName, dateStr)
}
