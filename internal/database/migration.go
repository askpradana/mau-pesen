package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	migrationsPath := "internal/db-migration/migrations"

	files, err := ioutil.ReadDir(migrationsPath)

	if err != nil {
		log.Fatal("Failed to read migrations folder: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), "sql") {
			continue
		}

		filePath := filepath.Join(migrationsPath, file.Name())
		sqlBytes, err := ioutil.ReadFile(filePath)

		if err != nil {
			log.Fatal("Failed to read migrations file %s: %v", file.Name(), err)
		}

		sql := string(sqlBytes)
		fmt.Printf("Running muigration: %s\n", file.Name(), err)

		if err := db.Exec(sql).Error; err != nil {
			log.Fatal("Migrations %s failed: %v", file.Name(), err)

		}
	}

	fmt.Println("All migrations applied successfully")
}
