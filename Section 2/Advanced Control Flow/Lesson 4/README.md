In this challenge, you'll practice using the goto statement to create a simple retry mechanism. You're building a connection system that attempts to establish a connection and retries a specific number of times if it fails.

You will receive two inputs:

A string representing the maximum number of retry attempts (e.g., "3")
A string representing the attempt number that will succeed (e.g., "2")
Your task is to:

Parse the maximum retry attempts from the first input
Parse the success attempt number from the second input
Use a goto statement with a label to implement the retry logic
Start with attempt number 1 and increment after each failed attempt
For each attempt that fails, print "Attempt [attempt_number] failed"
When the attempt number matches the success attempt number, print "Attempt [attempt_number] succeeded" and exit
If you exceed the maximum retry attempts without success, print "All attempts failed" and exit
Use the goto statement to jump back to the retry label when an attempt fails and you haven't exceeded the maximum attempts. The attempt numbers should start from 1.