Build a team member verification system that checks if specific people are part of a project team using the Go set idiom. This challenge demonstrates how to use the comma ok idiom to test for membership in a set.

You will receive two inputs:

A string containing team members separated by commas (e.g., "Alice,Bob,Charlie,Diana,Eve")
A string containing people to check separated by commas (e.g., "Bob,Frank,Alice,Grace,Charlie")
Your task is to:

Create a set using the Go idiom map[string]struct{} to store the team members
Parse the first input string by splitting it on commas to get individual team member names
Add each team member to your set using the empty struct literal {} as the value
Parse the second input string by splitting it on commas to get the list of people to verify
For each person to check, use the comma ok idiom _, ok := teamSet[person] to test membership
Display the verification results for each person being checked:
If the person is in the team: "[person] is on the team"
If the person is not in the team: "[person] is not on the team"
After checking all people, display a summary:
"Team size: [team_member_count]"
"People checked: [people_checked_count]"
"Team members found: [found_count]"
"Non-team members: [not_found_count]"
Use the strings package to split the input strings on commas. This challenge demonstrates how the comma ok idiom provides a safe and efficient way to check for membership in a Go set, returning both the value and a boolean indicating whether the key exists.