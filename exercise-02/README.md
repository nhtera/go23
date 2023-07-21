# Excecise 2

## Create a package and a command-line tool to sort input provided by the user.

### Goal:

- Inputs: Inputs: Number (integer or float) array, string array.
- Outputs: Sorted result based on the provided input type.
- Create a Go package with functions for sorting integer arrays, float arrays, and string arrays.
- Implement sorting logic for each data type using appropriate algorithms.
- Create a command-line tool (CLI) to parse the input from the command line.
- Determine the type of input (integer array, float array, string array, or mixed).
- Utilize the corresponding sorting function from the package to sort the elements.
- Output the sorted result.

### Hint:

- Use util package to access command line arguments: https://github.com/spf13/cobra, https://github.com/devfacet/gocmd
- Use the flag package to parse command line arguments.
- Create separate functions in the package for sorting each data type.
- Consider implementing generic sorting functions using interfaces to handle mixed input types.


### Example:

Example 1:
`go run sorted.go -int 5 2 10 1`

Output: `1 2 5 10`

Example 2:
`go run sorted.go -string apple orange banana`

Output: `apple banana orange`

Example 3:
`go run sorted.go -mix 5.5 apple 2.7 orange 3 banana`

Output: `2.7 3 5.5 apple banana orange`


## Solution
- Using Cobra to parse command line arguments (https://cobra.dev/)
- Package utils: contains sort function for each data type

### How to run:

- Parameters:

```
--type: input type (int, float, string, mix)
--input: input array values (separate by commas)
```


- Demo:

```
go run main.go --type int --input "5,2,10,1"
go run main.go --type string --input "apple,orange,banana"
go run main.go --type mix --input "5.5,apple,2.7,orange,3,banana"
```