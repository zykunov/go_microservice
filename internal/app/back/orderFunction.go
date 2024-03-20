package back

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zykunov/courseGoFirst/WB0/storage"
)

type CacheStruct struct {
	Cache map[string]string
}

func (c *CacheStruct) AllItemsPage(w http.ResponseWriter, r *http.Request) {

	var keys []string

	cache := c.Cache
	for key, _ := range cache {
		keys = append(keys, key)
	}

	tpl, _ := template.ParseFiles("templates/allOrders.html")
	tpl.Execute(w, keys)
}

func (c *CacheStruct) ItemPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	cache := c.Cache
	rawData := cache[vars["key"]]

	// rawData := `{"a":"v54vw45", "somethingElse": [1234, 5678]}`
	m := map[string]interface{}{}
	if err := json.Unmarshal([]byte(rawData), &m); err != nil {
		panic(err)
	}

	tpl, _ := template.ParseFiles("templates/order.html")
	tpl.Execute(w, m)

	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(rawData))
}

func GetAllOrders(db *sql.DB) {
	query := fmt.Sprintf("SELECT id, order_uid FROM %s", "orders")
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("error selecting from BD", err)
	}
	defer rows.Close()
}

func GetCache(db *sql.DB) (m map[string]string) {

	query := fmt.Sprintf("SELECT order_uid, data FROM %s", "orders")
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("error selecting from BD", err)
	}
	defer rows.Close()

	cacheMap := make(map[string]string)

	for rows.Next() {
		order := storage.Order{}
		if err := rows.Scan(&order.OrderUid, &order.Data); err != nil {
			fmt.Println("error while scaning rows", err)
		}
		cacheMap[order.OrderUid] = order.Data
	}
	if len(cacheMap) == 0 {
		fmt.Println("There is no cache in DB, waiting for messages")
	} else {
		fmt.Println("Get cache from DB")
	}

	return cacheMap
}
