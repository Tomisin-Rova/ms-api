package types

// GetCursor return the cursor of a AddressEdge
func (a AddressEdge) GetCursor() string {
	return a.Cursor
}

func (p PersonEdge) GetCursor() string {
	return p.Cursor
}
