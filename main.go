package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ferozjilla/home/food/food"
)

func main() {
	http.HandleFunc("/ingredients/", ingredientsHandler)

	fmt.Printf("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ingredientsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": // return all ingredients
		carrot := food.Ingredient{
			Name:     "Carrot",
			Metadata: "Small cut",
			Expiry:   time.Now(),
		}
		byteCarrot, err := json.Marshal(carrot)
		if err != nil {
			w.Write([]byte("Oh no! Sorry."))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(byteCarrot)
		if err != nil {
			w.Write([]byte("Oh no! Sorry."))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case "POST": // add a new ingredient
		var bodyBytes []byte
		_, err := r.Body.Read(bodyBytes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var ingredient food.Ingredient
		err = json.Unmarshal(bodyBytes, &ingredient)
		if err != nil {
			w.Write([]byte("Could not unmarshal json\nThis is probably because of an incorrect request JSON"))
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
		fmt.Fprintf(w, "stored object successfully!")

	default: // return error
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
