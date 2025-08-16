package pagination

type ResponseMetaData[T any] struct {
	CurrentPage  int     `json:"current_page"`
	Data         []T     `json:"data"`
	FirstPageUrl string  `json:"first_page_url"`
	NextPageUrl  *string `json:"next_page_url"`
	PerPage      int     `json:"per_page"`
	PrevPageUrl  *string `json:"prev_page_url"`
}

func NewResponseMetaData[T any](currentPage int, data []T, firstPageUrl string, nextPageUrl *string, perPage int, prevPageUrl *string) ResponseMetaData[T] {
	return ResponseMetaData[T]{
		CurrentPage:  currentPage,
		Data:         data,
		FirstPageUrl: firstPageUrl,
		NextPageUrl:  nextPageUrl,
		PerPage:      perPage,
		PrevPageUrl:  prevPageUrl,
	}
}
