package commands

import "github.com/spf13/cobra"

/**
code-review ask: (CREATE)
	1) Searches your directories files for "!CR @username ... !CR end"

	2) Creates code review 'request' and adds it to the @usernames's inbox, and your outbox
*/

// CreateCRRequest will put a request object into the recipients inbox
func CreateCRRequest() *cobra.Command {
	return &cobra.Command{
		Use: "ask",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Implement CeateCRRequesr() method body
			return nil
		},
	}
}
