package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

/* TODO, Add Token Server, make data server get the correct user
 *       Add Parameters for Public Data Server and Private Data Server
 *       How to Close Web Query from Other Crawler?
 */
func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := r.Header.Get("Authorization")
		s1 := strings.Replace(s, "Basic ", "", 1)
		fmt.Println(s, s1)
		ue, err := base64.StdEncoding.DecodeString(s1)
		if err == nil {
			//TODO
			fmt.Println("User Middleware :", string(ue))
		} else {
			//TODO
			fmt.Println(err)
		}
		next.ServeHTTP(w, r)
	})
}
