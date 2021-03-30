package types

// GetCursor return the cursor of a AddressEdge
func (a AddressEdge) GetCursor() string {
	return a.Cursor
}

func (p PersonEdge) GetCursor() string {
	return p.Cursor
}

func (c CDDEdge) GetCursor() string {
	return c.Cursor
}

func (ac AccountEdge) GetCursor() string {
	return ac.Cursor
}

func (ac TagEdge) GetCursor() string {
	return ac.Cursor
}
