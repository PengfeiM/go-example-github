package unittestexpl

import (
	"fmt"

	"gorm.io/gorm"
)

type testModel struct {
	name string
	val  int
}

func UpsertTestModelToDb(tm *testModel, tx *gorm.DB) error {
	if tm == nil {
		return fmt.Errorf("testModel is nil")
	}
	fmt.Printf("got valid testModel:[%v]", tm)

	if tx == nil {
		return fmt.Errorf("db tx is nil")
	}
	fmt.Printf("got vaild db:[%v]", tx)

	condition := &testModel{
		name: tm.name,
	}
	result := tx.Model(condition).Where(condition).Assign(tm).FirstOrCreate(condition)
	if result.Error != nil {
		fmt.Printf("upsert test model to db failed: [%s]", result.Error.Error())
		return result.Error
	}
	return nil
}
