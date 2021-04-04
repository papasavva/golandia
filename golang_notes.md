# Golang Notes

## Variables
### Declaration
- Short declarations in toplevel is not allowed
- Redeclaration in toplevel is not allowed
- := operator is not pure declare, but more like declare and assign
### Data types
- Go source files are always encoded in UTF-8 and strings are conventionally interpreted as UTF-8, so we can include Unicode code points in string literals.
- A slice has 3 components: a pointer, a length and a capacity.

### Productivity
#### Best IDE
- Use GoLand or IntelliJ with [Go plugin](https://plugins.jetbrains.com/plugin/9568-go).

#### File watchers
- Install [File Watchers plugin](https://plugins.jetbrains.com/plugin/7177-file-watchers) on IntelliJ/GoLand
- Add add go fmt and goimports. This will format the file and introduce imports to the code on save. 
- Then open directory as project from project browser.
- On file save, the code will be formatted and introduce the necessary imports.

## Error Handling
- Go communicates errors via an explicit, separate return value.
- By convention, error is the last return value and have type error.
- A nil value in the error, indicates that there was no error.