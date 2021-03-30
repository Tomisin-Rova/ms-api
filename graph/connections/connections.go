//go:generate genny -in=connection_template.go -out=gen_address_lookup.go gen "Name=AddressLookup NodeType=*types.Address EdgeType=types.AddressEdge ConnectionType=types.AddressConnection"
//go:generate genny -in=connection_template.go -out=gen_people_lookup.go gen "Name=PeopleLookup NodeType=*types.Person EdgeType=types.PersonEdge ConnectionType=types.PersonConnection"
//go:generate genny -in=connection_template.go -out=gen_cdd_lookup.go gen "Name=CddLookup NodeType=*types.Cdd EdgeType=types.CDDEdge ConnectionType=types.CDDConnection"
//go:generate genny -in=connection_template.go -out=gen_account_lookup.go gen "Name=AccountConnection NodeType=*types.Account EdgeType=types.AccountEdge ConnectionType=types.AccountConnection"
//go:generate genny -in=connection_template.go -out=gen_tag_lookup.go gen "Name=TagLookup NodeType=*types.Tag EdgeType=types.TagEdge ConnectionType=types.TagConnection"

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
		return 0, fmt.Errorf("Invalid cursor")
	}
	return offset, nil
}
