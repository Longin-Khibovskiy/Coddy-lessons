In this challenge, you'll practice using the fallthrough keyword in a switch statement to create a priority alert system. When a high-priority alert is triggered, the system should display multiple warning messages in sequence.

You will receive a string input representing an alert level. Create a switch statement that handles different alert levels and uses fallthrough to display cascading messages for higher priority alerts.

The input will be one of these alert levels: "critical", "high", "medium", or "low".

Your switch statement should work as follows:

For "critical": Print "CRITICAL: System shutdown imminent!" and fall through to print the high priority message as well
For "high": Print "HIGH: Immediate attention required!" and fall through to print the medium priority message as well
For "medium": Print "MEDIUM: Please review when possible"
For "low": Print "LOW: Informational only"
Use the fallthrough keyword so that critical alerts display both critical and high priority messages, and high alerts display both high and medium priority messages.