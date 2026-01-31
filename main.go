package main

import (
	"encoding/json"
	"fmt"
	"kasir-go-api/database"
	"kasir-go-api/handlers"
	"kasir-go-api/repositories"
	"kasir-go-api/services"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// ubah Config
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

// Simple logging middleware
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		go log.Printf("START: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
		go log.Printf("END: %s %s %v", r.Method, r.URL.Path, time.Since(start))
	}
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Run database migrations
	if err := database.RunMigrations(db, "database/migrations"); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Setup routes with logging
	http.HandleFunc("/api/products", logRequest(productHandler.HandleProducts))
	http.HandleFunc("/api/products/", logRequest(productHandler.HandleProductByID))

	//health check with logging
	http.HandleFunc("/health", logRequest(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(map[string]string{
			"status":  "ok",
			"message": "Server is running",
		})
	}))

	err = http.ListenAndServe(":"+config.Port, nil)

	fmt.Println("Server running at port " + config.Port)

	if err != nil {
		fmt.Println("Failed to run server", err)
	}
}
