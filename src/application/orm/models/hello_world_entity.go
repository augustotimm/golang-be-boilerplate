// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// HelloWorldEntity is an object representing the database table.
type HelloWorldEntity struct {
	ID   string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`

	R *helloWorldEntityR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L helloWorldEntityL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HelloWorldEntityColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

var HelloWorldEntityTableColumns = struct {
	ID   string
	Name string
}{
	ID:   "hello_world_entity.id",
	Name: "hello_world_entity.name",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) LIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" LIKE ?", x)
}
func (w whereHelpernull_String) NLIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT LIKE ?", x)
}
func (w whereHelpernull_String) ILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" ILIKE ?", x)
}
func (w whereHelpernull_String) NILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT ILIKE ?", x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var HelloWorldEntityWhere = struct {
	ID   whereHelperstring
	Name whereHelpernull_String
}{
	ID:   whereHelperstring{field: "\"hello_world_entity\".\"id\""},
	Name: whereHelpernull_String{field: "\"hello_world_entity\".\"name\""},
}

// HelloWorldEntityRels is where relationship names are stored.
var HelloWorldEntityRels = struct {
}{}

// helloWorldEntityR is where relationships are stored.
type helloWorldEntityR struct {
}

// NewStruct creates a new relationship struct
func (*helloWorldEntityR) NewStruct() *helloWorldEntityR {
	return &helloWorldEntityR{}
}

// helloWorldEntityL is where Load methods for each relationship are stored.
type helloWorldEntityL struct{}

var (
	helloWorldEntityAllColumns            = []string{"id", "name"}
	helloWorldEntityColumnsWithoutDefault = []string{}
	helloWorldEntityColumnsWithDefault    = []string{"id", "name"}
	helloWorldEntityPrimaryKeyColumns     = []string{"id"}
	helloWorldEntityGeneratedColumns      = []string{}
)

type (
	// HelloWorldEntitySlice is an alias for a slice of pointers to HelloWorldEntity.
	// This should almost always be used instead of []HelloWorldEntity.
	HelloWorldEntitySlice []*HelloWorldEntity
	// HelloWorldEntityHook is the signature for custom HelloWorldEntity hook methods
	HelloWorldEntityHook func(context.Context, boil.ContextExecutor, *HelloWorldEntity) error

	helloWorldEntityQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	helloWorldEntityType                 = reflect.TypeOf(&HelloWorldEntity{})
	helloWorldEntityMapping              = queries.MakeStructMapping(helloWorldEntityType)
	helloWorldEntityPrimaryKeyMapping, _ = queries.BindMapping(helloWorldEntityType, helloWorldEntityMapping, helloWorldEntityPrimaryKeyColumns)
	helloWorldEntityInsertCacheMut       sync.RWMutex
	helloWorldEntityInsertCache          = make(map[string]insertCache)
	helloWorldEntityUpdateCacheMut       sync.RWMutex
	helloWorldEntityUpdateCache          = make(map[string]updateCache)
	helloWorldEntityUpsertCacheMut       sync.RWMutex
	helloWorldEntityUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var helloWorldEntityAfterSelectMu sync.Mutex
var helloWorldEntityAfterSelectHooks []HelloWorldEntityHook

var helloWorldEntityBeforeInsertMu sync.Mutex
var helloWorldEntityBeforeInsertHooks []HelloWorldEntityHook
var helloWorldEntityAfterInsertMu sync.Mutex
var helloWorldEntityAfterInsertHooks []HelloWorldEntityHook

var helloWorldEntityBeforeUpdateMu sync.Mutex
var helloWorldEntityBeforeUpdateHooks []HelloWorldEntityHook
var helloWorldEntityAfterUpdateMu sync.Mutex
var helloWorldEntityAfterUpdateHooks []HelloWorldEntityHook

var helloWorldEntityBeforeDeleteMu sync.Mutex
var helloWorldEntityBeforeDeleteHooks []HelloWorldEntityHook
var helloWorldEntityAfterDeleteMu sync.Mutex
var helloWorldEntityAfterDeleteHooks []HelloWorldEntityHook

var helloWorldEntityBeforeUpsertMu sync.Mutex
var helloWorldEntityBeforeUpsertHooks []HelloWorldEntityHook
var helloWorldEntityAfterUpsertMu sync.Mutex
var helloWorldEntityAfterUpsertHooks []HelloWorldEntityHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *HelloWorldEntity) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range helloWorldEntityAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *HelloWorldEntity) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range helloWorldEntityBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *HelloWorldEntity) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range helloWorldEntityAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *HelloWorldEntity) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range helloWorldEntityBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *HelloWorldEntity) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range helloWorldEntityAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *HelloWorldEntity) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range helloWorldEntityBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *HelloWorldEntity) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range helloWorldEntityAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *HelloWorldEntity) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range helloWorldEntityBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *HelloWorldEntity) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range helloWorldEntityAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHelloWorldEntityHook registers your hook function for all future operations.
func AddHelloWorldEntityHook(hookPoint boil.HookPoint, helloWorldEntityHook HelloWorldEntityHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		helloWorldEntityAfterSelectMu.Lock()
		helloWorldEntityAfterSelectHooks = append(helloWorldEntityAfterSelectHooks, helloWorldEntityHook)
		helloWorldEntityAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		helloWorldEntityBeforeInsertMu.Lock()
		helloWorldEntityBeforeInsertHooks = append(helloWorldEntityBeforeInsertHooks, helloWorldEntityHook)
		helloWorldEntityBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		helloWorldEntityAfterInsertMu.Lock()
		helloWorldEntityAfterInsertHooks = append(helloWorldEntityAfterInsertHooks, helloWorldEntityHook)
		helloWorldEntityAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		helloWorldEntityBeforeUpdateMu.Lock()
		helloWorldEntityBeforeUpdateHooks = append(helloWorldEntityBeforeUpdateHooks, helloWorldEntityHook)
		helloWorldEntityBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		helloWorldEntityAfterUpdateMu.Lock()
		helloWorldEntityAfterUpdateHooks = append(helloWorldEntityAfterUpdateHooks, helloWorldEntityHook)
		helloWorldEntityAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		helloWorldEntityBeforeDeleteMu.Lock()
		helloWorldEntityBeforeDeleteHooks = append(helloWorldEntityBeforeDeleteHooks, helloWorldEntityHook)
		helloWorldEntityBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		helloWorldEntityAfterDeleteMu.Lock()
		helloWorldEntityAfterDeleteHooks = append(helloWorldEntityAfterDeleteHooks, helloWorldEntityHook)
		helloWorldEntityAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		helloWorldEntityBeforeUpsertMu.Lock()
		helloWorldEntityBeforeUpsertHooks = append(helloWorldEntityBeforeUpsertHooks, helloWorldEntityHook)
		helloWorldEntityBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		helloWorldEntityAfterUpsertMu.Lock()
		helloWorldEntityAfterUpsertHooks = append(helloWorldEntityAfterUpsertHooks, helloWorldEntityHook)
		helloWorldEntityAfterUpsertMu.Unlock()
	}
}

// One returns a single helloWorldEntity record from the query.
func (q helloWorldEntityQuery) One(ctx context.Context, exec boil.ContextExecutor) (*HelloWorldEntity, error) {
	o := &HelloWorldEntity{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for hello_world_entity")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all HelloWorldEntity records from the query.
func (q helloWorldEntityQuery) All(ctx context.Context, exec boil.ContextExecutor) (HelloWorldEntitySlice, error) {
	var o []*HelloWorldEntity

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to HelloWorldEntity slice")
	}

	if len(helloWorldEntityAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all HelloWorldEntity records in the query.
func (q helloWorldEntityQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count hello_world_entity rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q helloWorldEntityQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if hello_world_entity exists")
	}

	return count > 0, nil
}

// HelloWorldEntities retrieves all the records using an executor.
func HelloWorldEntities(mods ...qm.QueryMod) helloWorldEntityQuery {
	mods = append(mods, qm.From("\"hello_world_entity\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"hello_world_entity\".*"})
	}

	return helloWorldEntityQuery{q}
}

// FindHelloWorldEntity retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHelloWorldEntity(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*HelloWorldEntity, error) {
	helloWorldEntityObj := &HelloWorldEntity{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"hello_world_entity\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, helloWorldEntityObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from hello_world_entity")
	}

	if err = helloWorldEntityObj.doAfterSelectHooks(ctx, exec); err != nil {
		return helloWorldEntityObj, err
	}

	return helloWorldEntityObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *HelloWorldEntity) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no hello_world_entity provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(helloWorldEntityColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	helloWorldEntityInsertCacheMut.RLock()
	cache, cached := helloWorldEntityInsertCache[key]
	helloWorldEntityInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			helloWorldEntityAllColumns,
			helloWorldEntityColumnsWithDefault,
			helloWorldEntityColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(helloWorldEntityType, helloWorldEntityMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(helloWorldEntityType, helloWorldEntityMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"hello_world_entity\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"hello_world_entity\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into hello_world_entity")
	}

	if !cached {
		helloWorldEntityInsertCacheMut.Lock()
		helloWorldEntityInsertCache[key] = cache
		helloWorldEntityInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the HelloWorldEntity.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *HelloWorldEntity) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	helloWorldEntityUpdateCacheMut.RLock()
	cache, cached := helloWorldEntityUpdateCache[key]
	helloWorldEntityUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			helloWorldEntityAllColumns,
			helloWorldEntityPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update hello_world_entity, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"hello_world_entity\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, helloWorldEntityPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(helloWorldEntityType, helloWorldEntityMapping, append(wl, helloWorldEntityPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update hello_world_entity row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for hello_world_entity")
	}

	if !cached {
		helloWorldEntityUpdateCacheMut.Lock()
		helloWorldEntityUpdateCache[key] = cache
		helloWorldEntityUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q helloWorldEntityQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for hello_world_entity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for hello_world_entity")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HelloWorldEntitySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), helloWorldEntityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"hello_world_entity\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, helloWorldEntityPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in helloWorldEntity slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all helloWorldEntity")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *HelloWorldEntity) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no hello_world_entity provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(helloWorldEntityColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	helloWorldEntityUpsertCacheMut.RLock()
	cache, cached := helloWorldEntityUpsertCache[key]
	helloWorldEntityUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			helloWorldEntityAllColumns,
			helloWorldEntityColumnsWithDefault,
			helloWorldEntityColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			helloWorldEntityAllColumns,
			helloWorldEntityPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert hello_world_entity, could not build update column list")
		}

		ret := strmangle.SetComplement(helloWorldEntityAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(helloWorldEntityPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert hello_world_entity, could not build conflict column list")
			}

			conflict = make([]string, len(helloWorldEntityPrimaryKeyColumns))
			copy(conflict, helloWorldEntityPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"hello_world_entity\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(helloWorldEntityType, helloWorldEntityMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(helloWorldEntityType, helloWorldEntityMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert hello_world_entity")
	}

	if !cached {
		helloWorldEntityUpsertCacheMut.Lock()
		helloWorldEntityUpsertCache[key] = cache
		helloWorldEntityUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single HelloWorldEntity record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *HelloWorldEntity) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no HelloWorldEntity provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), helloWorldEntityPrimaryKeyMapping)
	sql := "DELETE FROM \"hello_world_entity\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from hello_world_entity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for hello_world_entity")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q helloWorldEntityQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no helloWorldEntityQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hello_world_entity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for hello_world_entity")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HelloWorldEntitySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(helloWorldEntityBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), helloWorldEntityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"hello_world_entity\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, helloWorldEntityPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from helloWorldEntity slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for hello_world_entity")
	}

	if len(helloWorldEntityAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *HelloWorldEntity) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHelloWorldEntity(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HelloWorldEntitySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HelloWorldEntitySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), helloWorldEntityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"hello_world_entity\".* FROM \"hello_world_entity\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, helloWorldEntityPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HelloWorldEntitySlice")
	}

	*o = slice

	return nil
}

// HelloWorldEntityExists checks if the HelloWorldEntity row exists.
func HelloWorldEntityExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"hello_world_entity\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if hello_world_entity exists")
	}

	return exists, nil
}

// Exists checks if the HelloWorldEntity row exists.
func (o *HelloWorldEntity) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return HelloWorldEntityExists(ctx, exec, o.ID)
}
