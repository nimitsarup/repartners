package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/nimitsarup/rep/handlers"
)

type API struct {
	Handlers handlers.HandlersInterface
}

func (a *API) UpdatePacks(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) <= 0 {
		http.Error(w, fmt.Sprintf("empty request body %v", err), http.StatusBadRequest)
		return
	}
	log.Printf("invoking UpdatePacks with [%s]", body)
	res, err := convertCommaSeparatedStringToIntArray(string(body))
	if err != nil {
		http.Error(w, fmt.Sprintf("bad request body %v", err), http.StatusBadRequest)
		return
	}
	err = a.Handlers.UpdatePacks(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("internal error %v", err), http.StatusInternalServerError)
		return
	}
}

func (a *API) GetPacksForItems(w http.ResponseWriter, r *http.Request) {
	itemsStr := r.URL.Query().Get("items")
	if itemsStr == "" {
		http.Error(w, "Please provide the order size using the 'items' query parameter.", http.StatusBadRequest)
		return
	}
	numItems, err := strconv.Atoi(itemsStr)
	if err != nil {
		http.Error(w, "Invalid order size. Please provide a valid integer.", http.StatusBadRequest)
		return
	}
	log.Printf("invoking GetPacksForItems with [%d]", numItems)
	if err := writeResponse(w, a.Handlers.GetPacksForItems(numItems)); err != nil {
		log.Println(err)
	}
}

func writeResponse(w http.ResponseWriter, resp interface{}) error {
	b, err := json.Marshal(resp)
	if err != nil {
		log.Printf("caught marshal error %v", err)
		return err
	}
	if _, err = w.Write(b); err != nil {
		log.Printf("caught write error %v", err)
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return nil
}

func convertCommaSeparatedStringToIntArray(s string) ([]int, error) {
	var result []int
	substrings := strings.Split(s, ",")
	for _, substring := range substrings {
		intVal, err := strconv.Atoi(strings.TrimSpace(substring))
		if err != nil {
			// Return the current result and the error if conversion fails
			return result, err
		}
		result = append(result, intVal)
	}
	return result, nil
}
