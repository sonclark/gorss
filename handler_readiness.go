package main

import "net/http"

//this is a specific function signature. Have to use if you want to define a http handler in the way that go standard library expect
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
