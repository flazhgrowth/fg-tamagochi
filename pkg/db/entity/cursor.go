package entity

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
)

type (
	CursorPaginationRequest struct {
		Size   int    `query:"size"`
		Cursor string `query:"cursor"`

		Field      string      `query:"-"`
		LastSeenID LastSeenVal `query:"-"`
	}
	CursorPaginationResponse struct {
		TotalData int     `json:"total_data"`
		Cursor    *string `json:"cursor"`
	}
)

// DecodeCursor decodes the base64-encoded cursor and populates Field and
// LastSeenID for the caller to build the next query from. An empty or
// malformed cursor is treated as the first page rather than an error.
func (pager *CursorPaginationRequest) DecodeCursor() *CursorPaginationRequest {
	if pager.Cursor == "" {
		return pager
	}

	dataByte, err := base64.StdEncoding.DecodeString(pager.Cursor)
	if err != nil {
		pager.Cursor = ""
		return pager
	}

	var decoded cursorPagination
	if err := json.Unmarshal(dataByte, &decoded); err != nil {
		pager.Cursor = ""
		return pager
	}

	pager.Field = decoded.Field
	pager.LastSeenID = decoded.LastSeenID
	return pager
}

func (page cursorPagination) encodeCursor() (string, error) {
	dataByte, err := json.Marshal(page)
	if err != nil {
		return "", err
	}

	encodedPage := base64.StdEncoding.EncodeToString(dataByte)

	return encodedPage, nil
}

// Calculate fills in the response's total count and, when hasNext is true,
// an encoded cursor pointing at lastSeenID. Pass hasNext=false (e.g. when
// fewer rows than the requested size were returned) to leave Cursor nil and
// signal the client there are no more pages.
func (pager *CursorPaginationResponse) Calculate(field string, lastSeenID LastSeenVal, totalData int, hasNext bool) error {
	pager.TotalData = totalData

	if !hasNext {
		pager.Cursor = nil
		return nil
	}

	cursorPage := cursorPagination{Field: field, LastSeenID: lastSeenID}
	cursor, err := cursorPage.encodeCursor()
	if err != nil {
		return err
	}
	pager.Cursor = &cursor

	return nil
}

type (
	LastSeenVal      string
	cursorPagination struct {
		Field      string      `json:"fld"`
		LastSeenID LastSeenVal `json:"lsid"`
	}
)

func (lsv LastSeenVal) Uint64() uint64 {
	val, err := strconv.ParseUint(string(lsv), 10, 64)
	if err != nil {
		return 0
	}

	return val
}
