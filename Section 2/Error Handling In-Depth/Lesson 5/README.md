Build a web server request handler that uses recover to gracefully handle panics and prevent the entire server from crashing. This challenge demonstrates how to use recover inside a deferred function to catch panics and maintain server stability.

You will receive two inputs:

A string containing request details in the format "endpoint,method,action" (e.g., "/api/users,GET,process")
A string containing server configuration in the format "server_name,max_requests" (e.g., "WebServer,100")
Your task is to:

Parse the first input by splitting on commas to get endpoint, method, and action
Parse the second input by splitting on commas to get server name and max requests
Convert the max requests string to an integer
Create a function called handleRequest that takes endpoint, method, action, and server name as parameters
Inside handleRequest, add a deferred function that uses recover to catch any panics:
If a panic is recovered (recover returns non-nil), print: "Server [server_name] recovered from panic: [panic_value]"
After handling the panic, print: "Request handling completed safely"
After the deferred function, simulate request processing based on the action:
If action is "process": print "Processing [method] request to [endpoint]"
If action is "panic_nil": print "Attempting dangerous operation" then call panic("nil pointer access")
If action is "panic_overflow": print "Attempting resource allocation" then call panic("memory overflow")
If action is "panic_timeout": print "Attempting database connection" then call panic("connection timeout")
For any other action: print "Unknown action: [action]" then call panic("invalid action")
In the main function, print the server startup message: "Starting [server_name] with max [max_requests] requests"
Call the handleRequest function with the parsed parameters
After the function call, print the server status: "Server [server_name] is still running"
Finally, print a summary:
"Server Summary:"
"Name: [server_name]"
"Max Requests: [max_requests]"
"Last Request: [method] [endpoint]"
"Action Performed: [action]"
"Status: Operational"
Use the strings package to split the input strings on commas and the strconv package to convert the max requests string to an integer. This challenge demonstrates how recover allows servers to handle unexpected panics gracefully, preventing individual request failures from bringing down the entire system.