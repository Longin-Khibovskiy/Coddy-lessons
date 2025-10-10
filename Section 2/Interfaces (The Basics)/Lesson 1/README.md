In this challenge, you'll practice defining interfaces in Go by creating a simple media player system. You'll define an interface that establishes a contract for different types of media devices.

You will receive one input:

A string representing a media action (e.g., "play", "pause", or "stop")
Your task is to:

Define an interface called MediaPlayer with three method signatures:
Play() that returns a string
Pause() that returns a string
Stop() that returns a string
Based on the input action, print the corresponding method signature from your interface definition in the format: "MediaPlayer interface requires: [MethodName]() string"
The method names should be capitalized (exported) as they are part of an interface definition. For the input "play", you should print information about the Play() method, for "pause" print about the Pause() method, and for "stop" print about the Stop() method.

This challenge focuses on the fundamental concept of defining interfaces with method signatures, establishing the contract that any implementing type must fulfill.