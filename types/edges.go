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

func (p PayeeEdge) GetCursor() string {
	return p.Cursor
}

func (p ProductEdge) GetCursor() string {
	return p.Cursor
}

func (t TransactionEdge) GetCursor() string {
	return t.Cursor
}
