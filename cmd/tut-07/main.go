package main

import "fmt"

func main(){
    var p *int32 = new(int32) // stores memaddr of int32 (points to nothing)
    var i int32               // the int32 that p will point to

    fmt.Printf("The value p points to is: %v", *p) // reference value to pointer
    fmt.Printf("\nThe value of i is: %v", i)

    p = &i  // p = i : p references to memaddr of i
            // p and i reference same int32 value in memory
    *p = 10 // reference value to pointer
            // make sure pointer isn't nil before running code

    fmt.Printf("\nThe value p points to is: %v", *p) // reference value to pointer
    fmt.Printf("\nThe value of i is: %v", i)

    var thing1 = [5]float64{1,2,3,4,5}
    fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
    var result [5]float64 = square(&thing1)
    fmt.Printf("\nThe result is: %v", result)
}

// functions can take pointers of arrays to make changes to them
// this is to save memory by not having to copy large data through func
func square(thing2 *[5]float64) [5]float64{
    fmt.Printf("\nThe memory location of the thing2 array is: %p", thing2)
    for i := range thing2{
        thing2[i] = thing2[i]*thing2[i]
    }
    return *thing2
}
