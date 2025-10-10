In this challenge, you'll create a communication system that demonstrates polymorphism through interfaces. You'll build a Speaker interface and implement it with different types that can "speak" in their own unique ways.

You will receive three inputs:

A string representing the number of speakers to create (e.g., "3")
A string containing speaker types separated by commas (e.g., "person,parrot,person")
A string containing speaker names separated by commas (e.g., "Alice,Polly,Bob")
Your task is to:

Define a Speaker interface with one method signature: Speak() that returns a string
Define a Person struct with a Name field (string)
Define a Parrot struct with a Name field (string)
Implement the Speak() method for both structs:
Person should return: "Hello, I'm [name]"
Parrot should return: "Squawk! [name] says hello!"
Create a function called makeAllSpeak that takes a slice of Speaker interfaces as a parameter
Inside makeAllSpeak, iterate through the slice and call the Speak() method on each speaker, printing the result
Parse the inputs to create the appropriate speaker instances based on the types and names provided
Store all speakers in a slice of Speaker interfaces
Call makeAllSpeak with your slice of speakers
The speakers should be created in the order they appear in the input. For example, if the types are "person,parrot,person" and names are "Alice,Polly,Bob", create a Person named Alice, then a Parrot named Polly, then a Person named Bob. This challenge demonstrates how different types can be treated uniformly through a shared interface, showcasing the power of polymorphism in Go.