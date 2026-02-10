package modules

import (
        "fmt"
        "net/http"
				"strings"
)

func ShowHeader(resp *http.Response) {
        fmt.Println("----------HEADER----------")
        for k, v := range resp.Header {
                fmt.Printf("%s%s: %s%s\n", ColorGreen, k, ColorYellow, strings.Join(v, ", "))
        }
				fmt.Printf("%s", ColorReset)
        fmt.Println("--------------------------")
}
