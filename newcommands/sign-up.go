package newcommands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

/* SAMPLE:
> code-review signup
> Signing you up for code-review:
> Enter your full name:
> Enter your email:
> Enter your password:
> Signinup...
> Your account has been created! use `code-review signin` to get started.
*/

// SignUp will hit the /signup endpoint in the code-review api
var SignUp = &cobra.Command{
	Use: "signup",
	RunE: func(cmd *cobra.Command, args []string) error {

		formBody := GetFormBody(map[string]string{
			"name":     "Enter your full name: ",
			"email":    "Enter your email address: ",
			"password": "Enter your password: ",
		})

		bytesArr, err := json.Marshal(formBody)
		if err != nil {
			return err
		}

		resp, err := http.Post("https://code-review-api1.herokuapp.com/signup", "application/json", bytes.NewBuffer(bytesArr))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)

		if res["error"] != nil {
			fmt.Print("Sign up failed: ")
			fmt.Print(res["error"])
		} else {
			fmt.Println(res["success"])
		}

		for _, cookie := range resp.Cookies() {
			if cookie.Name == "code-review" {
				err := StoreCookie(cookie.Value)
				if err != nil {
					return err
				}
			}
		}

		return nil
	},
}
