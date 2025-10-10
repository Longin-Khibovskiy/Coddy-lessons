Build a music playlist analyzer that processes song collections and displays detailed information about each playlist using the Go set idiom. This challenge demonstrates how to iterate over sets to examine and report on their contents.

You will receive two inputs:

A string containing rock songs separated by commas (e.g., "Bohemian Rhapsody,Stairway to Heaven,Hotel California,Sweet Child O Mine")
A string containing pop songs separated by commas (e.g., "Shape of You,Blinding Lights,Watermelon Sugar,Levitating")
Your task is to:

Create two sets using the Go idiom map[string]struct{} to store rock songs and pop songs separately
Parse both input strings by splitting them on commas to get individual song titles
Add each song to the appropriate set using the empty struct literal {} as the value
Display the rock playlist analysis:
Header: "Rock Playlist Analysis:"
Count: "Total rock songs: [count]"
Songs header: "Songs in rock playlist:"
Each song on a separate line: "♪ [song_title]"
Display the pop playlist analysis:
Header: "Pop Playlist Analysis:"
Count: "Total pop songs: [count]"
Songs header: "Songs in pop playlist:"
Each song on a separate line: "♪ [song_title]"
Create a combined playlist by merging both sets into a new set
Display the combined playlist analysis:
Header: "Combined Playlist Analysis:"
Count: "Total unique songs: [count]"
Songs header: "All songs in combined playlist:"
Each song on a separate line: "♫ [song_title]"
Finally, display a summary:
"Playlist Summary:"
"Rock songs: [rock_count]"
"Pop songs: [pop_count]"
"Combined unique songs: [combined_count]"
Use the strings package to split the input strings on commas. To iterate over each set, use the for...range loop with a single variable to access only the keys (song titles). This challenge demonstrates how iterating over Go sets allows you to examine all elements in a collection, making it perfect for analysis and reporting tasks.