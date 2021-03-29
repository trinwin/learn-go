# Errors Handling

- Error is interface - just another type like int, float...
- It's built-in to the language
- Any type that has the Error() string attached to it is an erro type
- We can create custom error type
- Always check for errors
- Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
- We can add information to errors to create a custom one
- In Go it’s idiomatic to communicate errors via an explicit, separate return value
- Blog about [Error Handling](https://blog.golang.org/error-handling-and-go)
