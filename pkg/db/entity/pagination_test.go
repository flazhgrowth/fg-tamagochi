package entity

import (
	"testing"

	"github.com/test-go/testify/assert"
)

func TestPaginationRequest_Calculate(t *testing.T) {
	type fields struct {
		Page int
		Size int
	}
	type args struct {
		totalData int
	}
	getInt := func(x int) *int {
		return &x
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PaginationResponse
	}{
		{
			name: "success_paginate_1",
			fields: fields{
				Page: 1,
				Size: 10,
			},
			args: args{totalData: 10},
			want: PaginationResponse{
				Page:      1,
				Size:      10,
				TotalPage: 1,
				TotalData: 10,
				NextPage:  nil,
				PrevPage:  nil,
			},
		},
		{
			name: "success_paginate_2",
			fields: fields{
				Page: 2,
				Size: 10,
			},
			args: args{totalData: 10},
			want: PaginationResponse{
				Page:      2,
				Size:      10,
				TotalPage: 1,
				TotalData: 10,
				NextPage:  nil,
				PrevPage:  getInt(1),
			},
		},
		{
			name: "success_paginate_3",
			fields: fields{
				Page: 0,
				Size: 10,
			},
			args: args{totalData: 10},
			want: PaginationResponse{
				Page:      1,
				Size:      10,
				TotalPage: 1,
				TotalData: 10,
				NextPage:  nil,
				PrevPage:  nil,
			},
		},
		{
			name: "success_paginate_4",
			fields: fields{
				Page: 1,
				Size: 10,
			},
			args: args{totalData: 21},
			want: PaginationResponse{
				Page:      1,
				Size:      10,
				TotalPage: 3,
				TotalData: 21,
				NextPage:  getInt(2),
				PrevPage:  nil,
			},
		},
		{
			name: "success_paginate_5",
			fields: fields{
				Page: 2,
				Size: 10,
			},
			args: args{totalData: 21},
			want: PaginationResponse{
				Page:      2,
				Size:      10,
				TotalPage: 3,
				TotalData: 21,
				NextPage:  getInt(3),
				PrevPage:  getInt(1),
			},
		},
		{
			name: "success_paginate_6",
			fields: fields{
				Page: 3,
				Size: 10,
			},
			args: args{totalData: 21},
			want: PaginationResponse{
				Page:      3,
				Size:      10,
				TotalPage: 3,
				TotalData: 21,
				NextPage:  nil,
				PrevPage:  getInt(2),
			},
		},
		{
			name: "success_paginate_6",
			fields: fields{
				Page: 4,
				Size: 10,
			},
			args: args{totalData: 21},
			want: PaginationResponse{
				Page:      4,
				Size:      10,
				TotalPage: 3,
				TotalData: 21,
				NextPage:  nil,
				PrevPage:  getInt(3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pagination := &PaginationRequest{
				Page: tt.fields.Page,
				Size: tt.fields.Size,
			}
			pagination.NormalizePagination(20)
			got := pagination.Calculate(tt.args.totalData)
			assert.Equal(t, tt.want, got)
		})
	}
}
