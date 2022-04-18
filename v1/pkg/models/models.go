package models

import "time"

// Currency ...
type Currency struct {
	ID    int
	Code  string
	Value float64
	// fecha unix
	LastUpdateAt int64

	CreatedAt int64
	UpdatedAt int64
}

// Request ...
type Request struct {
	Date     time.Time
	Duration string
	Status   StatusType

	CreatedAt int64
	UpdatedAt int64
}

// StatusType estado de una peticion
type StatusType string

const (
	// StatusOK la peticion se realizo correctamente
	StatusOK StatusType = "ok"
	// FailedOK existio error en la peticion
	FailedOK StatusType = "failed"
)
