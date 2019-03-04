package commands

import "github.com/spf13/cobra"

/**
code-review open :folder :requestId (READ)
        1) Will search the db for your outbox/inbox (:folder) for the specified request
        2) Will display the code in your terminal

        If folder is INBOX {
            3) Will update the 'opened' property on the request object.
            4) Will display this:
            code-review open inbox 12

                ... code ...

                To Respond, Copy the above code, add comments to it and enter 'code-review respond 12 "${yourCodeHere}"'
        }
        If folder is OUTBOX {
            If 'cr_response' propery has value {
                3) Will update 'opened' propery on the request object.
                4) Will display:
                Status: Answered dd/mm
                Request:
                ${from}:
                    ${cr_request}

                Response:
                ${to}:
                    ${cr_response}

                5) Will move this request from OUTBOX to HISTORY in yours and responder objects.
            } else {
                3) Will Display:
                Request:
                    ${cr_request}
                Status: Unanswered
            }

        }
*/

// OpenRequest will find and display the request containing supplied id parameter
func OpenRequest() *cobra.Command {
	return &cobra.Command{
		Use: "open",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Implement OpenRequest() method body
			return nil
		},
	}
}
