// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StocksColumns holds the columns for the "stocks" table.
	StocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "total", Type: field.TypeUint32},
		{Name: "locked", Type: field.TypeUint32},
		{Name: "in_service", Type: field.TypeUint32},
		{Name: "sold", Type: field.TypeUint32},
	}
	// StocksTable holds the schema information for the "stocks" table.
	StocksTable = &schema.Table{
		Name:       "stocks",
		Columns:    StocksColumns,
		PrimaryKey: []*schema.Column{StocksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "stock_good_id",
				Unique:  true,
				Columns: []*schema.Column{StocksColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StocksTable,
	}
)

func init() {
}
