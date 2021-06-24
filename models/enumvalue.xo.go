package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// EnumValue is a enum value.
type EnumValue struct {
	EnumValue  string `json:"enum_value"`  // enum_value
	ConstValue int    `json:"const_value"` // const_value
}

// PostgresEnumValues runs a custom query, returning results as EnumValue.
func PostgresEnumValues(ctx context.Context, db DB, schema, enum string) ([]*EnumValue, error) {
	// query
	const sqlstr = `SELECT ` +
		`e.enumlabel, ` + // ::varchar AS enum_value
		`e.enumsortorder ` + // ::integer AS const_value
		`FROM pg_type t ` +
		`JOIN ONLY pg_namespace n ON n.oid = t.typnamespace ` +
		`LEFT JOIN pg_enum e ON t.oid = e.enumtypid ` +
		`WHERE n.nspname = $1 ` +
		`AND t.typname = $2`
	// run
	logf(sqlstr, schema, enum)
	rows, err := db.QueryContext(ctx, sqlstr, schema, enum)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*EnumValue
	for rows.Next() {
		var ev EnumValue
		// scan
		if err := rows.Scan(&ev.EnumValue, &ev.ConstValue); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ev)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
