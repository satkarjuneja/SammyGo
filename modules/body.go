package modules

import (
	"fmt"
	"net/http"
	"io"
)

func ShowBody(resp *http.Response){
	var body, errBody = io.ReadAll(resp.Body)

	if errBody != nil {
		fmt.Printf("%sError loading request body!\n", ColorRed)
		return
	}
	
	fmt.Println("-------------BODY-------------")
	fmt.Println(string(body))
	fmt.Println("------------------------------")
}

