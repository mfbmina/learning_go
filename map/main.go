package main

import "fmt"

func main() {
  // var colors map[string]string

  // colors := make(map[int]string)

  colors := map[string]string {
    "red": "#ff0000",
    "green": "#4bf745",
  }

  colors["white"] = "#ffffff"

  // delete(colors, "white")

  fmt.Println(colors)

  printMap(colors)
}

func printMap(c map[string]string) {
  for color, hex := range c {
    fmt.Println(color, "in hex:", hex)
  }
}
