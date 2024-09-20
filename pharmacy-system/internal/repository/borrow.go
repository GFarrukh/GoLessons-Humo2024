package repository

import (
	"fmt"
	"log"
	"pharmacy-system/internal/model"
	"time"
)

// IsMedicationAvailable проверяет, доступна ли книга для выдачи.
func (r *Repository) IsMedicationAvailable(MedicationID int) (bool, error) {
	var count int64
	// Проверяем, есть ли записи, где книга выдана и еще не возвращена
	// select count(*) from borrows where MedicationID = MedicationID AND return_date IS NULL
	result := r.db.Model(&model.Borrow{}).Where("MedicationID = ? AND return_date IS NULL", MedicationID).Count(&count)
	if result.Error != nil {
		log.Printf("IsMedicationAvailable: Failed to check availability: %v\n", result.Error)
		return false, fmt.Errorf("Failed to check availability: %v\n", result.Error)
	}

	// Если count == 0, значит книга доступна для выдачи
	return count == 0, nil
}

// AddBorrow добавляет новую запись о выдаче книги.
func (r *Repository) AddBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	// insert into borrows (user_id, book_id, borrow_date, return_date) values (user_id, book_id, borrow_date, return_date) returning borrow_id
	result := r.db.Create(&borrow)
	if result.Error != nil {
		log.Printf("AddBorrow: Failed to add borrow: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add borrow: %v\n", result.Error)
	}

	return borrow, nil
}

func (r *Repository) AddBorrowV2(borrow *model.Borrow) (*model.Borrow, error) {
	tx := r.db.Begin()

	// insert into borrows (user_id, book_id, borrow_date, return_date)
	// values
	// (user_id, book_id, borrow_date, return_date) returning borrow_id
	result := tx.Create(&borrow)
	if result.Error != nil {
		log.Printf("AddBorrow: Failed to add borrow: %v\n", result.Error)
		tx.Rollback()
		return nil, fmt.Errorf("Failed to add borrow: %v\n", result.Error)
	}

	// insert into borrow_history (user_id, book_id, borrow_date, return_date)
	// values
	// (user_id, book_id, borrow_date, return_date) returning borrow_id
	borrowHistory := model.BorrowHistory{
		BorrowID:   borrow.BorrowID,
		ActionType: "borrow",
	}

	result = tx.Create(&borrowHistory)
	if result.Error != nil {
		log.Printf("AddBorrow: Failed to add borrow history: %v\n", result.Error)
		tx.Rollback()
		return nil, fmt.Errorf("Failed to add borrow history: %v\n", result.Error)
	}

	tx.Commit()

	return borrow, nil
}

// GetBorrows возвращает все записи о выдаче книг.
func (r *Repository) GetBorrows() ([]model.Borrow, error) {
	// select * from borrows
	var borrows []model.Borrow
	result := r.db.Find(&borrows)
	if result.Error != nil {
		log.Printf("GetBorrows: Failed to get borrows: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get borrows: %v\n", result.Error)
	}

	return borrows, nil
}

// GetBorrowsByUser возвращает все записи о выдаче книг пользователя.
func (r *Repository) GetBorrowsByUser(userID int) ([]model.Borrow, error) {
	var borrows []model.Borrow

	// select * from borrows where user_id = userID
	result := r.db.Where("UserID = ?", userID).Find(&borrows)
	if result.Error != nil {
		log.Printf("GetBorrowsByUser: Failed to get borrows: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get borrows: %v\n", result.Error)
	}

	return borrows, nil
}

// GetBorrowsByBook возвращает все записи о выдаче книги.
func (r *Repository) GetBorrowsByMedication(MedicationID int) ([]model.Borrow, error) {
	var borrows []model.Borrow

	// select * from borrows where book_id = bookID
	result := r.db.Where("MedicationID = ?", MedicationID).Find(&borrows)
	if result.Error != nil {
		log.Printf("GetBorrowsByMedication: Failed to get borrows: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get borrows: %v\n", result.Error)
	}

	return borrows, nil
}

// GetBorrowByID возвращает запись о выдаче книги по идентификатору.
func (r *Repository) GetBorrowByID(borrowID int) (*model.Borrow, error) {
	var borrow model.Borrow

	// select * from borrows where borrow_id = borrowID
	result := r.db.Where("borrow_id = ?", borrowID).Find(&borrow)
	if result.Error != nil {
		log.Printf("GetBorrowByID: Failed to get borrow: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get borrow: %v\n", result.Error)
	}

	return &borrow, nil
}

// DeleteBorrow удаляет запись о выдаче книги по идентификатору.
func (r *Repository) DeleteBorrow(borrowID int) (int, error) {
	// delete from borrows where borrow_id = borrowID returning borrow_id
	result := r.db.Where("borrow_id = ?", borrowID).Delete(&model.Borrow{})
	if result.Error != nil {
		log.Printf("DeleteBorrow: Failed to delete borrow: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete borrow: %v\n", result.Error)
	}

	return borrowID, nil
}

// GetBorrowsByUserAndBook возвращает запись о выдаче книги по идентификатору.
func (r *Repository) GetBorrowsByUserAndMedication(UserID, MedicationID int) ([]model.Borrow, error) {

	var borrow []model.Borrow
	// select * from borrows where UserID = UserID and MedicationID = MedicationID
	result := r.db.Where("UserID = ? AND MedicationID = ?", UserID, MedicationID).Find(&borrow)
	if result.Error != nil {
		log.Printf("GetBorrowsByUserAndMedication: Failed to get borrow: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get borrow: %v\n", result.Error)
	}

	return borrow, nil
}

// ReturnBook возвращает книгу в библиотеку.
func (r *Repository) ReturnMedication(borrowID int) error {
	// update borrows set return_date = now() where borrow_id = borrowID
	result := r.db.Model(&model.Borrow{}).Where("borrow_id = ?", borrowID).Update("return_date", time.Now())
	if result.Error != nil {
		log.Printf("ReturnMedication: Failed to return medication: %v\n", result.Error)
		return fmt.Errorf("Failed to return medication: %v\n", result.Error)
	}

	return nil
}

func (r *Repository) ReturnMedicationV2(borrowID int) error {
	tx := r.db.Begin()

	// update borrows set return_date = now() where book_id = bookID
	result := tx.Model(&model.Borrow{}).Where("borrow_id = ?", borrowID).Update("return_date", time.Now())
	if result.Error != nil {
		log.Printf("ReturnMedicationV2: Failed to return medication: %v\n", result.Error)
		tx.Rollback()
		return fmt.Errorf("Failed to return medication: %v\n", result.Error)
	}

	// insert into borrow_history (borrow_id, action_type, action_date) values (borrowID, "return", now())
	now := time.Now()

	borrowHistory := model.BorrowHistory{
		BorrowID:   borrowID,
		ActionType: "return",
		ActionDate: &now,
	}

	result = tx.Create(&borrowHistory)
	if result.Error != nil {
		log.Printf("ReturnMedicationV2: Failed to create borrow history: %v\n", result.Error)
		tx.Rollback()
		return fmt.Errorf("Failed to create borrow history: %v\n", result.Error)
	}

	err := tx.Commit().Error
	if err != nil {
		log.Printf("ReturnMedicationV2: Failed to commit transaction: %v\n", err)
		tx.Rollback()
		return fmt.Errorf("Failed to commit transaction: %v\n", err)
	}

	return nil
}
