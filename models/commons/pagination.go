package commons

import "math"

const (
	DEFAULTSIZE int = 10
	DEFAULTPAGE int = 1
)

type Paginate struct {
	Page      int `json:"page"`
	Size      int `json:"size"`
	TotalPage int `json:"totalPage"`
}

func ToPaginate(page int, size int, totalData int) Paginate {
	return Paginate{
		Page:      page,
		Size:      size,
		TotalPage: int(math.Ceil(float64(totalData) / float64(size))),
	}
}
