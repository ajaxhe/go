package main

import (
    "fmt"
)

func Unhex(c byte) byte {
    switch {
        case '0' <= c && c <= '9':
            return c - '0'
        case 'a' <= c && c <= 'f':
            return c - 'a' + 10
        case 'A' <= c && c <= 'F':
            return c - 'A' + 10
    }
    return 0
}

func ShouldEscape(c byte) bool {
    switch c {
        case ' ', '?', '=', '#', '+':
            return true
        case '&':
            return true
    }
    return false
}

func main() {
    fmt.Println(Unhex('F'))
    fmt.Println(Unhex('0'))
    fmt.Println(ShouldEscape('+'))
    fmt.Println(ShouldEscape('&'))
}
