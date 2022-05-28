package main

import "fmt"

func main() {
	var (
		l1 = []int{1, 2, 3}
		l2 = []int{4, 5}
		l3 = make([]int, len(l1)+len(l2))
	)
	copy(l3, l1)
	copy(l3[len(l1):], l2)
	fmt.Printf("%v", l3)

	//r := gin.Default()
	//r.GET("/hello", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"message": "hello",
	//	})
	//})
	//
	//r.Run()
}
