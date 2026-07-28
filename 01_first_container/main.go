package main

import "fmt"

func main(){
	fmt.Println("Hello from Docker - updated code !!")
}

//docker build -t first-image:v1 -- now check in docker desktop image must be created in image folder
//make some changes in main.go file build second image 
//docker build -t first-image:v2 -- now you can see v1 and v2 both images in docker desktop images folder 