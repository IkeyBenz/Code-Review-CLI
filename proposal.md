# Code Review CLI

## *Problem*
There is currently no efficient way to have code reviewed and commented on without either meeting up in person, or sending chunks of code back and fourth over text message.

## *Proposed Solution*: 
1) UserA marks a line in their project with `// !CR @UserB` and underneath explains what they're looking to have reviewed, then marks the end of their code chunk with `// !CR end`. 
2) UserA uses the cli to upload that request to UserB
3) The cli inserts a unique ID for that code review request on the same line as the comment: `// !CR @UserB -fdjndsdjfkjn` and uploads the request to UserB.
3) UserB logs in to cli and sees a CR request from UserA
4) UserB writes feedback on the code chunk and sends it back to UserA
5) UserA checks the cli for any updates
6) The cli notices UserB's response, parses the current directory for a line containing the ID of the original request, and inserts UserB's comments right next to the reviewed chunk of code.

