package task2

import (
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	ID      int     `gorm:"primaryKey;autoIncrement;column:id"`
	Name    string  `gorm:"column:name;not null"`
	Balance float64 `gorm:"column:balance;not null"`
}

type Transactions struct {
	ID            int     `gorm:"primaryKey;autoIncrement;column:id"`
	FromAccountId int     `gorm:"column:from_account_id;not null"`
	ToAccountId   int     `gorm:"column:to_account_id;not null"`
	Amount        float64 `gorm:"column:amount;not null"`
}

func initAccount(d *gorm.DB, from, to string) {
	fmt.Println("init account and transactions table")
	// Create the tables if they do not exist
	if err := d.Debug().AutoMigrate(&Account{}, &Transactions{}); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	accounts := []Account{
		{Name: from, Balance: 1000.0},
		{Name: to, Balance: 1000.0},
	}

	if err := d.Debug().Create(&accounts).Error; err != nil {
		panic(fmt.Sprintf("failed to create accounts: %v", err))
	}

	fmt.Println("account and transactions table initialized successfully")
}

// 假设转账账户与被转账户没有重名
func Transfer(db *gorm.DB, fromAccountName, toAccountName string, amount float64) error {
	initAccount(db, fromAccountName, toAccountName)
	// Check if the accounts exist and have sufficient balance
	var fromAccount Account
	var toAccount Account

	db.Transaction(func(tx *gorm.DB) error {
		db.Debug().Where("name=?", fromAccountName).First(&fromAccount)
		db.Debug().Where("name=?", toAccountName).First(&toAccount)

		fromAccount.Balance -= amount
		toAccount.Balance += amount

		transactionRecord := Transactions{
			FromAccountId: fromAccount.ID,
			ToAccountId:   toAccount.ID,
			Amount:        amount,
		}

		accounts := []Account{fromAccount, toAccount}

		if err := tx.Debug().Save(&accounts).Error; err != nil {
			return fmt.Errorf("failed to update accounts: %w", err)
		}

		if err := tx.Debug().Create(&transactionRecord).Error; err != nil {
			return fmt.Errorf("failed to record transaction: %w", err)
		}

		return tx.Debug().Commit().Error
	})

	return nil
}
