# Best practices

- use many, small interfaces
  - Single method interfaces are some of the most powerful and flexible
    - `io.Writer`, `io.Reader`, `interface{}`
- don't export intarfaces for types that will be consumed\
  publish the type, then others can define their own interfaces for it
- do export interface for types that wil be used by package
- deisgn functions and methods to receive interfaces whenever possible
