package main

import (
    "errors"
    "fmt"
)

func main(){
    answer := "Hello there"
    printMe(answer)

    var x int = 11
    var y int = 2
    var result, remainder, err = intDivision(x, y)

    switch{
    case err!=nil:
        fmt.Printf(err.Error())
    case remainder==0:
        fmt.Printf("The result of the int division is %v", result)
    default:
        fmt.Printf("The result of the int division is %v with remainder %v", result, remainder)
    }

}

func printMe(printValue string){
    fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
    var err error
    if denominator == 0 {
        err = errors.New("Cannot divide by 0")
        return 0, 0, err
    }
    var result int = numerator/denominator
    var remainder int = numerator%denominator
    return result, remainder, err

}
