package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/RemoteConfigIOSv2", getConfig).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getConfig(w http.ResponseWriter, r *http.Request) {

	jsonMap := getResponse(r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonMap)
}

func getResponse(fileName string) map[string]interface{} {

	path := strings.Join([]string{"./Models", fileName, ".json"}, "")

	content, error := os.ReadFile(path)

	if error != nil {
		fmt.Print(error)
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(content)), &jsonMap)

	return jsonMap
}
