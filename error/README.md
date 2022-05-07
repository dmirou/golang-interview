# Errors

## Links

- https://go.dev/doc/effective_go#errors
- Dave Cheney. Don’t just check errors, handle them gracefully
https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
  - Sentinel errors (`ErrSomething`)
    - Never inspect the output of error.Error
    - Sentinel errors become part of your public API
    - Sentinel errors create a dependency between two packages
    - Conclusion: avoid sentinel errors
  - Error types (type `BadRequestError` struct)
    - The caller can use a type assertion or type switch, error types must be made public.
    - If your code implements an interface whose contract requires a specific error type, all implementors of that interface need to depend on the package that defines the error type.
    - This intimate knowledge of a package’s types creates a strong coupling with the caller, making for a brittle API.
    - Conclusion: avoid error types
  - Opaque errors (recommended)
    - Assert errors for behaviour, not type
    ```
    type temporary interface {
        Temporary() bool
    }
  
    // IsTemporary returns true if err is temporary.
    func IsTemporary(err error) bool
    ```
    - Annotating errors `errors.Wrap(err, "open failed")`
    ```
    // IsTemporary returns true if err is temporary.
    func IsTemporary(err error) bool {
      te, ok := errors.Cause(err).(temporary)
      return ok && te.Temporary()
    }
    ```
  - Only log error once in the root caller
- Uber Go style Guide. Errors. https://github.com/uber-go/guide/blob/master/style.md#errors

  | Error matching? | Error Message | Guidance |
  |---------------|----------|---------------|
  | No       | static        | errors.New |
  | No       | dynamic       | fmt.Errorf |
  | Yes      | static        | top-level var with errors.New |
  | Yes      | dynamic       | custom error type |
  
  - use `errors.Is` to check if target error is in the error chain 
  ```
  if err := foo.Open(); err != nil {
    if errors.Is(err, foo.ErrCouldNotOpen) {
      // handle the error
    } else {
      // handle unknown error
    }
  }
  ```
  - use `errors.As` to get target error from the error chain and use it
  ```
  if err := foo.Open("testfile.txt"); err != nil {
    var notFound *NotFoundError
    if errors.As(err, &notFound) {
      // handle the error, use exported field of NotFoundError
    } else {
      // handle unknown error
    }
  }
  ```
  
  - Error Wrapping (add context with `fmt.Errorf`)
    - Use `%w` if the caller should have access to the underlying error. 
      This is a good default for most wrapped errors, but be aware that callers may 
      begin to rely on this behavior. So for cases where the wrapped error is a known 
      var or type, document and test it as part of your function's contract.
    - Use `%v` to obfuscate the underlying error. Callers will be unable to match it, 
      but you can switch to %w in the future if needed.
  - Error Naming
    - For error values stored as global variables, use the prefix Err or err depending 
      on whether they're exported. `ErrBrokenLink = errors.New("link is broken")`
    - For custom error types, use the suffix `Error` instead.
      `type NotFoundError struct`