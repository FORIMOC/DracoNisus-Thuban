package main

import (
	"fmt"

	"forimoc.DracoNisus-Thuban/util"
	"gorm.io/gorm"
)

func main() {
	a := util.HandleNotFoundErr(gorm.ErrDuplicatedKey, nil)
	fmt.Println(a)
}
