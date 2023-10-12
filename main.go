package main

import (
	router "github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/controller"
)

func main() {
	r := router.Get()
	r.Run()
}
