package main

import (
	"net/http"
	"time"
)

func Racer(urlone, urltwo string) (winner string) {
	startone := time.Now()
	http.Get(urlone)
	startoneDuration := time.Since(startone)

	starttwo := time.Now()
	http.Get(urltwo)
	starttwoDuration := time.Since(starttwo)

	if startoneDuration < starttwoDuration {
		return urlone
	}

	return urltwo

}
