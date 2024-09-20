package service

import (
	"errors"
	"fmt"
	"pharmacy-system/internal/model"
	"time"
)

var (
	ErrMedicationNotAvailable    = errors.New("medication not available")
	ErrBorrowsNotFound           = errors.New("no borrows found")
	ErrBorrowNotFound            = errors.New("borrow not found")
	ErrMedicationAlreadyReturned = errors.New("medication already returned")
	ErrDifferentUser             = errors.New("different user")
	ErrUserNotFound              = errors.New("user not found")
)

func (s *Service) CreateBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	// Проверяем, существует ли пользователь
	_, err := s.Repository.GetUserByID(borrow.UserID)
	if err != nil {
		if errors.As(err, &ErrRecordNotFound) {
			//return nil, fmt.Errorf("user with id %d doesn't exist", borrow.UserID)
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Проверяем, имеется ли книга в наличии
	MedicationAvailable, err := s.Repository.IsMedicationAvailable(borrow.MedicationID)
	if err != nil {
		return nil, err
	}
	// Проверяем, доступна ли книга для выдачи
	if !MedicationAvailable {
		return nil, ErrMedicationNotAvailable
	}
	// Установка даты выдачи и возврата
	now := time.Now()
	borrow.BorrowDate = &now

	return s.Repository.AddBorrow(borrow)
}

func (s *Service) ReturnMedication(UserID int, borrowID int) error {
	// Получаем информацию о выдаче книги по ID
	borrowByID, err := s.Repository.GetBorrowByID(borrowID)
	if err != nil {
		if errors.As(err, &ErrRecordNotFound) {
			return ErrBorrowNotFound
		}
		return err
	}

	// Проверяем, вернулась ли книга
	if borrowByID.ReturnDate != nil {
		//return fmt.Errorf("Medication already returned with ID %d", borrowID)
		return ErrMedicationAlreadyReturned
	}
	// Получаем информацию о книге по ID
	_, err = s.Repository.GetMedicationByBorrow(borrowID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return fmt.Errorf("Medication not found with borrow ID %d", borrowID)
		}
		return err
	}

	// Проверяем, существует ли пользователь
	//_, err = s.Repository.GetUserByID(borrowByID.UserID)
	//if err != nil {
	//	if errors.As(err, &ErrRecordNotFound) {
	//		return fmt.Errorf("user not found with ID %d", borrowByID.UserID)
	//	}
	//	return err
	//}
	if UserID != borrowByID.UserID {
		return ErrDifferentUser
	}

	return s.Repository.ReturnMedication(borrowID)
}

func (s *Service) GetBorrows() ([]model.Borrow, error) {
	borrows, err := s.Repository.GetBorrows()
	if err != nil {
		return nil, err
	}
	if len(borrows) == 0 {
		return nil, ErrBorrowsNotFound
	}

	return borrows, nil
}

func (s *Service) GetBorrowByID(borrowID int) (*model.Borrow, error) {
	borrowByID, err := s.Repository.GetBorrowByID(borrowID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("borrow not found with ID %d", borrowID)
		}
		return nil, err
	}

	return borrowByID, nil
}

func (s *Service) GetBorrowsByUser(userID int) ([]model.Borrow, error) {
	// Проверяем, существует ли пользователь
	_, err := s.Repository.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with ID %d", userID)
		}
		return nil, err
	}

	// Получаем все выдачи пользователя
	borrowsByUser, err := s.Repository.GetBorrowsByUser(userID)
	if err != nil {
		return nil, err
	}
	if len(borrowsByUser) == 0 {
		return nil, fmt.Errorf("no borrows found with user ID %d", userID)
	}

	return borrowsByUser, nil
}

func (s *Service) GetBorrowsByMedication(MedicationID int) ([]model.Borrow, error) {
	// Проверяем, существует ли книга
	_, err := s.Repository.GetMedicationByID(MedicationID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found with ID %d", MedicationID)
		}
		return nil, err
	}

	// Получаем все выдачи книги
	borrowsByMedication, err := s.Repository.GetBorrowsByMedication(MedicationID)
	if err != nil {
		return nil, err
	}
	if len(borrowsByMedication) == 0 {
		return nil, fmt.Errorf("no borrows found with book ID %d", MedicationID)
	}

	return borrowsByMedication, nil
}

func (s *Service) GetBorrowsByUserAndMedication(userID, MedicationID int) ([]model.Borrow, error) {
	// Проверяем, существует ли пользователь
	_, err := s.Repository.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with ID %d", userID)
		}
		return nil, err
	}

	// Проверяем, существует ли книга
	_, err = s.Repository.GetMedicationByID(MedicationID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("Medication not found with ID %d", MedicationID)
		}
		return nil, err
	}

	// Получаем все выдачи пользователя и книги
	borrowByUserAndMedication, err := s.Repository.GetBorrowsByUserAndMedication(userID, MedicationID)
	if err != nil {
		return nil, err
	}
	if len(borrowByUserAndMedication) == 0 {
		return nil, fmt.Errorf("borrow not found with user ID %d and Medication ID %d", userID, MedicationID)
	}

	return borrowByUserAndMedication, nil
}
