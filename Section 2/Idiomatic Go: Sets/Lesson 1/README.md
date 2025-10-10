Create a programming language preference tracker that demonstrates the Go set idiom for storing unique items. This challenge shows how to use the map[string]struct{} pattern to track programming languages without duplicates.

You will receive two inputs:

A string representing the number of languages to process (e.g., "6")
A string containing programming languages separated by commas (e.g., "Go,Python,JavaScript,Go,Java,Python,C++,JavaScript")
Your task is to:

Create a set using the Go idiom map[string]struct{} to store unique programming languages
Parse the input string by splitting it on commas to get individual language names
Add each language to your set using the empty struct literal {} as the value
Display the processing results by printing each language as it's encountered:
If the language is new to the set: "Added: [language]"
If the language already exists in the set: "Already exists: [language]"
After processing all languages, display a summary:
"Total languages processed: [total_count]"
"Unique languages: [unique_count]"
Finally, list all unique languages in the set:
Header: "Programming languages in set:"
Each language on a separate line: "- [language]"
Use the strings package to split the input string and the strconv package to convert the count string to an integer. To check if a language already exists in the set before adding it, use the comma ok idiom: _, exists := languageSet[language]. This challenge demonstrates how the Go set idiom provides an efficient way to track unique items and prevent duplicates in your data.