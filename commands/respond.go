package commands

import "github.com/spf13/cobra"

/**
code-review respond :requestId "a shit ton of code with comments" (RESPOND)
        1) Will update the 'cr_response' of the request with the code parameter,
        2) Will change the 'opened' property of the response to false
        3) Will change the 'date_responded' of the request to be the current date

*/

// RespondToRequest will set cr_response of requestId to whatever the response string includes
func RespondToRequest() *cobra.Command {
	return &cobra.Command{
		Use: "ask",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Implement RespondToRequest() method body
			return nil
		},
	}
}
