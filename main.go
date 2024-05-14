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

	r.HandleFunc("/config", getConfig).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getConfig(w http.ResponseWriter, r *http.Request) {

	jsonData := getResponse("RemoteConfigIOSv2")

	var model RemoteConfigIOSv2Model

	err := json.Unmarshal([]byte(jsonData), &model)
	if err != nil {
		fmt.Println("Ошибка чтения JSON-данных:", err)
	}
	fmt.Println(model)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model)
}

// func main() {

// 	resp := getResponse("RemoteConfigIOSv2")
// 	fmt.Println(string(resp))
// }

func getResponse(fileName string) string {

	path := strings.Join([]string{"./Response/", fileName, ".json"}, "")

	content, error := os.ReadFile(path)

	if error != nil {
		fmt.Print(error)
	}

	fmt.Print(string(content))

	return content
}
