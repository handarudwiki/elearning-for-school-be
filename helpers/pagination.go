package helpers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

func GetPaginationParams(ctx *fiber.Ctx, pageDefault int, sizeDefault int) (page int, size int) {
	size, err := strconv.Atoi(ctx.Query("size"))
	if err != nil || size == 0 {
		size = sizeDefault
	}

	page, err = strconv.Atoi(ctx.Query("page"))
	if err != nil || page == 0 {
		page = pageDefault
	}

	return page, size
}
