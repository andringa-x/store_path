package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/andringa-x/store_path/api"
	"github.com/andringa-x/store_path/internal/tools"
	log "github.com/sirupsen/logrus"
)

func PostSortList(w http.ResponseWriter, r *http.Request) {
	var storePathBody = api.StorePathBody{}

	err := json.NewDecoder(r.Body).Decode(&storePathBody)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	log.Info(storePathBody)
	aisles, aisleMap := tools.JsonMapToAisles("aisles.json")

	query := tools.QueryBuilder(aisles, aisleMap, storePathBody.Path, storePathBody.List)

	geminiReponse, err := api.GeminiCall(query)
	if err != nil {
		log.Error(err)
		api.GeminiErrorHandler(w)
		return
	}

	var response = api.StorePathReponse{
		List: geminiReponse,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
