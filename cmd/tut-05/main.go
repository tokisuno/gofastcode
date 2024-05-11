package main

import (
    "fmt"
    "strings"
)

func main(){
    var myString = "résumé"
    var betterString = []rune("résumé")
    var indexed = myString[0]

    fmt.Println(betterString)
    fmt.Println(indexed)
    fmt.Printf("%v, %T", indexed, indexed)

    for i, v := range myString{
        fmt.Println(i, v)
    }

    var strSlice = []string{"p", "u", "t", "a"}
    var strBuilder strings.Builder
    for i := range strSlice{
        strBuilder.WriteString(strSlice[i])
    }
    var catStr = strBuilder.String()
    fmt.Printf("\n%v", catStr)


}
