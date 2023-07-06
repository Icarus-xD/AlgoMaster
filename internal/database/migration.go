package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"github.com/Icarus-xD/AlgoMaster/internal/model"
)

func runMigration(db *gorm.DB) error {
	migration := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "1",
			Migrate: func (tx *gorm.DB) error {
				err := tx.AutoMigrate(&model.TaskPrice{})
				if err != nil {
					return err
				}

				var count int64
				err = tx.Model(&model.TaskPrice{}).Where("type = ?", "FMNUA").Count(&count).Error
				if err != nil {
					return err
				}

				if count == 0 {
					price := model.TaskPrice{
						Type: "FMNUA",
						Price: 20000,
					}

					err := tx.Create(&price).Error
					if err != nil {
						return err
					}
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Where("type = ?", "FMNUA").Delete(&model.TaskPrice{}).Error
			},
		},
		{
			ID: "2",
			Migrate: func(tx *gorm.DB) error  {
				err := tx.AutoMigrate(&model.User{})
				if err != nil {
					return err
				}

				return nil
			},
		},
	})

	err := migration.Migrate()
	if err != nil {
		return nil
	}

	return nil
}