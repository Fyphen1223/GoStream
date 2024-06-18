package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("GoStream v0.0.1");
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/*", handle);
	fmt.Println("Server started on port 8080");
	fmt.Println("GoStream is ready to accept connection");
	server.ListenAndServe();
}

func handle(w http.ResponseWriter, r *http.Request) (){
	val := validateRequest(w, r);
	if !val {
		return;
	}

	path := r.URL.Path;

	if path == "/loadTracks" {
		loadTracks(w, r);
		return;
	}

	handle404(w);
}

func validateRequest(w http.ResponseWriter, r *http.Request) (bool){
	auth := r.Header.Get("auth");
	if auth != "123456" {
		w.WriteHeader(http.StatusUnauthorized);
		fmt.Fprint(w, "Unauthorized");
		return false;
	}
	return true;
};

func handle404 (w http.ResponseWriter) (){
	w.WriteHeader(http.StatusNotFound);
	fmt.Fprint(w, "Not Found");
};

func loadTracks(w http.ResponseWriter, r *http.Request) (){
	if r.Method != "GET" {
		handelIllegalMethod(w);
		return;
	}
	identifier := r.URL.Query().Get("identifier");
	fmt.Fprint(w, "This is your identifier:" + identifier);
}

func handelIllegalMethod(w http.ResponseWriter) (){
	w.WriteHeader(http.StatusMethodNotAllowed);
	fmt.Fprint(w, "Method Not Allowed");
}