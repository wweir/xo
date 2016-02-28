// Package sqlite3 contains the types for schema ''.
package sqlite3

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// AuthUserGroup represents a row from 'auth_user_groups'.
type AuthUserGroup struct {
	ID      int // id
	UserID  int // user_id
	GroupID int // group_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AuthUserGroup exists in the database.
func (aug *AuthUserGroup) Exists() bool {
	return aug._exists
}

// Deleted provides information if the AuthUserGroup has been deleted from the database.
func (aug *AuthUserGroup) Deleted() bool {
	return aug._deleted
}

// Insert inserts the AuthUserGroup to the database.
func (aug *AuthUserGroup) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if aug._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO auth_user_groups (` +
		`user_id, group_id` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, aug.UserID, aug.GroupID)
	res, err := db.Exec(sqlstr, aug.UserID, aug.GroupID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	aug.ID = int(id)
	aug._exists = true

	return nil
}

// Update updates the AuthUserGroup in the database.
func (aug *AuthUserGroup) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !aug._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if aug._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE auth_user_groups SET ` +
		`user_id = ?, group_id = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, aug.UserID, aug.GroupID, aug.ID)
	_, err = db.Exec(sqlstr, aug.UserID, aug.GroupID, aug.ID)
	return err
}

// Save saves the AuthUserGroup to the database.
func (aug *AuthUserGroup) Save(db XODB) error {
	if aug.Exists() {
		return aug.Update(db)
	}

	return aug.Insert(db)
}

// Delete deletes the AuthUserGroup from the database.
func (aug *AuthUserGroup) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !aug._exists {
		return nil
	}

	// if deleted, bail
	if aug._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM auth_user_groups WHERE id = ?`

	// run query
	XOLog(sqlstr, aug.ID)
	_, err = db.Exec(sqlstr, aug.ID)
	if err != nil {
		return err
	}

	// set deleted
	aug._deleted = true

	return nil
}

// AuthGroup returns the AuthGroup associated with the AuthUserGroup's GroupID (group_id).
//
// Generated from foreign key 'auth_user_groups_group_id_fkey'.
func (aug *AuthUserGroup) AuthGroup(db XODB) (*AuthGroup, error) {
	return AuthGroupByID(db, aug.GroupID)
}

// AuthUser returns the AuthUser associated with the AuthUserGroup's UserID (user_id).
//
// Generated from foreign key 'auth_user_groups_user_id_fkey'.
func (aug *AuthUserGroup) AuthUser(db XODB) (*AuthUser, error) {
	return AuthUserByID(db, aug.UserID)
}

// AuthUserGroupsByGroupID retrieves a row from 'auth_user_groups' as a AuthUserGroup.
//
// Generated from index 'auth_user_groups_0e939a4f'.
func AuthUserGroupsByGroupID(db XODB, groupID int) ([]*AuthUserGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, group_id ` +
		`FROM auth_user_groups ` +
		`WHERE group_id = ?`

	// run query
	XOLog(sqlstr, groupID)
	q, err := db.Query(sqlstr, groupID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthUserGroup{}
	for q.Next() {
		aug := AuthUserGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&aug.ID, &aug.UserID, &aug.GroupID)
		if err != nil {
			return nil, err
		}

		res = append(res, &aug)
	}

	return res, nil
}

// AuthUserGroupsByUserID retrieves a row from 'auth_user_groups' as a AuthUserGroup.
//
// Generated from index 'auth_user_groups_e8701ad4'.
func AuthUserGroupsByUserID(db XODB, userID int) ([]*AuthUserGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, group_id ` +
		`FROM auth_user_groups ` +
		`WHERE user_id = ?`

	// run query
	XOLog(sqlstr, userID)
	q, err := db.Query(sqlstr, userID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthUserGroup{}
	for q.Next() {
		aug := AuthUserGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&aug.ID, &aug.UserID, &aug.GroupID)
		if err != nil {
			return nil, err
		}

		res = append(res, &aug)
	}

	return res, nil
}

// AuthUserGroupByID retrieves a row from 'auth_user_groups' as a AuthUserGroup.
//
// Generated from index 'auth_user_groups_id_pkey'.
func AuthUserGroupByID(db XODB, id int) (*AuthUserGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, group_id ` +
		`FROM auth_user_groups ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	aug := AuthUserGroup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&aug.ID, &aug.UserID, &aug.GroupID)
	if err != nil {
		return nil, err
	}

	return &aug, nil
}

// AuthUserGroupByUserIDGroupID retrieves a row from 'auth_user_groups' as a AuthUserGroup.
//
// Generated from index 'auth_user_groups_user_id_94350c0c_uniq'.
func AuthUserGroupByUserIDGroupID(db XODB, userID int, groupID int) (*AuthUserGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, group_id ` +
		`FROM auth_user_groups ` +
		`WHERE user_id = ? AND group_id = ?`

	// run query
	XOLog(sqlstr, userID, groupID)
	aug := AuthUserGroup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, userID, groupID).Scan(&aug.ID, &aug.UserID, &aug.GroupID)
	if err != nil {
		return nil, err
	}

	return &aug, nil
}
