package main

import pb "go-proto-example/src/protos/snazzy/items"

import "fmt"

func main() {
    shirt := &pb.Shirt{ Color: "red", Size: pb.Shirt_MEDIUM}
    fmt.Println("Hello shirt ", shirt.String())
}
