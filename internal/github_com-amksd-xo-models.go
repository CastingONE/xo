// Code generated by 'yaegi extract github.com/CastingONE/xo/models'. DO NOT EDIT.

package internal

import (
	"context"
	"database/sql"
	"github.com/CastingONE/xo/models"
	"reflect"
)

func init() {
	Symbols["github.com/CastingONE/xo/models/models"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrAlreadyExists":          reflect.ValueOf(models.ErrAlreadyExists),
		"ErrDoesNotExist":           reflect.ValueOf(models.ErrDoesNotExist),
		"ErrInvalidStringSlice":     reflect.ValueOf(models.ErrInvalidStringSlice),
		"ErrMarkedForDeletion":      reflect.ValueOf(models.ErrMarkedForDeletion),
		"Errorf":                    reflect.ValueOf(models.Errorf),
		"Logf":                      reflect.ValueOf(models.Logf),
		"MysqlEnumValues":           reflect.ValueOf(models.MysqlEnumValues),
		"MysqlEnums":                reflect.ValueOf(models.MysqlEnums),
		"MysqlIndexColumns":         reflect.ValueOf(models.MysqlIndexColumns),
		"MysqlProcParams":           reflect.ValueOf(models.MysqlProcParams),
		"MysqlProcs":                reflect.ValueOf(models.MysqlProcs),
		"MysqlSchema":               reflect.ValueOf(models.MysqlSchema),
		"MysqlTableColumns":         reflect.ValueOf(models.MysqlTableColumns),
		"MysqlTableForeignKeys":     reflect.ValueOf(models.MysqlTableForeignKeys),
		"MysqlTableIndexes":         reflect.ValueOf(models.MysqlTableIndexes),
		"MysqlTableSequences":       reflect.ValueOf(models.MysqlTableSequences),
		"MysqlTables":               reflect.ValueOf(models.MysqlTables),
		"MysqlViewCreate":           reflect.ValueOf(models.MysqlViewCreate),
		"MysqlViewDrop":             reflect.ValueOf(models.MysqlViewDrop),
		"OracleIndexColumns":        reflect.ValueOf(models.OracleIndexColumns),
		"OracleProcParams":          reflect.ValueOf(models.OracleProcParams),
		"OracleProcs":               reflect.ValueOf(models.OracleProcs),
		"OracleSchema":              reflect.ValueOf(models.OracleSchema),
		"OracleTableColumns":        reflect.ValueOf(models.OracleTableColumns),
		"OracleTableForeignKeys":    reflect.ValueOf(models.OracleTableForeignKeys),
		"OracleTableIndexes":        reflect.ValueOf(models.OracleTableIndexes),
		"OracleTableSequences":      reflect.ValueOf(models.OracleTableSequences),
		"OracleTables":              reflect.ValueOf(models.OracleTables),
		"OracleViewCreate":          reflect.ValueOf(models.OracleViewCreate),
		"OracleViewDrop":            reflect.ValueOf(models.OracleViewDrop),
		"OracleViewTruncate":        reflect.ValueOf(models.OracleViewTruncate),
		"PostgresEnumValues":        reflect.ValueOf(models.PostgresEnumValues),
		"PostgresEnums":             reflect.ValueOf(models.PostgresEnums),
		"PostgresGetColOrder":       reflect.ValueOf(models.PostgresGetColOrder),
		"PostgresIndexColumns":      reflect.ValueOf(models.PostgresIndexColumns),
		"PostgresProcParams":        reflect.ValueOf(models.PostgresProcParams),
		"PostgresProcs":             reflect.ValueOf(models.PostgresProcs),
		"PostgresSchema":            reflect.ValueOf(models.PostgresSchema),
		"PostgresTableColumns":      reflect.ValueOf(models.PostgresTableColumns),
		"PostgresTableForeignKeys":  reflect.ValueOf(models.PostgresTableForeignKeys),
		"PostgresTableIndexes":      reflect.ValueOf(models.PostgresTableIndexes),
		"PostgresTableSequences":    reflect.ValueOf(models.PostgresTableSequences),
		"PostgresTables":            reflect.ValueOf(models.PostgresTables),
		"PostgresViewCreate":        reflect.ValueOf(models.PostgresViewCreate),
		"PostgresViewDrop":          reflect.ValueOf(models.PostgresViewDrop),
		"PostgresViewSchema":        reflect.ValueOf(models.PostgresViewSchema),
		"SetErrorLogger":            reflect.ValueOf(models.SetErrorLogger),
		"SetLogger":                 reflect.ValueOf(models.SetLogger),
		"Sqlite3IndexColumns":       reflect.ValueOf(models.Sqlite3IndexColumns),
		"Sqlite3Schema":             reflect.ValueOf(models.Sqlite3Schema),
		"Sqlite3TableColumns":       reflect.ValueOf(models.Sqlite3TableColumns),
		"Sqlite3TableForeignKeys":   reflect.ValueOf(models.Sqlite3TableForeignKeys),
		"Sqlite3TableIndexes":       reflect.ValueOf(models.Sqlite3TableIndexes),
		"Sqlite3TableSequences":     reflect.ValueOf(models.Sqlite3TableSequences),
		"Sqlite3Tables":             reflect.ValueOf(models.Sqlite3Tables),
		"Sqlite3ViewCreate":         reflect.ValueOf(models.Sqlite3ViewCreate),
		"Sqlite3ViewDrop":           reflect.ValueOf(models.Sqlite3ViewDrop),
		"SqlserverIndexColumns":     reflect.ValueOf(models.SqlserverIndexColumns),
		"SqlserverProcParams":       reflect.ValueOf(models.SqlserverProcParams),
		"SqlserverProcs":            reflect.ValueOf(models.SqlserverProcs),
		"SqlserverSchema":           reflect.ValueOf(models.SqlserverSchema),
		"SqlserverTableColumns":     reflect.ValueOf(models.SqlserverTableColumns),
		"SqlserverTableForeignKeys": reflect.ValueOf(models.SqlserverTableForeignKeys),
		"SqlserverTableIndexes":     reflect.ValueOf(models.SqlserverTableIndexes),
		"SqlserverTableSequences":   reflect.ValueOf(models.SqlserverTableSequences),
		"SqlserverTables":           reflect.ValueOf(models.SqlserverTables),
		"SqlserverViewCreate":       reflect.ValueOf(models.SqlserverViewCreate),
		"SqlserverViewDrop":         reflect.ValueOf(models.SqlserverViewDrop),

		// type definitions
		"Column":           reflect.ValueOf((*models.Column)(nil)),
		"DB":               reflect.ValueOf((*models.DB)(nil)),
		"Enum":             reflect.ValueOf((*models.Enum)(nil)),
		"EnumValue":        reflect.ValueOf((*models.EnumValue)(nil)),
		"ErrDecodeFailed":  reflect.ValueOf((*models.ErrDecodeFailed)(nil)),
		"ErrInsertFailed":  reflect.ValueOf((*models.ErrInsertFailed)(nil)),
		"ErrUpdateFailed":  reflect.ValueOf((*models.ErrUpdateFailed)(nil)),
		"ErrUpsertFailed":  reflect.ValueOf((*models.ErrUpsertFailed)(nil)),
		"Error":            reflect.ValueOf((*models.Error)(nil)),
		"ForeignKey":       reflect.ValueOf((*models.ForeignKey)(nil)),
		"Index":            reflect.ValueOf((*models.Index)(nil)),
		"IndexColumn":      reflect.ValueOf((*models.IndexColumn)(nil)),
		"MysqlEnumValue":   reflect.ValueOf((*models.MysqlEnumValue)(nil)),
		"PostgresColOrder": reflect.ValueOf((*models.PostgresColOrder)(nil)),
		"Proc":             reflect.ValueOf((*models.Proc)(nil)),
		"ProcParam":        reflect.ValueOf((*models.ProcParam)(nil)),
		"Sequence":         reflect.ValueOf((*models.Sequence)(nil)),
		"StringSlice":      reflect.ValueOf((*models.StringSlice)(nil)),
		"Table":            reflect.ValueOf((*models.Table)(nil)),

		// interface wrapper definitions
		"_DB": reflect.ValueOf((*_github_com_CastingONE_xo_models_DB)(nil)),
	}
}

// _github_com_CastingONE_xo_models_DB is an interface wrapper for DB type
type _github_com_CastingONE_xo_models_DB struct {
	IValue           interface{}
	WExecContext     func(a0 context.Context, a1 string, a2 ...interface{}) (sql.Result, error)
	WQueryContext    func(a0 context.Context, a1 string, a2 ...interface{}) (*sql.Rows, error)
	WQueryRowContext func(a0 context.Context, a1 string, a2 ...interface{}) *sql.Row
}

func (W _github_com_CastingONE_xo_models_DB) ExecContext(a0 context.Context, a1 string, a2 ...interface{}) (sql.Result, error) {
	return W.WExecContext(a0, a1, a2...)
}
func (W _github_com_CastingONE_xo_models_DB) QueryContext(a0 context.Context, a1 string, a2 ...interface{}) (*sql.Rows, error) {
	return W.WQueryContext(a0, a1, a2...)
}
func (W _github_com_CastingONE_xo_models_DB) QueryRowContext(a0 context.Context, a1 string, a2 ...interface{}) *sql.Row {
	return W.WQueryRowContext(a0, a1, a2...)
}
