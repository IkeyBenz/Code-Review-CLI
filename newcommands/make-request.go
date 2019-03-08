package newcommands

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// MakeRequest hits the /requests post route in the core-review api
var MakeRequest = &cobra.Command{
	Use: "request",
	RunE: func(cmd *cobra.Command, args []string) error {

		// formBody := GetFormBody(map[string]string{
		// 	"answerer": "Enter recipient's email address: ",
		// 	// You're going to have to get a different method to grab this as there will be \n characters in it
		// 	"cr_request": "Enter your code in quotes: ",
		// })

		// fmt.Println("Sending your request...")

		// bytesArr, err := json.Marshal(formBody)
		// if err != nil {
		// 	return err
		// }

		// req, err := http.NewRequest("POST", "https://code-review-api1.herokuapp.com/requests", bytes.NewBuffer(bytesArr))
		// if err != nil {
		// 	return err
		// }

		// token, err := ReadCookie()
		// req.AddCookie(&http.Cookie{Name: "code-review", Value: token})
		// req.Header.Set("Content-Type", "application/json")

		// client := &http.Client{}
		// resp, err := client.Do(req)
		// if err != nil {
		// 	return err
		// }
		// defer resp.Body.Close()

		// var res map[string]interface{}

		// json.NewDecoder(resp.Body).Decode(&res)

		// fmt.Println(res)
		// fmt.Println(res["data"])

		// Ask for relative file path
		filePath := GetInputs([]string{"Enter the relative path of the file containing your code: "})[0]
		_, err := parseFileForCodeBlock(filePath)
		if err != nil {
			return err
		}
		return nil

	},
}

func parseFileForCodeBlock(relativePath string) (string, error) {

	// Get file path
	dir, _ := os.Getwd()
	filePath := path.Join(dir, relativePath)
	fmt.Println(filePath)

	// Read file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}
