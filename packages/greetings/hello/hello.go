package main

import (
	// "context"
	"os"
	"strings"
	"fmt"
	"github.com/techswarn/tempconv"

	// _ "github.com/go-sql-driver/mysql"
	// "github.com/xo/dburl"
)

func Main(args map[string]interface{}) map[string]interface{} {
	fmt.Println(os.Args)
	fmt.Println(strings.ToLower("WE LOVE DO"))
	tempconv.CTof(34)


	name, ok := args["name"].(string)
	if !ok {
		name = "stranger"
	}
	msg := make(map[string]interface{})
	msg["body"] = "Hello " + name + "!"
	return msg
}
