package commons

const (
	DEFAULTSIZE int = 10
	DEFAULTPAGE int = 1
)

type Paginate struct {
	Page      int `json:"page"`
	Size      int `json:"size"`
	TotalPage int `json:"totalPage"`
}

func ToPaginate(page int, size int, totalPage int) Paginate {
	return Paginate{
		Page:      page,
		Size:      size,
		TotalPage: totalPage,
	}
}
