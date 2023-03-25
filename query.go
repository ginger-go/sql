package sql

import "gorm.io/gorm"

// FindOne finds one record that matches the given clause.
func FindOne[T any](tx *gorm.DB, clause *Clause) (*T, error) {
	if clause != nil {
		tx = clause.Consume(tx)
	}
	output := new(T)
	err := tx.First(output).Error
	return output, err
}

// FindAll finds all records that match the given clause.
func FindAll[T any](tx *gorm.DB, clause *Clause) ([]T, error) {
	if clause != nil {
		tx = clause.Consume(tx)
	}
	output := make([]T, 0)
	err := tx.Find(&output).Error
	return output, err
}

// Count counts the number of records that match the given clause.
func Count(tx *gorm.DB, clause *Clause) (int64, error) {
	if clause != nil {
		tx = clause.Consume(tx)
	}
	var count int64
	err := tx.Count(&count).Error
	return count, err
}

// FindAllComplex finds all records that match the given clause and applies the given sort and pagination.
func FindAllComplex[T any](tx *gorm.DB, clause *Clause, sort *Sort, pagination *Pagination) ([]T, *Pagination, error) {
	if clause != nil {
		tx = clause.Consume(tx)
	}
	if sort != nil {
		tx = sort.Consume(tx)
	}
	if pagination != nil {
		tx = pagination.Consume(tx)
	}
	output := make([]T, 0)
	err := tx.Find(&output).Error
	if err != nil {
		return nil, nil, err
	}
	if pagination != nil {
		pagination.Total, err = Count(tx, clause)
		if err != nil {
			return nil, nil, err
		}
	}
	return output, pagination, err
}