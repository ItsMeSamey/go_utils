//! Call Load(<filename>, os.Setenv), eg `Load(".env", os.Setenv)`
//!
//! You can use custom setter for logging, for example to print env vars that are being set, use
//!   Load(<filename>, func (k, v string) error { fmt.Println("Set: ", k, ":=", v); return os.Getenv(k, v) })

package utils

import (
  "os"
  "bytes"
  "errors"
  "unsafe"
  "strconv"
)

// Loads the provided env file.
func Load(file string, setter func (k, v string) error) error {
  // BytesToString converts a slice of bytes to string without memory allocation.
  BytesToString := func (b []byte) string {
    return unsafe.String(unsafe.SliceData(b), len(b))
  }

  data, err := os.ReadFile(file)
  if err != nil {
    return errors.New("error reading env file `" + file + "`\n" + err.Error())
  }

  for _, line := range bytes.Split(data, []byte{'\n'}) {
    i := bytes.IndexByte(line, '=')
    if i == -1 || i+1 >= len(line) { continue }

    varName := BytesToString(line[:i])
    varVal := BytesToString(line[i+1:])

    if varVal[0] == '"' || varVal[0] == '\'' {
      if varVal, err = strconv.Unquote(varVal); err != nil {
        return errors.New("Unescaping error for `"+ varName +"` and val `" + varVal + "`\n"+ err.Error())
      }
    }

    if err := setter(varName, varVal); err != nil {
      return errors.New("Error setting `"+ varName +"` to `" + varVal + "`\n"+ err.Error())
    }
  }

  return nil
}

