package main

// import (
// 	"flag"
// 	"fmt"

// 	"github.com/pnkj-kmr/looking-glass/controllers"
// )

// func main() {
// 	var src, dst string
// 	var proto int

// 	flag.IntVar(&proto, "p", 1, "Protocol Values [1-6]")
// 	flag.StringVar(&src, "s", "127.0.0.1", "Source IP")
// 	flag.StringVar(&dst, "d", "127.0.0.1", "Destination IP")
// 	flag.Parse()

// 	out := make(chan []byte)
// 	err := controllers.Execute(src, dst, proto, out)
// 	if err != nil {
// 		fmt.Println("......", err.Error())
// 	}

// 	fmt.Println("starting...")
// 	for range out {
// 	}
// 	fmt.Println("exiting...")
// }
