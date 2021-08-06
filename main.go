package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/keriwisnu/dblocking/models"
	"github.com/keriwisnu/dblocking/sale"
	"github.com/keriwisnu/dblocking/utils"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/sale/update", UpdateSales)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

//Db row locking
func UpdateSales(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "content type application / json only", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var fsale models.Sales

		if err := json.NewDecoder(r.Body).Decode(&fsale); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := fsale.Update(ctx, fsale); err != nil {
			error := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, error, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Berhasil Membeli",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Gagal Membeli", http.StatusMethodNotAllowed)
	return
}

