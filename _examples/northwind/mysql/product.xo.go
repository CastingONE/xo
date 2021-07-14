package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Product represents a row from 'northwind.products'.
type Product struct {
	ProductID       int16           `json:"product_id"`        // product_id
	ProductName     string          `json:"product_name"`      // product_name
	SupplierID      sql.NullInt64   `json:"supplier_id"`       // supplier_id
	CategoryID      sql.NullInt64   `json:"category_id"`       // category_id
	QuantityPerUnit sql.NullString  `json:"quantity_per_unit"` // quantity_per_unit
	UnitPrice       sql.NullFloat64 `json:"unit_price"`        // unit_price
	UnitsInStock    sql.NullInt64   `json:"units_in_stock"`    // units_in_stock
	UnitsOnOrder    sql.NullInt64   `json:"units_on_order"`    // units_on_order
	ReorderLevel    sql.NullInt64   `json:"reorder_level"`     // reorder_level
	Discontinued    int             `json:"discontinued"`      // discontinued
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Product exists in the database.
func (p *Product) Exists() bool {
	return p._exists
}

// Deleted returns true when the Product has been marked for deletion from
// the database.
func (p *Product) Deleted() bool {
	return p._deleted
}

// Insert inserts the Product to the database.
func (p *Product) Insert(ctx context.Context, db DB) error {
	switch {
	case p._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO northwind.products (` +
		`product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, p.ProductID, p.ProductName, p.SupplierID, p.CategoryID, p.QuantityPerUnit, p.UnitPrice, p.UnitsInStock, p.UnitsOnOrder, p.ReorderLevel, p.Discontinued)
	if _, err := db.ExecContext(ctx, sqlstr, p.ProductID, p.ProductName, p.SupplierID, p.CategoryID, p.QuantityPerUnit, p.UnitPrice, p.UnitsInStock, p.UnitsOnOrder, p.ReorderLevel, p.Discontinued); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Update updates a Product in the database.
func (p *Product) Update(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE northwind.products SET ` +
		`product_name = ?, supplier_id = ?, category_id = ?, quantity_per_unit = ?, unit_price = ?, units_in_stock = ?, units_on_order = ?, reorder_level = ?, discontinued = ? ` +
		`WHERE product_id = ?`
	// run
	logf(sqlstr, p.ProductName, p.SupplierID, p.CategoryID, p.QuantityPerUnit, p.UnitPrice, p.UnitsInStock, p.UnitsOnOrder, p.ReorderLevel, p.Discontinued, p.ProductID)
	if _, err := db.ExecContext(ctx, sqlstr, p.ProductName, p.SupplierID, p.CategoryID, p.QuantityPerUnit, p.UnitPrice, p.UnitsInStock, p.UnitsOnOrder, p.ReorderLevel, p.Discontinued, p.ProductID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Product to the database.
func (p *Product) Save(ctx context.Context, db DB) error {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for Product.
func (p *Product) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO northwind.products (` +
		`product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`product_id = VALUES(product_id), product_name = VALUES(product_name), supplier_id = VALUES(supplier_id), category_id = VALUES(category_id), quantity_per_unit = VALUES(quantity_per_unit), unit_price = VALUES(unit_price), units_in_stock = VALUES(units_in_stock), units_on_order = VALUES(units_on_order), reorder_level = VALUES(reorder_level), discontinued = VALUES(discontinued)`
	// run
	logf(sqlstr, p.ProductID, p.ProductName, p.SupplierID, p.CategoryID, p.QuantityPerUnit, p.UnitPrice, p.UnitsInStock, p.UnitsOnOrder, p.ReorderLevel, p.Discontinued)
	if _, err := db.ExecContext(ctx, sqlstr, p.ProductID, p.ProductName, p.SupplierID, p.CategoryID, p.QuantityPerUnit, p.UnitPrice, p.UnitsInStock, p.UnitsOnOrder, p.ReorderLevel, p.Discontinued); err != nil {
		return err
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the Product from the database.
func (p *Product) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM northwind.products ` +
		`WHERE product_id = ?`
	// run
	logf(sqlstr, p.ProductID)
	if _, err := db.ExecContext(ctx, sqlstr, p.ProductID); err != nil {
		return logerror(err)
	}
	// set deleted
	p._deleted = true
	return nil
}

// ProductsByCategoryID retrieves a row from 'northwind.products' as a Product.
//
// Generated from index 'category_id'.
func ProductsByCategoryID(ctx context.Context, db DB, categoryID sql.NullInt64) ([]*Product, error) {
	// query
	const sqlstr = `SELECT ` +
		`product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued ` +
		`FROM northwind.products ` +
		`WHERE category_id = ?`
	// run
	logf(sqlstr, categoryID)
	rows, err := db.QueryContext(ctx, sqlstr, categoryID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Product
	for rows.Next() {
		p := Product{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&p.ProductID, &p.ProductName, &p.SupplierID, &p.CategoryID, &p.QuantityPerUnit, &p.UnitPrice, &p.UnitsInStock, &p.UnitsOnOrder, &p.ReorderLevel, &p.Discontinued); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ProductByProductID retrieves a row from 'northwind.products' as a Product.
//
// Generated from index 'products_product_id_pkey'.
func ProductByProductID(ctx context.Context, db DB, productID int16) (*Product, error) {
	// query
	const sqlstr = `SELECT ` +
		`product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued ` +
		`FROM northwind.products ` +
		`WHERE product_id = ?`
	// run
	logf(sqlstr, productID)
	p := Product{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, productID).Scan(&p.ProductID, &p.ProductName, &p.SupplierID, &p.CategoryID, &p.QuantityPerUnit, &p.UnitPrice, &p.UnitsInStock, &p.UnitsOnOrder, &p.ReorderLevel, &p.Discontinued); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}

// ProductsBySupplierID retrieves a row from 'northwind.products' as a Product.
//
// Generated from index 'supplier_id'.
func ProductsBySupplierID(ctx context.Context, db DB, supplierID sql.NullInt64) ([]*Product, error) {
	// query
	const sqlstr = `SELECT ` +
		`product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued ` +
		`FROM northwind.products ` +
		`WHERE supplier_id = ?`
	// run
	logf(sqlstr, supplierID)
	rows, err := db.QueryContext(ctx, sqlstr, supplierID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Product
	for rows.Next() {
		p := Product{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&p.ProductID, &p.ProductName, &p.SupplierID, &p.CategoryID, &p.QuantityPerUnit, &p.UnitPrice, &p.UnitsInStock, &p.UnitsOnOrder, &p.ReorderLevel, &p.Discontinued); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Supplier returns the Supplier associated with the Product's (SupplierID).
//
// Generated from foreign key 'products_ibfk_1'.
func (p *Product) Supplier(ctx context.Context, db DB) (*Supplier, error) {
	return SupplierBySupplierID(ctx, db, int16(p.SupplierID.Int64))
}

// Category returns the Category associated with the Product's (CategoryID).
//
// Generated from foreign key 'products_ibfk_2'.
func (p *Product) Category(ctx context.Context, db DB) (*Category, error) {
	return CategoryByCategoryID(ctx, db, int16(p.CategoryID.Int64))
}
