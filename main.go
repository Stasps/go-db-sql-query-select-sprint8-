package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale

	// Подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Выполнение SELECT-запроса
	rows, err := db.Query("SELECT product, volume, date FROM sales WHERE client = :client", sql.Named("client", client))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Заполнение массива в переменной sales объектами Sale
	for rows.Next() {
		var sale Sale // временный объект для текущей строки

		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			return nil, err
		}

		sales = append(sales, sale)
	}

	// далее, оригинальный прекод
	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
