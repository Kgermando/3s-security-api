package services

import ( 
	"github.com/kgermando/backend/database" 
)

// CreateResource creates a new resource
func CreateResource(resource interface{}) error {
	db := database.DB
	if err := db.Create(resource).Error; err != nil {
		return err
	}
	return nil
}

// GetResources retrieves a paginated list of resources
func GetResources(page int, pageSize int, resources interface{}) error {
	db := database.DB
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(resources).Error; err != nil {
		return err
	}
	return nil
}
 
// GetResource retrieves a single resource by ID
func GetResource(id string, resource interface{}) error {
	db := database.DB
	if err := db.First(resource, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// UpdateResource updates an existing resource by ID
func UpdateResource(id string, resource interface{}) error {
	db := database.DB
	if err := db.Model(resource).Where("id = ?", id).Updates(resource).Error; err != nil {
		return err
	}
	return nil
}

// DeleteResource deletes a resource by ID
func DeleteResource(id string, resource interface{}) error {
	db := database.DB
	if err := db.Where("id = ?", id).Delete(resource).Error; err != nil {
		return err
	}
	return nil
}
