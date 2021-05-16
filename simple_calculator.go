package main

import "fmt"

func main() {
  fmt.Println(sum(1,1));
  fmt.Println(sub(2,1));
  fmt.Println(div(4,2));
  fmt.Println(multi(2,2));
}

func sum(x int, y int) int {
  return x + y;
}

func sub(x int, y int) int {
  return x - y;
}

func div(x int, y int) int {
  return x / y;
}

func multi(x int, y int) int {
  return x * y;
}
