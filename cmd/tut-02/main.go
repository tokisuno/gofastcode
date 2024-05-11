package main

import "fmt"

func main(){
    var intNum int
    fmt.Println(intNum)

    var floatNum float64 = 123.1
    fmt.Println(floatNum)

    var floatNum32 float32 = 10.1
    var intNum32 int32 = 2
    var result float32 = floatNum32 + float32(intNum32)
    fmt.Println(result)

    var string1 string = "My string"
    var string2 string = `Hello
This is me typing`

    fmt.Println(string1)
    fmt.Println(string2)

    const myConst string = "const value"
    fmt.Println(myConst)
}
