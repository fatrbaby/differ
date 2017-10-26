 
### differ

find the same file in your disk!

```go
diff := differ.New(directory)
diff.Scan()
fmt.Println(diff.Sames())
```