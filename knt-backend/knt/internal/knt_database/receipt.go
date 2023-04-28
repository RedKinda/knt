package kntdatabase

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

//This file handles everything having to do with purchases and receipts

func MakeTransaction(userId int, purchase PurchaseRequest, db *sql.DB) (int, error) {
	//Begins the transaction
	//This is important because if ANY error were to occur we need to reset the database to its original state
	transaction, err := db.Begin()
	if err != nil {
		return 0, err
	}
	// get user
	user, err := GetUser(db, userId)
	if err != nil {
		return 0, err
	}
	//pin validation
	if !ValidatePin(purchase.Password, user, db) {
		return 0, errors.New("Unauthorized")
	}
	//calculate cost
	cost, err := calculateCost(purchase.Data, db)
	if err != nil {
		return 0, err
	}
	//validate users balance
	if cost > user.Balance {
		return 0, errors.New("Insufficient balance on user")
	}
	//generate receipt
	receiptId, err := generateReceipt(transaction, purchase.Data)
	if err != nil {
		return 0, err
	}
	//make a transaction
	err = generateTransaction(transaction, userId, user.Balance, cost, user.Balance-cost, receiptId)
	if err != nil {
		return 0, err
	}
	//subtract balance
	err = setBalance(transaction, userId, user.Balance-cost)
	if err != nil {
		return 0, err
	}

	//Add tax
	err = addTaxTotals(purchase.Data, db, transaction, cost)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	//Finilize the transaction
	err = transaction.Commit()
	if err != nil {
		return 0, err
	}
	return cost, nil
}

// Calculates the total cost of the purchased products
func calculateCost(entries []PurchaseEntry, db *sql.DB) (int, error) {
	var cost int

	for _, element := range entries {
		value, err := getProductValue(element, db)
		if err != nil {
			return 0, err
		}
		cost += value
	}

	return cost, nil
}

// returns the value of a specific entry
func getProductValue(entry PurchaseEntry, db *sql.DB) (int, error) {
	value, err := getFirstSingleValue[int](queryBuilder(db, "select price from product where id = ?", entry.ProductId))
	return value * entry.Amount, err
}

// Generated the receipt and stores it in the database
func generateReceipt(transaction *sql.Tx, entries []PurchaseEntry) (int64, error) {
	dataString, err := json.Marshal(entries)
	if err != nil {
		return 0, err
	}

	receiptId, err := addToTransaction(transaction, "INSERT INTO receipts (data, timestamp) VALUES (?, datetime())", dataString)
	if err != nil {
		return 0, err
	}

	return receiptId, nil
}

// Generates the transaction and stores it in the database
func generateTransaction(transaction *sql.Tx, userId int, startingBal int, deltaBal int, finalBal int, receiptId int64) error {
	_, err := addToTransaction(transaction,
		"INSERT INTO transactions (user_id, starting_balance, delta_balance, final_balance, receipt_id) VALUES (?, ?, ?, ?, ?)",
		userId, startingBal, deltaBal, finalBal, receiptId)
	return err
}

// Sets the users balance to a specified ammount
func setBalance(transaction *sql.Tx, userId int, balance int) error {
	_, err := addToTransaction(transaction, "UPDATE user SET balance = ? WHERE id = ?",
		balance, userId)
	return err
}

// Function designed to calculate total ammounts for tax reasons
// The taxes are stored in their own tables in the form of tax{year}
func addTaxTotals(entries []PurchaseEntry, db *sql.DB, transaction *sql.Tx, cost int) error {
	year := time.Now().Year()

	//Ensure yearly tables existance
	err := ensureTaxTableExists(transaction, year)
	if err != nil {
		return err
	}
	//Go through all the entries and apply the operation on all of them
	for _, entry := range entries {
		//check if entry exists
		tax, err := getFirstEntry[TaxEntry](queryBuilder(db, "SELECT * FROM tax where year = ? and product_id = ?",
			year, entry.ProductId))
		if err != nil {
			return err
		}
		if tax.Id == 0 {
			//The row for this product doesnt already exist in the table so create a new row for it
			_, err = addToTransaction(transaction, "INSERT INTO tax (product_id, amount, totalCost, year) VALUES (?, ?, ?, ?)",
				entry.ProductId, entry.Amount, cost, year)
			if err != nil {
				return err
			}
			continue
		}
		//The product exists update it
		_, err = addToTransaction(transaction, "UPDATE tax SET amount = ?, totalCost = ? WHERE year = ? and id = ?",
			tax.Amount+entry.Amount, tax.TotalCost+cost, year, tax.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func ensureTaxTableExists(transaction *sql.Tx, year int) error {
	_, err := addToTransaction(transaction, "CREATE TABLE IF NOT EXISTS tax (id INTEGER PRIMARY KEY AUTOINCREMENT, product_id INT, amount INT, totalCost INT, year INT)")
	return err
}
