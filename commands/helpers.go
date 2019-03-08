package commands

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// GetInputs takes a slice of prompts, and returns user input for them
func GetInputs(prompts []string) []string {

	answers := make([]string, len(prompts))

	for index, prompt := range prompts {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(prompt + " ")
		answer, _ := reader.ReadString('\n')

		answers[index] = answer
	}

	return answers
}

// GetFormBody takes an map of property:question and using stdin
// returns a map of property:answer
func GetFormBody(body map[string]string) map[string]string {

	answers := make(map[string]string, len(body))

	for prop, prompt := range body {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(prompt + " ")
		answer, _ := reader.ReadString('\n')

		answers[prop] = strings.TrimSpace(answer)
	}

	return answers
}

// StoreCookie writes text to a given filename
func StoreCookie(text string) error {
	// Write value to file called cookie.txt
	f, err := os.Create("code-review-cookie.txt")
	if err != nil {
		return err
	}

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}

// ReadCookie reads text from a given filename
func ReadCookie() (string, error) {
	dat, err := ioutil.ReadFile("code-review-cookie.txt")
	if err != nil {
		return "", err
	}

	return string(dat), nil
}
