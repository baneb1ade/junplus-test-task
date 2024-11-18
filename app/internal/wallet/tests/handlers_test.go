package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"

	//"github.com/go-playground/validator/v10"
	"javacode-test-task/app/internal/wallet"
	"testing"
)

type MockWalletService struct{}

func TestHandlers(t *testing.T) {
	walletService := &MockWalletService{}
	v := validator.New()
	getAllWalletsHandler := wallet.GetAllWallets(walletService)
	findWalletByUUIDHandler := wallet.FindWalletByUUID(walletService)
	updateWalletByUUIDHandler := wallet.UpdateWalletByUUID(v, walletService)

	t.Run("Get all wallets", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/wallets/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := httprouter.New()
		router.GET("/api/v1/wallets/", getAllWalletsHandler)
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("Find wallet by UUID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/wallets/:id/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := httprouter.New()
		router.GET("/api/v1/wallets/:id/", findWalletByUUIDHandler)
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("Update wallet by UUID", func(t *testing.T) {
		payload := wallet.UpdateWalletRequest{
			UUID:          "a3c8a350-5b69-4d75-a16e-8d5bfa2b7a29",
			OperationType: "deposit",
			Amount:        100,
		}
		marshaled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/wallet/", bytes.NewBuffer(marshaled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := httprouter.New()
		router.POST("/api/v1/wallet/", updateWalletByUUIDHandler)
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})
}

func (mw *MockWalletService) GetWallets(ctx context.Context) ([]wallet.Wallet, error) {
	return nil, nil
}

func (mw *MockWalletService) GetWalletByUUID(ctx context.Context, uuid string) (wallet.Wallet, error) {
	return wallet.Wallet{}, nil
}

func (mw *MockWalletService) ChangeBalance(ctx context.Context, update wallet.UpdateWalletRequest) (wallet.Wallet, error) {
	return wallet.Wallet{}, nil
}
