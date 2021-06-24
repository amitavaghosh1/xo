package sqlite3

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// OrderDetail represents a row from 'order_details'.
type OrderDetail struct {
	OrderID   int     `json:"order_id"`   // order_id
	ProductID int     `json:"product_id"` // product_id
	UnitPrice float64 `json:"unit_price"` // unit_price
	Quantity  int     `json:"quantity"`   // quantity
	Discount  float64 `json:"discount"`   // discount
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the OrderDetail exists in the database.
func (od *OrderDetail) Exists() bool {
	return od._exists
}

// Deleted returns true when the OrderDetail has been marked for deletion from
// the database.
func (od *OrderDetail) Deleted() bool {
	return od._deleted
}

// Insert inserts the OrderDetail to the database.
func (od *OrderDetail) Insert(ctx context.Context, db DB) error {
	switch {
	case od._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case od._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO order_details (` +
		`order_id, product_id, unit_price, quantity, discount` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)`
	// run
	logf(sqlstr, od.OrderID, od.ProductID, od.UnitPrice, od.Quantity, od.Discount)
	if _, err := db.ExecContext(ctx, sqlstr, od.OrderID, od.ProductID, od.UnitPrice, od.Quantity, od.Discount); err != nil {
		return logerror(err)
	}
	// set exists
	od._exists = true
	return nil
}

// Update updates a OrderDetail in the database.
func (od *OrderDetail) Update(ctx context.Context, db DB) error {
	switch {
	case !od._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case od._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE order_details SET ` +
		`unit_price = $1, quantity = $2, discount = $3` +
		` WHERE order_id = $4 AND product_id = $5`
	// run
	logf(sqlstr, od.OrderID, od.UnitPrice, od.Quantity, od.Discount, od.OrderID, od.ProductID)
	if _, err := db.ExecContext(ctx, sqlstr, od.OrderID, od.UnitPrice, od.Quantity, od.Discount, od.OrderID, od.ProductID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the OrderDetail to the database.
func (od *OrderDetail) Save(ctx context.Context, db DB) error {
	if od.Exists() {
		return od.Update(ctx, db)
	}
	return od.Insert(ctx, db)
}

// Delete deletes the OrderDetail from the database.
func (od *OrderDetail) Delete(ctx context.Context, db DB) error {
	switch {
	case !od._exists: // doesn't exist
		return nil
	case od._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM order_details WHERE order_id = $1 AND product_id = $2`
	// run
	logf(sqlstr, od.OrderID, od.ProductID)
	if _, err := db.ExecContext(ctx, sqlstr, od.OrderID, od.ProductID); err != nil {
		return logerror(err)
	}
	// set deleted
	od._deleted = true
	return nil
}

// OrderDetailByOrderIDProductID retrieves a row from 'order_details' as a OrderDetail.
//
// Generated from index 'sqlite_autoindex_order_details_1'.
func OrderDetailByOrderIDProductID(ctx context.Context, db DB, orderID, productID int) (*OrderDetail, error) {
	// query
	const sqlstr = `SELECT ` +
		`order_id, product_id, unit_price, quantity, discount ` +
		`FROM order_details ` +
		`WHERE order_id = $1 AND product_id = $2`
	// run
	logf(sqlstr, orderID, productID)
	od := OrderDetail{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, orderID, productID).Scan(&od.OrderID, &od.ProductID, &od.UnitPrice, &od.Quantity, &od.Discount); err != nil {
		return nil, logerror(err)
	}
	return &od, nil
}

// Order returns the Order associated with the OrderDetail's OrderID (order_id).
//
// Generated from foreign key 'order_details_order_id_fkey'.
func (od *OrderDetail) Order(ctx context.Context, db DB) (*Order, error) {
	return OrderByOrderID(ctx, db, od.OrderID)
}

// Product returns the Product associated with the OrderDetail's ProductID (product_id).
//
// Generated from foreign key 'order_details_product_id_fkey'.
func (od *OrderDetail) Product(ctx context.Context, db DB) (*Product, error) {
	return ProductByProductID(ctx, db, od.ProductID)
}
