# GO 2.0

As of the latest survey on what GO developers struggle with most, the proposal for Go 2.0 includes changes to the following aspects of the language:

1. Package management (modules)
2. Generics
3. Error Handling

## Package Management & Modules

## Generics

##  Error Wrapping

Check out the `errors.go` file in this directory which has a code snippet from our earlier discussion on errors. 

In this version of error handling, we get the message: 
`open notes.txt: no such file or directory`, but we are missing the type of erorr or any kind of stack trace indicating where this error came from. 

A common approach is to manually wrap the error in a larger structure providing additional information, like a timestamp. 

```go
type WrappedError struct {
    err error
    time time.Time
}

func wrap(err Error) *WrappedError {
    newErr := &WrappedError{err: err, time: time.Now()}
    return newErr
}

// other code 
if err != nil {
    return wrap(err)
}
```

Downsides to this include:
- Inconsistencies among various approaches to wrapping bewteen developers
- Having to hand rolling any useful information.
- Inconsistencies in the resulting type assertions

Go 2.0 is proposing the ability to unwrap errors without having to manually write your own functions around it by introducing the `Wrapper` interface.

```go
type Wrapper interface {
    // Returns the next error in the error chain, or nil
    Unwrap() error
}
```

All types that satisfy the `Wrapper` interface will have a layer of error inspection.

Take our WrappedError example above. 

```go
func (err *WrappedError) Unwrap() error {
    return err.err
}
```

// DISCUSS ERRORS AS AND IS

https://github.com/golang/go/issues/29934
https://golang.org/pkg/errors/
https://github.com/golang/go/wiki/ErrorValueFAQ
http://ieftimov.com/post/deep-dive-in....