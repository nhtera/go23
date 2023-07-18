# Excercise 01

## Reordering Names Based on Country Code

### Goal: Create a program that reorders names based on the specified country's format.

- Inputs: First name, last name, middle name (optional), and country code.
- Outputs: Reordered name based on the country's format.
- Parse command line arguments to extract the first name, last name, middle name (if provided), and the country code.
- Determine the name format based on the country code. Reorder the name according to the determined format.
- Output the reordered name.

### Hint:
- Use the os package to access command line arguments os.Args.
- Handle cases where the middle name is not provided.
- Consider using a switch statement to determine the name format based on the country code.

### Example: 
`go run reorder_names.go John Smith VN`

Output: `Smith John`


`go run reorder_names.go Emily Rose Watson US`

Output: `Emily Watson Rose`