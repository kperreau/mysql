package mysql

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Multi64 []int64

func (m Multi64) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if len(m) == 0 {
		return clause.Expr{SQL: "()", Vars: nil}
	}
	var sb strings.Builder
	sb.WriteByte('(')
	for i, v := range m {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.FormatInt(v, 10))
	}
	sb.WriteByte(')')
	return clause.Expr{SQL: sb.String(), Vars: nil}
}

func (m Multi64) Value() (driver.Value, error) {
	if len(m) == 0 {
		return "()", nil
	}
	var sb strings.Builder
	sb.WriteByte('(')
	for i, v := range m {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.FormatInt(v, 10))
	}
	sb.WriteByte(')')
	return sb.String(), nil
}

func (Multi64) GormDataType() string {
	return "multi64"
}

func (m *Multi64) Scan(value interface{}) error {
	if value == nil {
		*m = Multi64{}
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("multi64: cannot convert non-byte value")
	}

	str := string(b)
	if str == "" {
		*m = Multi64{}
		return nil
	}

	parts := strings.Split(str, ",")
	out := make(Multi64, 0, len(parts))
	for _, p := range parts {
		n, err := strconv.ParseInt(strings.TrimSpace(p), 10, 64)
		if err != nil {
			return err
		}
		out = append(out, n)
	}
	*m = out
	return nil
}

type Multi []int32

func (m Multi) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if len(m) == 0 {
		return clause.Expr{SQL: "()", Vars: nil}
	}
	var sb strings.Builder
	sb.WriteByte('(')
	for i, v := range m {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(int(v)))
	}
	sb.WriteByte(')')
	return clause.Expr{SQL: sb.String(), Vars: nil}
}

func (m Multi) Value() (driver.Value, error) {
	if len(m) == 0 {
		return "()", nil
	}
	var sb strings.Builder
	sb.WriteByte('(')
	for i, v := range m {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(int(v)))
	}
	sb.WriteByte(')')
	return sb.String(), nil
}

func (Multi) GormDataType() string {
	return "multi"
}

func (m *Multi) Scan(value interface{}) error {
	if value == nil {
		*m = Multi{}
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("multi: cannot convert non-byte value")
	}

	str := string(b)
	if str == "" {
		*m = Multi{}
		return nil
	}

	parts := strings.Split(str, ",")
	out := make(Multi, 0, len(parts))
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return err
		}
		out = append(out, int32(n))
	}
	*m = out
	return nil
}

type DocID uint64

func (DocID) GormDataType() string {
	return "docId"
}

func (id DocID) Value() (driver.Value, error) {
	return uint64(id), nil
}

func (id *DocID) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		if v < 0 {
			// Manticore encode parfois des uint64 en int64 signé
			*id = DocID(uint64(v))
		} else {
			*id = DocID(v)
		}
		return nil
	case uint64:
		*id = DocID(v)
		return nil
	default:
		return fmt.Errorf("DocID: unsupported Scan type: %T", value)
	}
}
