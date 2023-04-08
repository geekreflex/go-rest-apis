package main

import (
	"blog-api/controller"
	"blog-api/model"
)



func main() {
	model.Init()
	controller.Start()
}