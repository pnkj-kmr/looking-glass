package main

// func main() {
// 	test_x()
// }

// func test_x() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)
// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			time.Sleep(1 * time.Second)
// 			c1 <- time.Now().String()
// 		}
// 		close(c1)
// 	}()
// 	go func() {
// 		for i := 0; i < 3; i++ {
// 			time.Sleep(2 * time.Second)
// 			c2 <- time.Now().String()
// 		}
// 		close(c2)
// 	}()
// 	var res1, res2 string
// 	ch1_close, ch2_close := true, true
// 	for {
// 		select {
// 		case res1, ch1_close = <-c1:
// 			if ch1_close {
// 				fmt.Println("from c1:", res1, ch1_close)
// 			}
// 		case res2, ch2_close = <-c2:
// 			if ch2_close {
// 				fmt.Println("from c2:", res2, ch2_close)
// 			}
// 		}
// 		if (!ch1_close) && (!ch2_close) {
// 			fmt.Println("CLOSING....... 2000000")
// 			break
// 		}
// 	}
// }
