//! To use this, wrap every error return with `WithStack`
//! To enable stack traces set `SHOW_ERROR_STACK_TRACES` environment variable to `true`

package utils

import (
  "os"
  "unsafe"
  "runtime"
)

// when this is true, errors will have stack trace
// changing this to `const show = false` should eliminate any performance penalty
var show = os.Getenv("SHOW_ERROR_STACK_TRACES") == "true"

// Max size of the stak trace
const maxTraceLen = 1024 * 16


// init funcction that is called automatically
func init() { if !show { runtime.StartTrace() } }

type ErrorWithStack struct {
  current     string
  originalLen int
  original    error
}

func (e ErrorWithStack) Error() string { return e.current }
func (e ErrorWithStack) Unwrap() error { return e.original }
func (e ErrorWithStack) WithoutStack() string { return e.current[:e.originalLen] }

// Done this way to reduce cost when `show` is false

// Adds stack trace to errors when in dev mode i.e. `ENV="dev"`
var WithStack = func () func (err error) error {
  if !show { return func (err error) error { return err } }

  return func (err error) error {
    if err == nil { return nil }
    if _, ok := err.(ErrorWithStack); ok { return err }

    out := make([]byte, maxTraceLen, maxTraceLen)
    originalString := err.Error()
    return ErrorWithStack{
      current: originalString + "\n##-STACK-##\n" + unsafe.String(unsafe.SliceData(out), runtime.Stack(out, false)),
      originalLen: len(originalString),
      original: err,
    }
  }
}()

