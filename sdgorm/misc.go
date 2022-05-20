package sdgorm

import (
	"database/sql"
	"github.com/gaorx/stardust4/sderr"
	"gorm.io/gorm"
)

func Transaction[R any](db *gorm.DB, action func(tx *gorm.DB) (R, error), opts ...*sql.TxOptions) (R, error) {
	var r R
	err := db.Transaction(func(tx *gorm.DB) error {
		r0, err := action(tx)
		if err != nil {
			return err
		}
		r = r0
		return nil
	}, opts...)
	if err != nil {
		return r, sderr.Wrap(err, "sdgorm transaction error")
	}
	return r, nil
}

func First[T any](tx *gorm.DB) (T, error) {
	var r T
	dbr := tx.First(&r)
	if dbr.Error != nil {
		return r, dbr.Error
	}
	return r, nil
}

func Last[T any](tx *gorm.DB) (T, error) {
	var r T
	dbr := tx.Last(&r)
	if dbr.Error != nil {
		return r, dbr.Error
	}
	return r, nil
}

func Take[T any](tx *gorm.DB) (T, error) {
	var r T
	dbr := tx.Take(&r)
	if dbr.Error != nil {
		return r, dbr.Error
	}
	return r, nil
}

func Find[T any](tx *gorm.DB) (T, error) {
	var r T
	dbr := tx.Find(&r)
	if dbr.Error != nil {
		return r, dbr.Error
	}
	return r, nil
}
