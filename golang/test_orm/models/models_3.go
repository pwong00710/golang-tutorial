package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	// auto-populate columns: id, created_at, updated_at, deleted_at
	gorm.Model
	// Or alternatively write:
	//Model gorm.Model `gorm:"embedded"`

	// If you don't want to include that many columns, simply use:
	//ID uint
	// Which gorm will still set it as primary_key

	// Set column type manually
	Username string `sql:"type:VARCHAR(255)"`

	// Set default value
	LastName string `sql:"DEFAULT:'Smith'"`

	// Ignored attribute will be treated as attr instead of column
	IgnoredField bool `sql:"-"`

	// Custom primary ket
	UserID int `gorm:"primary_key"`

	// Custom column name instead of default snake_case format
	FirstName string `gorm:"column:FirstName"`

	// AUTO_INCREMENT can only be set on key field
	Count int `gorm:"AUTO_INCREMENT"`

	// Not Null & Unique field
	//Username string `sql:"not null;unique"`

	Salary float64

	Birthday time.Time

	CreditCard CreditCard // One-To-One relationship (has one - use CreditCard's UserID as foreign key)

	Emails []Email // One-To-Many relationship (has many - use Email's UserID as foreign key)

	BillingAddress   Address // One-To-One relationship (belongs to - use BillingAddressID as foreign key)
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One relationship (belongs to - use ShippingAddressID as foreign key)
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`                         // Ignore this field
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}

func (u *User) TableName() string {
	// custom table name, this is default
	return "users"
}

/*
func (u *User) BeforeSave() (err error) {
	if u.Role != "admin" {
		err = errors.New("Permission denied.")
	}
	return
}
*/

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // Foreign key (belongs to), tag `index` will create index for this column
	Email      string `gorm:"type:varchar(100);unique_index"` // `type` set sql type, `unique_index` will create unique index for this column
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // Set field as not nullable and unique
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // Create index with name, and will create combined index if find other fields defined same name
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}
