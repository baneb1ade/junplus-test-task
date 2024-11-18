package wallet

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetAllWallets(s Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		wallets, err := s.GetWallets(context.Background())
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			sendJSONError(w, http.StatusOK, SWWErr, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(wallets); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"status": "something went wrong"})
			return
		}
	}
}

func FindWalletByUUID(s Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		wallet, err := s.GetWalletByUUID(context.Background(), p.ByName("id"))
		if err != nil {
			sendJSONError(w, http.StatusOK, SWWErr, err)
			return
		}
		response, err := json.Marshal(wallet)
		if err != nil {
			sendJSONError(w, http.StatusOK, SWWErr, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func UpdateWalletByUUID(v *validator.Validate, s Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req UpdateWalletRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			sendJSONError(w, http.StatusBadRequest, validationErr, err)
			return
		}

		if err := v.Struct(req); err != nil {
			sendJSONError(w, http.StatusBadRequest, validationErr, err)
			return
		}
		wallet, err := s.ChangeBalance(context.Background(), req)
		if err != nil {
			sendJSONError(w, http.StatusOK, "Error", err)
			return
		}

		response, err := json.Marshal(wallet)
		if err != nil {
			sendJSONError(w, http.StatusOK, SWWErr, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func sendJSONError(w http.ResponseWriter, status int, err string, details error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{err: details.Error()})

}
