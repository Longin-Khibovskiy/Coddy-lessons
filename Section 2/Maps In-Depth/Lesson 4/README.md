Build a text analysis tool that counts word frequencies in a document. This challenge demonstrates how maps excel at aggregating data by tracking how many times each unique word appears in a piece of text.

You will receive two inputs:

A string representing the text to analyze (e.g., "the quick brown fox jumps over the lazy dog the fox is quick")
A string representing the minimum frequency threshold (e.g., "2")
Your task is to:

Create a function called countWords that takes a string of text and returns a map[string]int where:
Keys are individual words (strings)
Values are the frequency count of each word (integers)
Parse the input text by splitting it into individual words using spaces as delimiters
Count the frequency of each word in the text:
Convert all words to lowercase for consistent counting
For each word, increment its count in the map
If a word doesn't exist in the map yet, it will start at 0 and be incremented to 1
Filter the results to show only words that appear at least as many times as the threshold
Display the results in the following format:
Header: "Word Frequency Analysis:"
For each qualifying word: "[word]: [count]"
Display words in alphabetical order
Calculate and display summary statistics:
"Total unique words: [total_unique_count]"
"Words above threshold: [filtered_count]"
"Most frequent word: [word] ([count] times)"
Use the strings package to split the text and convert to lowercase, the strconv package to convert the threshold string to an integer, and the sort package to sort the words alphabetically. This challenge demonstrates how maps provide an elegant solution for counting and aggregating data, with the zero-value behavior of maps making the counting logic simple and efficient.