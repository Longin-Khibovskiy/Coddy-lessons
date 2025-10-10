In this challenge, you'll combine labeled break statements with switch and fallthrough to create a security monitoring system. The system searches through multiple security zones for threats and categorizes them with cascading alert levels.

You will receive two inputs:

A string representing the threat level to search for (e.g., "high")
A string representing the maximum zones to search (e.g., "3")
You are provided with the following security zones data:

zones := [][]string{
{"low", "medium", "low"},
{"medium", "high", "low"},
{"critical", "medium", "high"},
{"low", "critical", "medium"}
}
Your task is to:

Parse the threat level from the first input
Parse the maximum zones to search from the second input
Use nested loops with a labeled break to search through the zones
When you find the target threat level, print "Threat found at zone [zone_number] position [position]" and immediately exit both loops using labeled break
After finding the threat (or if no threat is found), use a switch statement with fallthrough to display security alerts based on the threat level you were searching for
The switch statement should work as follows:

For "critical": Print "CRITICAL: Lockdown initiated!" and fall through to print the high alert as well
For "high": Print "HIGH: Security breach detected!" and fall through to print the medium alert as well
For "medium": Print "MEDIUM: Increased monitoring active"
For "low": Print "LOW: Standard security protocols"
If the threat is not found after searching the specified number of zones, print "Threat not found in searched zones" before displaying the security alerts. Use 0-based indexing for zone numbers and positions within zones.