package main

import (
	"consumer/repositories"
	"migration/orm/db"
)

func main() {

	//
	// dsn := "host=localhost user=postgres password=1234 dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Bangkok"

	// Mysql
	dsn := "root:1234@tcp(127.0.0.1:3306)/micro?charset=utf8mb4&parseTime=True&loc=Local"

	db := db.GetConnection(dsn, false)

	err := db.Migrator().CreateTable(repositories.BankAccount{})

	if err != nil {
		panic(err)
	}

	// err := db.Migrator().CreateTable(repository.ItemToCategory{})

	// if err != nil {
	// 	panic(err)
	// }

	// db.AutoMigrate(repository.Item{}, repository.Category{})
	// db.AutoMigrate(repository.Category{})

	// tx := db.Create(&repository.Category{
	// 	Name: "Code",
	// })

	// item := repository.Item{
	// 	Name: "item A",
	// }

	// db.Create(&item)

	// category := repository.Category{
	// 	Name: "Category A",
	// 	Slug: "category-a",
	// }

	// tx = db.Create(&category)

	// fmt.Println(tx.RowsAffected)

	// itemToCategory := repository.ItemToCategory{
	// 	ItemID:     item.ID,
	// 	CategoryID: category.ID,
	// }

	// db.Create(&itemToCategory)

	// result := []repository.ItemToCategory{}

	// db.Joins("Item").Joins("Category").Where("item_id=@itemId", sql.Named("itemId", 2)).Find(&result)

	// fmt.Println(result)

	// for _, v := range result {
	// 	fmt.Printf("%v: %v | %v: %v\n", v.Item.ID, v.Item.Name, v.CategoryID, v.Category.Name)
	// }
}

// func create(db *gorm.DB, value interface{})
