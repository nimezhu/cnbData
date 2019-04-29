package main

import (
	"strings"

	"github.com/rs/cors"
)

var (
	CORS = []string{
		"https://vis.nucleome.org",
		"http://vis.nucleome.org",
		"https://browser.nucleome.org",
		"http://browser.nucleome.org",
		"https://v.nucleome.org",
		"http://v.nucleome.org",
		"http://www.nucleome.org",
		"https://www.nucleome.org",
		"http://nucleome.org",
		"https://nucleome.org",
		"https://nbrowser.github.io",
		"https://genome.compbio.cs.cmu.edu",
		"http://genome.compbio.cs.cmu.edu:8080",
		"chrome-extension://djcdicpaejhpgncicoglfckiappkoeof",
		/* for development */
		"http://x7.andrew.cmu.edu:8080",
		"https://dev.nucleome.org",
		"http://dev.nucleome.org",
		"chrome-extension://gedcoafficobgcagmjpplpnmempkcpfp",
		"https://vbio.app",
	}
)

func getCors(customCors string) cors.Options {
	if customCors != "" {
		otherCors := strings.Split(customCors, ";")
		for _, s := range otherCors {
			CORS = append(CORS, s)
		}

	}
	corsOptions := cors.Options{
		AllowedOrigins:   CORS,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization"},
	}
	return corsOptions

}
