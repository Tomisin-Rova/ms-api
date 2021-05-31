//go:generate genny -in=connection_template.go -out=gen_address_lookup.go gen "Name=AddressLookup NodeType=*types.Address EdgeType=types.AddressEdge ConnectionType=types.AddressConnection"
//go:generate genny -in=connection_template.go -out=gen_people_lookup.go gen "Name=PeopleLookup NodeType=*types.Person EdgeType=types.PersonEdge ConnectionType=types.PersonConnection"
//go:generate genny -in=connection_template.go -out=gen_cdd_lookup.go gen "Name=CddLookup NodeType=*types.Cdd EdgeType=types.CDDEdge ConnectionType=types.CDDConnection"
//go:generate genny -in=connection_template.go -out=gen_account_lookup.go gen "Name=AccountConnection NodeType=*types.Account EdgeType=types.AccountEdge ConnectionType=types.AccountConnection"
//go:generate genny -in=connection_template.go -out=gen_tag_lookup.go gen "Name=TagLookup NodeType=*types.Tag EdgeType=types.TagEdge ConnectionType=types.TagConnection"
//go:generate genny -in=connection_template.go -out=gen_payee_lookup.go gen "Name=PayeeConnection NodeType=*types.Payee EdgeType=types.PayeeEdge ConnectionType=types.PayeeConnection"
//go:generate genny -in=connection_template.go -out=gen_product_lookup.go gen "Name=ProductConnection NodeType=*types.Product EdgeType=types.ProductEdge ConnectionType=types.ProductConnection"
//go:generate genny -in=connection_template.go -out=gen_transaction_lookup.go gen "Name=TransactionConnection NodeType=*types.Transaction EdgeType=types.TransactionEdge ConnectionType=types.TransactionConnection"
//go:generate genny -in=connection_template.go -out=gen_tags_lookup.go gen "Name=TagConnection NodeType=*types.Tag EdgeType=types.TagEdge ConnectionType=types.TagConnection"

// Package connections implement a generic GraphQL relay connection
package connections

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

const cursorPrefix = "cursor:"

// Edge define the contract for an edge in a relay connection
type Edge interface {
	GetCursor() string
}

// Pagination stores the pagination details based on its cursor-based counterpart
type Pagination struct {
	After  string
	First  int64
	Before string
	Last   int64
}

// OffsetToCursor create the cursor string from an offset
func OffsetToCursor(offset int) string {
	str := fmt.Sprintf("%v%v", cursorPrefix, offset)
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// CursorToOffset re-derives the offset from the cursor string.
func CursorToOffset(cursor string) (int, error) {
	str := ""
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err == nil {
		str = string(b)
	}
	str = strings.Replace(str, cursorPrefix, "", -1)
	offset, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("invalid cursor")
	}
	return offset, nil
}

// IdToCursor create the cursor string from an id
func IdToCursor(id string) string {
	return base64.StdEncoding.EncodeToString([]byte(id))
}

// CursorToId re-derives the id from the cursor string.
func CursorToId(cursor string) (string, error) {
	str := ""
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", fmt.Errorf("invalid cursor")
	}
	str = string(b)
	return str, nil
}

// Paginate resolves values for pagination
func Paginate(first *int64, after *string, last *int64, before *string) (Pagination, error) {
	var pagAfter, pagBefore string
	var pagFirst, pagLast int64
	pagFirst = 0
	pagLast = 0
	var err error
	if after != nil {
		pagAfter, err = CursorToId(*after)
		if err != nil {
			return Pagination{}, err
		}
	}
	if before != nil {
		pagBefore, err = CursorToId(*before)
		if err != nil {
			return Pagination{}, err
		}
	}
	if first != nil {
		pagFirst = *first
	}
	if last != nil {
		pagLast = *last
	}
	if first == nil && after == nil && last == nil && before == nil {
		pagFirst = 100
	}

	return Pagination{First: pagFirst, After: pagAfter, Last: pagLast, Before: pagBefore}, err
}
