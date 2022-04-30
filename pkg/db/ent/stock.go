// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/stock-manager/pkg/db/ent/stock"
	"github.com/google/uuid"
)

// Stock is the model entity for the Stock schema.
type Stock struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// Total holds the value of the "total" field.
	Total uint32 `json:"total,omitempty"`
	// Locked holds the value of the "locked" field.
	Locked uint32 `json:"locked,omitempty"`
	// InService holds the value of the "in_service" field.
	InService uint32 `json:"in_service,omitempty"`
	// Sold holds the value of the "sold" field.
	Sold uint32 `json:"sold,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Stock) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case stock.FieldCreatedAt, stock.FieldUpdatedAt, stock.FieldDeletedAt, stock.FieldTotal, stock.FieldLocked, stock.FieldInService, stock.FieldSold:
			values[i] = new(sql.NullInt64)
		case stock.FieldID, stock.FieldGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Stock", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Stock fields.
func (s *Stock) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case stock.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case stock.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = uint32(value.Int64)
			}
		case stock.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = uint32(value.Int64)
			}
		case stock.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = uint32(value.Int64)
			}
		case stock.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				s.GoodID = *value
			}
		case stock.FieldTotal:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total", values[i])
			} else if value.Valid {
				s.Total = uint32(value.Int64)
			}
		case stock.FieldLocked:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				s.Locked = uint32(value.Int64)
			}
		case stock.FieldInService:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field in_service", values[i])
			} else if value.Valid {
				s.InService = uint32(value.Int64)
			}
		case stock.FieldSold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sold", values[i])
			} else if value.Valid {
				s.Sold = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Stock.
// Note that you need to call Stock.Unwrap() before calling this method if this Stock
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Stock) Update() *StockUpdateOne {
	return (&StockClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Stock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Stock) Unwrap() *Stock {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Stock is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Stock) String() string {
	var builder strings.Builder
	builder.WriteString("Stock(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", s.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", s.UpdatedAt))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", s.DeletedAt))
	builder.WriteString(", good_id=")
	builder.WriteString(fmt.Sprintf("%v", s.GoodID))
	builder.WriteString(", total=")
	builder.WriteString(fmt.Sprintf("%v", s.Total))
	builder.WriteString(", locked=")
	builder.WriteString(fmt.Sprintf("%v", s.Locked))
	builder.WriteString(", in_service=")
	builder.WriteString(fmt.Sprintf("%v", s.InService))
	builder.WriteString(", sold=")
	builder.WriteString(fmt.Sprintf("%v", s.Sold))
	builder.WriteByte(')')
	return builder.String()
}

// Stocks is a parsable slice of Stock.
type Stocks []*Stock

func (s Stocks) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
