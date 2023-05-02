package repository

type Category struct {
	ID   uint
	Name string
	Slug string `gorm="unique;not null;index;size:100"`
}
