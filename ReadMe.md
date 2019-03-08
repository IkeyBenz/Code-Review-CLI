[![Go Report Card](https://goreportcard.com/badge/github.com/IkeyBenz/Code-Review-CLI)](https://goreportcard.com/report/github.com/IkeyBenz/Code-Review-CLI)

# Code Review CLI 

This command line interface empowers its users by allowing them to efficiently send a request for code review to a peer, from the comfort of their IDE. 

# Usage:
#### There are four main commands:
1) `code-review signup`: Prompts you to enter name, email, and password. Creates your account for code-review and signs you in.
2) `code-review request`: Prompts you to enter email of recipient and block of code that you'd like to have reviewed. Sends your request to the intendede reviewer.
3) `code-review status`: Retrieves and displays your current outbox, inbox, and history of code-review requests.
4) `code-review respond`: Prompts you to enter requestId and revised code snippet. Sends the revision to the person who requested the review.

# Installation
1) Clone this repository
2) In terminal in folder you cloned, run `go run && go install`
3) Run `code-review` to use!