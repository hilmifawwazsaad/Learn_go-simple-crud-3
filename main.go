package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-crud-3/database"
	"simple-crud-3/handlers"
	"simple-crud-3/repositories"
	"simple-crud-3/services"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONNECTION"`
}

func main() {
	viper.AutomaticEnv()

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONNECTION"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Setup
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	// Setup routes
	http.HandleFunc("/api/category", categoryHandler.HandleCategories)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)

	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout)

	http.HandleFunc("/api/report", reportHandler.HandleReport)
	http.HandleFunc("/api/report/hari-ini", reportHandler.HandleReportToday)

	// /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Welcome to the API",
		})
	})

	// api/health
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API is running",
		})
	})

	// Start server
	log.Println("Server started on :" + config.Port)
	fmt.Println("Server started on :" + config.Port)

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
