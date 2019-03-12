package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

var password string

var cache map[string]bool

func initCache() {
	cache = make(map[string]bool)
}

var sessionTime = 24 * 60 * 60 * time.Second

func Signin(w http.ResponseWriter, r *http.Request) {
	/*
		u, ok1 := r.URL.Query()["user"]
		if !ok1 {
			w.Write([]byte(fmt.Sprintf("{'Error':'No User Found'}")))
			return
		}
	*/
	p, ok2 := r.URL.Query()["password"]
	if !ok2 {
		w.Write([]byte(fmt.Sprintf("{'Error':'No User Password Found'}")))
		return
	}
	if p[0] != password {
		w.Write([]byte(fmt.Sprintf("{'Error':'Wrong Password'}")))
		return

	}

	//Check User Status
	uNew, _ := uuid.NewV4()
	sessionToken := uNew.String()
	cache[sessionToken] = true
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(sessionTime), //TODO
	})
	go func() {
		time.Sleep(sessionTime)
		if _, ok := cache[sessionToken]; ok {
			delete(cache, sessionToken)
		}
	}()
}

func Signout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.Write([]byte(fmt.Sprintf("{'Error':'No Cookie'}")))
			return
		}
		w.Write([]byte(fmt.Sprintf("{'Error':'Other Error'}")))
		return
	}
	sessionToken := c.Value
	if _, ok := cache[sessionToken]; ok {
		delete(cache, sessionToken)
		w.Write([]byte(fmt.Sprintf("{'Info':'Session Deleted'}")))
	} else {
		w.Write([]byte(fmt.Sprintf("{'Error':'Session Does Not Exists'}")))

	}
}

var noSecureMap = map[string]bool{
	"/signout":   true,
	"/signin":    true,
	"/main.html": true,
}

func noSecure(url string) bool {
	_, ok := noSecureMap[url]
	return ok
}

func secureMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if noSecure(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.Write([]byte(fmt.Sprintf("{'Error':'No Cookie'}")))
				return
			}
			w.Write([]byte(fmt.Sprintf("{'Error':'Other Error'}")))
			return
		}
		sessionToken := c.Value

		response, ok := cache[sessionToken]
		if !ok {
			w.Write([]byte(fmt.Sprintf("No Session")))
			return
		}

		if !response {
			w.Write([]byte(fmt.Sprintf("No Reps")))
			return
		}
		next.ServeHTTP(w, r)
	})
}
func mainHtml(w http.ResponseWriter, r *http.Request) {
	s := `
	<html>
	<head>
	</head>
	<body>
	<h3>CMU Nucleome Data Server </h3>
	<form action="signin">
  Password:<br>
  <input type="text" name="password">
  <br>
  <input type="submit" value="Submit">
  </form>
	</body>
	</html>
	`
	w.Write([]byte(s))
}
