// main single package for api, we want to keep everything in the same package for simplicity
package main

func main() {
    server := NewAPIServer(":3000")
    server.Run()
}
