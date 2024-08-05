package config

import (
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// dbUrl := fmt.Sprintf(`postgres://%s:%s@localhost:%s/crud`, "root", "secret", "5432")

	// db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	// if err != nil {
	// 	return nil, err
	// }

	return &gorm.DB{}, nil
}
