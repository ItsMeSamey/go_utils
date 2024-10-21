//! WARNING: String in golang are immutable
//!   Therefor, bytes returned from S2B MUST not be changed
//! and
//!   bytes given to B2S MUST not be changed as long as the string is in use

package utils

import "unsafe"

// WARNING: DONOT MUTATE RETURNED VALUE
func S2B(s string) []byte {
  return unsafe.Slice(unsafe.StringData(s), len(s))
}

// WARNING: DONOT MUTATE THE ENTERED VALUE AS LONG AS THE RETURNED STRING IS IN USE
func B2S(b []byte) string {
  return unsafe.String(unsafe.SliceData(b), len(b))
}

