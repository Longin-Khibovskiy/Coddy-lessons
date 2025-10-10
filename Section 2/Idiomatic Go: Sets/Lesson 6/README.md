Build a username registration system that prevents duplicate usernames using the Go set idiom. This challenge demonstrates how to create a robust user management system that ensures username uniqueness across all registrations.

You will receive two inputs:

A string containing existing usernames separated by commas (e.g., "alice123,bob_dev,charlie2024,diana_codes")
A string containing new username registration attempts separated by commas (e.g., "eve_user,alice123,frank_dev,bob_dev,grace2024,alice123")
Your task is to:

Create a set using the Go idiom map[string]struct{} to store registered usernames
Parse the first input string by splitting it on commas to get existing usernames
Add each existing username to your set using the empty struct literal {} as the value
Parse the second input string by splitting it on commas to get new registration attempts
For each registration attempt, check if the username already exists in the set using the comma ok idiom
Display the registration process for each attempt:
If the username is available: "Registered: [username]" and add it to the set
If the username already exists: "Username taken: [username]"
After processing all registration attempts, display a comprehensive summary:
"Registration Summary:"
"Initial usernames: [initial_count]"
"Registration attempts: [attempts_count]"
"Successful registrations: [successful_count]"
"Rejected (duplicate): [rejected_count]"
"Total registered usernames: [total_count]"
Finally, list all registered usernames in the system:
Header: "All registered usernames:"
Each username on a separate line: "@[username]"
Use the strings package to split the input strings on commas. This challenge demonstrates how the Go set idiom provides an efficient solution for preventing duplicates in user registration systems, ensuring data integrity while providing fast lookup performance for username validation.