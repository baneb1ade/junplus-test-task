package tests

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"javacode-test-task/app/internal/wallet"
	"javacode-test-task/app/internal/wallet/db"
	"javacode-test-task/app/pkg/client/psql"
	"javacode-test-task/app/pkg/logger"
	"os"
	"testing"
)

type TestStorageConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func loadTestConfig(t *testing.T) *TestStorageConfig {
	err := godotenv.Load("../../../../config.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	requiredEnvVars := []string{
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_HOST",
		"POSTGRES_PORT",
		"POSTGRES_DB",
		"SERVER_ADDRESS",
		"SERVER_PORT",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			t.Fatalf("%s must be set", envVar)
		}
	}
	return &TestStorageConfig{
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		DBName:     os.Getenv("POSTGRES_DB"),
	}
}

func TestService(t *testing.T) {
	cfg := loadTestConfig(t)
	if cfg.DBHost == "postgres" {
		cfg.DBHost = "localhost"
	}
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort, // Преобразуем int в строку
		cfg.DBName,
	)
	psqlClient, err := psql.NewClient(context.Background(), dsn)
	if err != nil {
		t.Fatal(err)
	}
	log := logger.SetupLogger()
	storage := db.NewRepository(psqlClient, log)

	t.Run("Get One Test", func(t *testing.T) {
		expectWallet := wallet.Wallet{
			UUID:    "a3c8a350-5b69-4d75-a16e-8d5bfa2b7a29",
			Balance: 150.75,
		}
		w, err := storage.GetOne(context.Background(), expectWallet.UUID)
		if err != nil {
			t.Fatal(err)
		}

		if w.UUID != expectWallet.UUID {
			t.Errorf("UUID mismatch: got %s, want %s", w.UUID, expectWallet.UUID)
		}
		if w.Balance != expectWallet.Balance {
			t.Errorf("Balance mismatch: got %.2f, want %.2f", w.Balance, expectWallet.Balance)
		}
	})

	t.Run("Get All Test", func(t *testing.T) {
		expectWallets := []wallet.Wallet{
			wallet.Wallet{
				UUID:    "a3c8a350-5b69-4d75-a16e-8d5bfa2b7a29",
				Balance: 150.75,
			},
			wallet.Wallet{
				UUID:    "bbd9c3f1-8a5f-4f3e-87e6-9c8b4a9d69c0",
				Balance: 2000,
			},
			wallet.Wallet{
				UUID:    "dde3f8e2-91a7-47fc-b09e-4f52934912a8",
				Balance: 750.25,
			},
			wallet.Wallet{
				UUID:    "cc7a9d85-f728-4c44-b55b-34e354f5937a",
				Balance: 500.5,
			},
		}
		ws, err := storage.GetAll(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if len(ws) != len(expectWallets) {
			t.Errorf("len(ws): got %d, want %d", len(ws), len(expectWallets))
		}
	})

	t.Run("Update One Test", func(t *testing.T) {
		walletUUID := "a3c8a350-5b69-4d75-a16e-8d5bfa2b7a29"
		var orgBalance float32 = 150.75
		w, err := storage.GetOne(context.Background(), walletUUID)
		if err != nil {
			t.Fatalf("Error in GetOne Err: %v", err)
		}
		w.Balance = 100
		err = storage.UpdateOne(context.Background(), w)
		if err != nil {
			t.Fatalf("Error in UpdateOne Err: %v", err)
		}
		updatedWallet, err := storage.GetOne(context.Background(), walletUUID)
		if err != nil {
			t.Fatalf("Error in GetOne Err: %v", err)
		}
		if updatedWallet.UUID != w.UUID {
			t.Fatalf("UUID mismatch: got %s, want %s", updatedWallet.UUID, w.UUID)
		}
		if updatedWallet.Balance != w.Balance {
			t.Fatalf("Balance mismatch: got %.2f, want %.2f", updatedWallet.Balance, w.Balance)
		}

		w.Balance = orgBalance
		err = storage.UpdateOne(context.Background(), w)
		if err != nil {
			t.Fatalf("Error in UpdateOne Err: %v", err)
		}
	})

}
