package mysql

import (
	"context"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Multi64 []int64
type Multi []int32

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

func (Multi64) GormDataType() string {
	return "multi64"
}

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
		sb.WriteString(strconv.FormatInt(int64(v), 10))
	}
	sb.WriteByte(')')
	return clause.Expr{SQL: sb.String(), Vars: nil}
}

func (Multi) GormDataType() string {
	return "multi"
}
