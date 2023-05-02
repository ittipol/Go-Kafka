package ex

import (
	"database/sql"
	"fmt"
	"migration/repository"

	"gorm.io/gorm"
)

func innerJoin(db *gorm.DB) {
	result := []repository.ItemToCategory{}

	db.Joins("Item").Joins("Category").Where("item_id=@itemId", sql.Named("itemId", 2)).Find(&result)

	fmt.Println(result)

	for _, v := range result {
		fmt.Printf("%v: %v | %v: %v\n", v.Item.ID, v.Item.Name, v.CategoryID, v.Category.Name)
	}
}

func joinWithConditionEx(db *gorm.DB) {
	result := []repository.ItemToCategory{}

	db.Joins("Item", db.Where(&repository.Item{Model: gorm.Model{ID: 2}})).Find(&result)
}
