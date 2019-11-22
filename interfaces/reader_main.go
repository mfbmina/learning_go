package main

import (
  "fmt"
  "io"
  "os"
)

type fileLogger struct{}

func main() {
  file, err := os.Open(os.Args[1])

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fl := fileLogger{}
  io.Copy(fl, file)
}

func (fileLogger) Write(bs []byte) (int, error) {
  fmt.Println(string(bs))
  return len(bs), nil
}
