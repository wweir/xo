// Package postgres contains the types for schema 'public'.
package postgres

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// AuthGroupPermission represents a row from public.auth_group_permissions.
type AuthGroupPermission struct {
	ID           int // id
	GroupID      int // group_id
	PermissionID int // permission_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AuthGroupPermission exists in the database.
func (agp *AuthGroupPermission) Exists() bool {
	return agp._exists
}

// Deleted provides information if the AuthGroupPermission has been deleted from the database.
func (agp *AuthGroupPermission) Deleted() bool {
	return agp._deleted
}

// Insert inserts the AuthGroupPermission to the database.
func (agp *AuthGroupPermission) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if agp._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.auth_group_permissions (` +
		`group_id, permission_id` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, agp.GroupID, agp.PermissionID)
	err = db.QueryRow(sqlstr, agp.GroupID, agp.PermissionID).Scan(&agp.ID)
	if err != nil {
		return err
	}

	// set existence
	agp._exists = true

	return nil
}

// Update updates the AuthGroupPermission in the database.
func (agp *AuthGroupPermission) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !agp._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if agp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.auth_group_permissions SET (` +
		`group_id, permission_id` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE id = $3`

	// run query
	XOLog(sqlstr, agp.GroupID, agp.PermissionID, agp.ID)
	_, err = db.Exec(sqlstr, agp.GroupID, agp.PermissionID, agp.ID)
	return err
}

// Save saves the AuthGroupPermission to the database.
func (agp *AuthGroupPermission) Save(db XODB) error {
	if agp.Exists() {
		return agp.Update(db)
	}

	return agp.Insert(db)
}

// Upsert performs an upsert for AuthGroupPermission.
//
// NOTE: PostgreSQL 9.5+ only
func (agp *AuthGroupPermission) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if agp._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.auth_group_permissions (` +
		`id, group_id, permission_id` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, group_id, permission_id` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.group_id, EXCLUDED.permission_id` +
		`)`

	// run query
	XOLog(sqlstr, agp.ID, agp.GroupID, agp.PermissionID)
	_, err = db.Exec(sqlstr, agp.ID, agp.GroupID, agp.PermissionID)
	if err != nil {
		return err
	}

	// set existence
	agp._exists = true

	return nil
}

// Delete deletes the AuthGroupPermission from the database.
func (agp *AuthGroupPermission) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !agp._exists {
		return nil
	}

	// if deleted, bail
	if agp._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.auth_group_permissions WHERE id = $1`

	// run query
	XOLog(sqlstr, agp.ID)
	_, err = db.Exec(sqlstr, agp.ID)
	if err != nil {
		return err
	}

	// set deleted
	agp._deleted = true

	return nil
}

// AuthPermission returns the AuthPermission associated with the AuthGroupPermission's PermissionID (permission_id).
//
// Generated from foreign key 'auth_group_permiss_permission_id_84c5c92e_fk_auth_permission_id'.
func (agp *AuthGroupPermission) AuthPermission(db XODB) (*AuthPermission, error) {
	return AuthPermissionByID(db, agp.PermissionID)
}

// AuthGroup returns the AuthGroup associated with the AuthGroupPermission's GroupID (group_id).
//
// Generated from foreign key 'auth_group_permissions_group_id_b120cbf9_fk_auth_group_id'.
func (agp *AuthGroupPermission) AuthGroup(db XODB) (*AuthGroup, error) {
	return AuthGroupByID(db, agp.GroupID)
}

// AuthGroupPermissionsByGroupID retrieves a row from 'public.auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_0e939a4f'.
func AuthGroupPermissionsByGroupID(db XODB, groupID int) ([]*AuthGroupPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM public.auth_group_permissions ` +
		`WHERE group_id = $1`

	// run query
	XOLog(sqlstr, groupID)
	q, err := db.Query(sqlstr, groupID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthGroupPermission{}
	for q.Next() {
		agp := AuthGroupPermission{
			_exists: true,
		}

		// scan
		err = q.Scan(&agp.ID, &agp.GroupID, &agp.PermissionID)
		if err != nil {
			return nil, err
		}

		res = append(res, &agp)
	}

	return res, nil
}

// AuthGroupPermissionsByPermissionID retrieves a row from 'public.auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_8373b171'.
func AuthGroupPermissionsByPermissionID(db XODB, permissionID int) ([]*AuthGroupPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM public.auth_group_permissions ` +
		`WHERE permission_id = $1`

	// run query
	XOLog(sqlstr, permissionID)
	q, err := db.Query(sqlstr, permissionID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthGroupPermission{}
	for q.Next() {
		agp := AuthGroupPermission{
			_exists: true,
		}

		// scan
		err = q.Scan(&agp.ID, &agp.GroupID, &agp.PermissionID)
		if err != nil {
			return nil, err
		}

		res = append(res, &agp)
	}

	return res, nil
}

// AuthGroupPermissionByGroupIDPermissionID retrieves a row from 'public.auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_group_id_0cd325b0_uniq'.
func AuthGroupPermissionByGroupIDPermissionID(db XODB, groupID int, permissionID int) (*AuthGroupPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM public.auth_group_permissions ` +
		`WHERE group_id = $1 AND permission_id = $2`

	// run query
	XOLog(sqlstr, groupID, permissionID)
	agp := AuthGroupPermission{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, groupID, permissionID).Scan(&agp.ID, &agp.GroupID, &agp.PermissionID)
	if err != nil {
		return nil, err
	}

	return &agp, nil
}

// AuthGroupPermissionByID retrieves a row from 'public.auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_pkey'.
func AuthGroupPermissionByID(db XODB, id int) (*AuthGroupPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM public.auth_group_permissions ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	agp := AuthGroupPermission{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&agp.ID, &agp.GroupID, &agp.PermissionID)
	if err != nil {
		return nil, err
	}

	return &agp, nil
}
