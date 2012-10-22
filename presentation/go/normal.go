package main

import (
   "fmt"
   "time"
)

func main() {
   fmt.Println("Timer 1: 3s")
   time.Sleep(3 * 1e9)
   fmt.Println("Timer 1 done")

   fmt.Println("Timer 2: 1s")
   time.Sleep(1 * 1e9)
   fmt.Println("Timer 2 done")
}
