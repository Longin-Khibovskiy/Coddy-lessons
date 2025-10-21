Fix the visibility issues in a Go package by correcting the naming convention for exported and unexported identifiers. You'll work with a simulated multi-package scenario where some identifiers are incorrectly named, causing compilation errors when accessed from other packages.

You will receive two inputs:

A string containing package definitions in the format "package_name:identifier1:type1:visibility1,identifier2:type2:visibility2|package_name:identifier1:type1:visibility1" (e.g., "utils:helper:function:private,Calculate:function:public|main:name:variable:private,Version:variable:public")
A string containing access attempts in the format "accessing_package.identifier1,accessing_package.identifier2,accessing_package.identifier3" (e.g., "main.helper,main.Calculate,utils.Version")
Your task is to:

Parse the first input by splitting on pipes to get individual package definitions
For each package definition, split on colons to get package name and its identifiers
For each identifier definition, split on colons to get identifier name, type, and intended visibility
Create a function called fixIdentifierName that takes an identifier name and intended visibility, and returns the corrected name:
If visibility is "public", ensure the identifier starts with a capital letter
If visibility is "private", ensure the identifier starts with a lowercase letter
Display the package analysis header: "=== PACKAGE VISIBILITY ANALYSIS ==="
For each package, display the package information:
"Package: [package_name]"
For each identifier in the package:
If the identifier name needs fixing: "- [original_name] ([type]) -> [corrected_name] (visibility: [visibility])"
If the identifier name is correct: "- [identifier_name] ([type]) (visibility: [visibility]) ✓"
Parse the second input by splitting on commas to get individual access attempts
For each access attempt, split on dots to get the accessing package and the identifier being accessed
Display the access validation header: "=== ACCESS VALIDATION ==="
For each access attempt, determine if the access is valid:
Find the package that contains the identifier
Check if the corrected identifier name starts with a capital letter (exported)
If the identifier is exported: "[accessing_package] accessing [target_package].[corrected_identifier_name]: ✓ ALLOWED (exported)"
If the identifier is unexported: "[accessing_package] accessing [target_package].[corrected_identifier_name]: ✗ DENIED (unexported)"
If the identifier doesn't exist: "[accessing_package] accessing [target_package].[identifier_name]: ✗ NOT FOUND"
Count and display the summary statistics:
"=== SUMMARY ==="
"Total packages analyzed: [number_of_packages]"
"Total identifiers processed: [total_number_of_identifiers]"
"Identifiers requiring fixes: [number_of_identifiers_that_needed_fixing]"
"Access attempts: [number_of_access_attempts]"
"Allowed accesses: [number_of_allowed_accesses]"
"Denied accesses: [number_of_denied_accesses]"
Display the final recommendations:
If there were identifiers requiring fixes: "Recommendation: Fix identifier naming to follow Go conventions"
If all identifiers were correct: "All identifiers follow proper Go naming conventions"
Use the strings package to split the input strings, the fmt package for formatted output, and string manipulation functions to check and modify identifier names. This challenge demonstrates how Go's naming convention controls package visibility and helps you understand the difference between exported and unexported identifiers.