# Best practices

- Don't create goroutines in libraries\
  Let consumer control concurrency
- When creating a gorutines, know how it will end\
  Avoid subtle memory leaks
- Check for race conditions at compile time\
  `go run -race .`
