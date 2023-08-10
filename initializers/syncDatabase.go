package initializers

import (
	"example/authzed/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.Account{}, &models.Folder{}, &models.Document{})
	if err != nil {
		panic("Failed sync database")
	}

	/// Validate Assumption
	var folder models.Folder
	DB.First(&folder, "name = ?", "root")
	if folder.ID == 0 {
		// Invalid state -> need to set up
		rootFolder := models.Folder{
			Name:       "root",
			ParentId:   nil,
			SubFolders: []models.Folder{},
			Documents:  []models.Document{},
		}
		result := DB.Create(&rootFolder)
		if result.Error != nil {
			panic("Init state failed")
		}
	}
}
