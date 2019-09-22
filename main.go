package main

import (
	"fmt"

	gormdao "github.com/CRL-Studio/AuthServer/src/dao/gorm"
	"github.com/CRL-Studio/AuthServer/src/route"
)

func main() {
	//defer logger.Close()
	defer gormdao.Close()
	fmt.Printf("Auth Service Start\n")
	route.Run()
}
