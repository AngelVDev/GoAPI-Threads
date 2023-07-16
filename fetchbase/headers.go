package fetchbase

import "net/http"

func initHeaders() http.Header {
	header := http.Header{}
	header.Add("content-type", "application/x-www-form-urlencoded")
	header.Add("user-agent", "Threads API midu client")
	header.Add("x-ig-app-id", THREADS_APP_ID)
	header.Add("x-fb-lsd", "jdFoLBsUcm9h-j90PeanuC")
	return header
}
