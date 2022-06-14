package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/FMusk99/gorm/model/order"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

type OrderRepository interface {
	GetOrders() order.EcomOrder
}

type orderRepository struct {
	db *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: Db,
	}
}

func (or *orderRepository) GetOrders() order.EcomOrder {
	var order order.EcomOrder
	or.db.First(&order, 1)

	return order
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	or := NewOrderRepository(Db)
	order := or.GetOrders()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	par, _ := json.Marshal(order)
	w.Write([]byte(par))
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "ecom_demo:Va46Q8ECtueGb5dTW@tcp(10.12.30.11:3306)/ecom_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db = db
	// Read
	r := mux.NewRouter()
	r.HandleFunc("/orders", getOrders)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
