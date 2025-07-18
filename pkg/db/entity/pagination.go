package entity

type PaginationRequest struct {
	Page int `schema:"page"`
	Size int `schema:"size"`
}

func (pagination *PaginationRequest) NormalizePagination(defaultSize int) *PaginationRequest {
	if pagination.Page < 1 {
		pagination.Page = 1
	}

	if pagination.Size < 1 {
		pagination.Size = defaultSize
	}

	return pagination
}

type PaginationResponse struct {
	Page      int  `json:"page"`
	Size      int  `json:"size"`
	TotalPage int  `json:"total_page"`
	TotalData int  `json:"total_data"`
	NextPage  *int `json:"next_page"`
	PrevPage  *int `json:"prev_page"`
}

func (pagination *PaginationRequest) Calculate(totalData int) PaginationResponse {
	resp := PaginationResponse{
		Page:      pagination.Page,
		Size:      pagination.Size,
		TotalData: totalData,
	}
	totalPages := totalData / resp.Size
	if totalData%resp.Size != 0 {
		totalPages += 1
	}
	resp.TotalPage = totalPages

	nextPage := resp.Page + 1
	if nextPage > totalPages {
		nextPage = 0
	}

	prevPage := resp.Page - 1
	if prevPage <= 0 {
		prevPage = 0
	}

	if nextPage > 0 {
		resp.NextPage = &nextPage
	}

	if prevPage > 0 {
		resp.PrevPage = &prevPage
	}

	return resp
}
