package helpers

import "gorm.io/gorm"

func Search(query string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == "" {
			return db // Return the DB without additional filtering if no query is provided.
		}
		return db.Where("name ILIKE ?", "%"+query+"%")
	}
}

func SearchTitle(query string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query == "" {
			return db // Return the DB without additional filtering if no query is provided.
		}
		return db.Where("title ILIKE ?", "%"+query+"%")
	}
}

func FilterIsActive(IsActive *bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if IsActive != nil {
			return db.Where("is_active = ?", IsActive) // Return the DB without additional filtering if no query is provided.
		}
		return db
	}
}

func FilterRole(role *int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if role != nil {
			return db.Where("role = ?", role) // Return the DB without additional filtering if no query is provided.
		}
		return db
	}
}

// func FilterLectureID(lectureID *int) func(*gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		if lectureID != nil {
// 			return db.Where("lecture_id = ?", lectureID) // Return the DB without additional filtering if no query is provided.
// 		}
// 		return db
// 	}
// }
