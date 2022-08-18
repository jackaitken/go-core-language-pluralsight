// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	healthHandler := func(w http.ResponseWriter, req *http.Request) {
// 		w.Header().Add("Content-Type", "applicaton/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"status": "OK"}`))
// 	}

// 	pingHandler := func(w http.ResponseWriter, req *http.Request) {
// 		w.Header().Add("Content-Type", "applicaton/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"status:OK", "data": {"message": "pong"} }`))
// 	}

// 	greetUserHandler := func(w http.ResponseWriter, req *http.Request) {
// 		names := req.URL.Query()["name"]
// 		var name string
// 		if len(names) == 1 {
// 			name = names[0]
// 		}

// 		m := map[string]string{"name": name}
// 		enc := json.NewEncoder(w)
// 		enc.Encode(m)

// 		w.Header().Add("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 	}

// 	http.HandleFunc("/", greetUserHandler)
// 	http.HandleFunc("/health", healthHandler)
// 	http.HandleFunc("/ping", pingHandler)

// 	err := http.ListenAndServe(":3000", nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
