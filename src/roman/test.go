package main

import (
    "fmt"
    "sort"

)

func main() {
    m := map[string]string{"1": "I", "2": "II", "3": "III", "4": "IV"}
    mk := make([]string, len(m))
    i := 0
    for k, _ := range m {
        mk[i] = k
        i++
    }
    sort.Strings(mk)

    fmt.Println(mk)
}

