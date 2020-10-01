package app

import (
	"fmt"
	"net/http"

	"github.com/kr/pretty"
)

func HandleError(res http.ResponseWriter, err error) {
	fmt.Println("=== ERROR")
	pretty.Println(err)
	res.WriteHeader(500)
	res.Write([]byte(err.Error()))
}
