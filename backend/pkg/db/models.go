// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type BackgroundType string

const (
	BackgroundTypeBTNONE    BackgroundType = "BT_NONE"
	BackgroundTypeBTOUTSIDE BackgroundType = "BT_OUTSIDE"
	BackgroundTypeBTFULL    BackgroundType = "BT_FULL"
)

func (e *BackgroundType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BackgroundType(s)
	case string:
		*e = BackgroundType(s)
	default:
		return fmt.Errorf("unsupported scan type for BackgroundType: %T", src)
	}
	return nil
}

type NullBackgroundType struct {
	BackgroundType BackgroundType
	Valid          bool // Valid is true if BackgroundType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBackgroundType) Scan(value interface{}) error {
	if value == nil {
		ns.BackgroundType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BackgroundType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBackgroundType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BackgroundType), nil
}

type Category string

const (
	CategoryCBASKETBALLCOMPETITION Category = "C_BASKETBALL_COMPETITION"
	CategoryCBASKETBALLCOMPETITOR  Category = "C_BASKETBALL_COMPETITOR"
	CategoryCFOOTBALLCOMPETITION   Category = "C_FOOTBALL_COMPETITION"
	CategoryCFOOTBALLCOMPETITOR    Category = "C_FOOTBALL_COMPETITOR"
)

func (e *Category) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Category(s)
	case string:
		*e = Category(s)
	default:
		return fmt.Errorf("unsupported scan type for Category: %T", src)
	}
	return nil
}

type NullCategory struct {
	Category Category
	Valid    bool // Valid is true if Category is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCategory) Scan(value interface{}) error {
	if value == nil {
		ns.Category, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Category.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCategory) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Category), nil
}

type Image struct {
	ID             int64
	Category       Category
	BackgroundType NullBackgroundType
	LabelerID      pgtype.Int8
	Name           string
	DisplayName    pgtype.Text
	Url1           string
	Url2           string
	Url3           string
	UrlSelected    pgtype.Int2
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
	Region         pgtype.Text
	DisplayOrder   pgtype.Int4
	Url1Title      string
	Url2Title      string
	Url3Title      string
}

type User struct {
	ID          int64
	Username    string
	Password    string
	DisplayName string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	Email       string
}
