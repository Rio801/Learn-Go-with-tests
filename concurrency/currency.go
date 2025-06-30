package main

import (
	"net/http"
	"time"
)

func measureDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	ending := time.Since(start)
	return ending
}

func Racer(urlone, urltwo string) (winner string) {
	startoneDuration := measureDuration(urlone)
	starttwoDuration := measureDuration(urltwo)

	if startoneDuration < starttwoDuration {
		return urlone
	}

	return urltwo

}
