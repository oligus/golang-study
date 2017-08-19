# Go Langugage Cheat Sheet

## Strings

##### Concatinate


```go
str := "my cool string"
str2 := "not " + str
```
  
```go
str := "my cool string"
str2 := strings.Join([]string{"not", str}, " ")
```

Result:
`str2` = `"not my cool string"` 