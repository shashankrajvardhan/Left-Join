package main

import (
	"fmt"
	"log"
)

func join() {
	query := `
	SELECT orders.orderid, orders.customerid, customer.customername, orders.orderdate
	FROM orders
	LEFT JOIN customer ON orders.customerid = customer.customerid;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			orderdate           string
			customername        any
			orderid, customerid int
		)

		err := rows.Scan(&orderid, &customerid, &customername, &orderdate)
		if err != nil {
			log.Fatal(err)
		}
		if customername == nil {
			customername = ""
		}
		fmt.Printf("orderid: %d, customerid: %d, customername: %s, orderdate: %s\n", orderid, customerid, customername, orderdate)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
