package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// ShowStatus will display the authenticated user's outbox, inbox, and history
var ShowStatus = &cobra.Command{
	Use: "status",
	RunE: func(cmd *cobra.Command, args []string) error {

		req, err := http.NewRequest("GET", "https://code-review-api1.herokuapp.com/my-requests", nil)
		if err != nil {
			return err
		}

		token, err := ReadCookie()
		req.AddCookie(&http.Cookie{Name: "code-review", Value: token})

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)

		b, _ := json.MarshalIndent(res, "", "\t")
		fmt.Println(string(b))

		return nil
	},
}

func formattedStatus(status map[string]map[string]interface{}) string {
	// outstrings := []string{
	// 	fmt.Sprintf("Account: %s", status["email"]), "\n",
	// 	"\tOutbox:\n",
	// }
	// for _, req := range status["outbox"] {

	// }
	return ""
}

func formattedRequest(req map[string]string) string {
	outstrings := ""

	return outstrings
}
