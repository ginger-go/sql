package sql

import "gorm.io/gorm"

type Sort struct {
	By  string `form:"by"`
	Asc bool   `form:"asc"`
}

func (s *Sort) Consume(tx *gorm.DB) *gorm.DB {
	if s.By != "" {
		if s.Asc {
			tx = tx.Order(s.By)
		} else {
			tx = tx.Order(s.By + " DESC")
		}
	}
	return tx
}
