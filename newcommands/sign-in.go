package newcommands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

/* SAMPLE:
> code-review signin
> Signing you in to code-review:
> Enter your email:
> Enter your password:
> Signing you in...
> You are now signed into code-review
*/

// SignIn will hit the /signup endpoint in the code-review api
var SignIn = &cobra.Command{
	Use: "signin",
	RunE: func(cmd *cobra.Command, args []string) error {

		formBody := GetFormBody(map[string]string{
			"email":    "Enter your email address: ",
			"password": "Enter your password: ",
		})

		bytesArr, err := json.Marshal(formBody)
		if err != nil {
			return err
		}

		resp, err := http.Post("https://code-review-api1.herokuapp.com/signin", "application/json", bytes.NewBuffer(bytesArr))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)

		fmt.Println(res)
		fmt.Println(res["data"])

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
