// Package ischema contains the types for schema 'information_schema'.
package ischema

// Code generated by xo. DO NOT EDIT.

// PgKeysequal calls the stored procedure 'information_schema._pg_keysequal(smallint[], smallint[]) boolean' on db.
func PgKeysequal(db XODB, v0 []int16, v1 []int16) (bool, error) {
	var err error

	// sql query
	const sqlstr = `SELECT information_schema._pg_keysequal($1, $2)`

	// run query
	var ret bool
	XOLog(sqlstr, v0, v1)
	err = db.QueryRow(sqlstr, v0, v1).Scan(&ret)
	if err != nil {
		return false, err
	}

	return ret, nil
}
