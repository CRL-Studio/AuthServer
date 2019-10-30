package main

import (
	"fmt"

	gormdao "github.com/CRL-Studio/AuthServer/src/dao/gorm"
	"github.com/CRL-Studio/AuthServer/src/route"
	"github.com/CRL-Studio/AuthServer/src/utils/env"
	"github.com/CRL-Studio/AuthServer/src/utils/logger"
)

func main() {
	defer logger.Close()
	defer gormdao.Close()
	defer env.Set()
	fmt.Printf("Auth Service Start\n")
	route.Run()
}
