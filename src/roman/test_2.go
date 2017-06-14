package main

import "fmt"

func main() {
    //on an array, range returns the index
   /* a := [...]string{"a", "b", "c", "d"}
    for i := range a {
        fmt.Println("Array item", i, "is", a[i])
    }*/

    //on a map, range returns the key
    capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo" }
   /* for key := range capitals {
        fmt.Println("Map item: Capital of", key, "is", capitals[key])
    }
*/
    //range can also return two items, the index/key and the corresponding value
    for key2, val := range capitals {
        if key2 == "France"{
            fmt.Println("Map item: Capital of", key2, "is", val)
        }

    }
}