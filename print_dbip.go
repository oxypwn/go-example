package main

import "os"
import "fmt"

func main() {
   fmt.Println(os.Getenv("DB_PORT_8080_TCP_ADDR"))
}
