/////////////////////////////////////////
// util.go - Password termui utilities
// Mike Schilli, 2022 (m@perlmeister.com)
/////////////////////////////////////////
package main

func mask(s string) string {
  masked := []byte(s)

  tomask := false

  for i := 0; i < len(s); i++ {
    if tomask {
      masked[i] = '*'
    } else {
      masked[i] = s[i]
    }
    if s[i] == ' ' {
      tomask = true
    }
  }
  return string(masked)
}

