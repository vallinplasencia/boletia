package resp

// Currency ...
type Currency struct {
	ID    int     `json:"id"`
	Code  string  `json:"code"`
	Value float64 `json:"value"`
	// fecha unix
	LastUpdateAt int64 `json:"last_update_at"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
