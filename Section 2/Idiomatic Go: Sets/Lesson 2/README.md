Build a skill tracking system that manages a developer's technical abilities using the Go set idiom. This challenge demonstrates how to add new skills to a set while preventing duplicates and tracking the learning progress.

You will receive two inputs:

A string representing the number of skills to process (e.g., "5")
A string containing skills separated by commas (e.g., "Go,Docker,Kubernetes,Go,React,Docker,Python")
Your task is to:

Create a set using the Go idiom map[string]struct{} to store unique technical skills
Initialize the set with three starter skills: "Programming", "Problem Solving", and "Communication"
Parse the input string by splitting it on commas to get individual skill names
For each skill from the input, add it to your set using the empty struct literal {} as the value
Display the processing results by printing each skill as it's processed:
If the skill is new to the set: "Learning new skill: [skill]"
If the skill already exists in the set: "Already mastered: [skill]"
After processing all skills, display a progress summary:
"Skills processed: [total_processed]"
"New skills learned: [new_skills_count]"
"Total skills mastered: [total_unique_skills]"
Finally, list all skills in the developer's skill set:
Header: "Complete skill set:"
Each skill on a separate line: "✓ [skill]"
Use the strings package to split the input string and the strconv package to convert the count string to an integer. To check if a skill already exists before adding it, use the comma ok idiom: _, exists := skillSet[skill]. This challenge demonstrates how adding elements to a Go set automatically handles uniqueness while providing fast lookup performance for membership testing.