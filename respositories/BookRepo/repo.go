package BookRepo

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"
)

func Index() ([]models.Book, error) {
	var book []models.Book
	result := helpers.DB.Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func Store(book *models.Book) (*models.Book, error) {
	result := helpers.DB.Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}
func Show(id any) (*models.Book, error) {
	var book models.Book
	result := helpers.DB.Where("book_id=?", id).First(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}
func Update(data models.Book, id any) (*models.Book, error) {
	var book models.Book
	result := helpers.DB.Where("book_id = ?", id).First(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	helpers.DB.Model(&book).Updates(&data)
	return &book, nil
}
func Destroy(id any) error {
	var book models.Book
	result := helpers.DB.Where("book_id=?", id).First(&book)
	if result.Error != nil {
		return result.Error
	}
	helpers.DB.Delete(&book)
	return nil
}
