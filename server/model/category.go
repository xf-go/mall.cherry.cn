package model

// Category - table category
type Category struct {
	ID   uint
	Name string
}

func GetCategories() ([]*Category, error) {
	var categories []*Category
	err := DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
