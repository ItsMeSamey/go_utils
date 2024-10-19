//! To use this, wrap every error return with `WithStack`
//! To enable stack traces set `SHOW_ERROR_STACK_TRACES` environment variable to `true`

package utils

import (
  "os"
  "errors"
  "runtime"
)

// when this is true, errors will have stack trace
// changing this to `const show = false` should eliminate any performance penalty
var show = os.Getenv("SHOW_ERROR_STACK_TRACES") == "true"

// Max size of the stak trace
const maxTraceLen = 1024 * 16


// init funcction that is called automatically
func init() { if !show { runtime.StartTrace() } }

type ErrorWithStack struct { error }

// Done this way to reduce cost when `show` is false

// Adds stack trace to errors when in dev mode i.e. `ENV="dev"`
var WithStack = func () func (err error) error {
  if !show { return func (err error) error { return err } }

  return func (err error) error {
    if err == nil { return nil }
    if _, ok := err.(ErrorWithStack); ok { return err }

    out := make([]byte, maxTraceLen, maxTraceLen)
    // runtime.Stack is called with all = false to prevent world stopping!
    err = ErrorWithStack{ errors.New(err.Error() + "\n-- STACK --\n" + string(out[:runtime.Stack(out, false)])) }
    return err
  }
}()

