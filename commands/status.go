package commands

import "github.com/spf13/cobra"

/**
code-review status: (INDEX)
        1) Goes through your outbox and checks to see if any requests's dateResponded properties have a value
        2) Displays the one's that have been answered and the ones that are still in outbox

                code-review status

                OUTBOX:
                 ID# 4) To: @Ikey Subject: "Efficiency Question" Status: Unanswered

                INBOX:
                 ID# 12) From: @Ikey Subject: "Efficiency Question"
*/

// GetRequestStatus will index and display all requests in your outbox and inbox
func GetRequestStatus() *cobra.Command {
	return &cobra.Command{
		Use: "ask",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Implement RespondToRequest() method body
			return nil
		},
	}
}
