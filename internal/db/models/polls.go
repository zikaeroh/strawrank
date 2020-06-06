// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Poll is an object representing the database table.
type Poll struct {
	ID        int64             `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time         `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time         `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	Question  string            `boil:"question" json:"question" toml:"question" yaml:"question"`
	Choices   types.StringArray `boil:"choices" json:"choices,omitempty" toml:"choices" yaml:"choices,omitempty"`
	CheckMode string            `boil:"check_mode" json:"check_mode" toml:"check_mode" yaml:"check_mode"`

	R *pollR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pollL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PollColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	Question  string
	Choices   string
	CheckMode string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Question:  "question",
	Choices:   "choices",
	CheckMode: "check_mode",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
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
	return qm.WhereIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_StringArray) IsNull() qm.QueryMod { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_StringArray) IsNotNull() qm.QueryMod {
	return qmhelper.WhereIsNotNull(w.field)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var PollWhere = struct {
	ID        whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	Question  whereHelperstring
	Choices   whereHelpertypes_StringArray
	CheckMode whereHelperstring
}{
	ID:        whereHelperint64{field: "\"polls\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"polls\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"polls\".\"updated_at\""},
	Question:  whereHelperstring{field: "\"polls\".\"question\""},
	Choices:   whereHelpertypes_StringArray{field: "\"polls\".\"choices\""},
	CheckMode: whereHelperstring{field: "\"polls\".\"check_mode\""},
}

// PollRels is where relationship names are stored.
var PollRels = struct {
	Ballots string
}{
	Ballots: "Ballots",
}

// pollR is where relationships are stored.
type pollR struct {
	Ballots BallotSlice `boil:"Ballots" json:"Ballots" toml:"Ballots" yaml:"Ballots"`
}

// NewStruct creates a new relationship struct
func (*pollR) NewStruct() *pollR {
	return &pollR{}
}

// pollL is where Load methods for each relationship are stored.
type pollL struct{}

var (
	pollAllColumns            = []string{"id", "created_at", "updated_at", "question", "choices", "check_mode"}
	pollColumnsWithoutDefault = []string{"question", "choices"}
	pollColumnsWithDefault    = []string{"id", "created_at", "updated_at", "check_mode"}
	pollPrimaryKeyColumns     = []string{"id"}
)

type (
	// PollSlice is an alias for a slice of pointers to Poll.
	// This should generally be used opposed to []Poll.
	PollSlice []*Poll

	pollQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pollType                 = reflect.TypeOf(&Poll{})
	pollMapping              = queries.MakeStructMapping(pollType)
	pollPrimaryKeyMapping, _ = queries.BindMapping(pollType, pollMapping, pollPrimaryKeyColumns)
	pollInsertCacheMut       sync.RWMutex
	pollInsertCache          = make(map[string]insertCache)
	pollUpdateCacheMut       sync.RWMutex
	pollUpdateCache          = make(map[string]updateCache)
	pollUpsertCacheMut       sync.RWMutex
	pollUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single poll record from the query.
func (q pollQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Poll, error) {
	o := &Poll{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for polls")
	}

	return o, nil
}

// All returns all Poll records from the query.
func (q pollQuery) All(ctx context.Context, exec boil.ContextExecutor) (PollSlice, error) {
	var o []*Poll

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Poll slice")
	}

	return o, nil
}

// Count returns the count of all Poll records in the query.
func (q pollQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count polls rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pollQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if polls exists")
	}

	return count > 0, nil
}

// Ballots retrieves all the ballot's Ballots with an executor.
func (o *Poll) Ballots(mods ...qm.QueryMod) ballotQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"ballots\".\"poll_id\"=?", o.ID),
	)

	query := Ballots(queryMods...)
	queries.SetFrom(query.Query, "\"ballots\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"ballots\".*"})
	}

	return query
}

// LoadBallots allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (pollL) LoadBallots(ctx context.Context, e boil.ContextExecutor, singular bool, maybePoll interface{}, mods queries.Applicator) error {
	var slice []*Poll
	var object *Poll

	if singular {
		object = maybePoll.(*Poll)
	} else {
		slice = *maybePoll.(*[]*Poll)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pollR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pollR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`ballots`),
		qm.WhereIn(`ballots.poll_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ballots")
	}

	var resultSlice []*Ballot
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ballots")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on ballots")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for ballots")
	}

	if singular {
		object.R.Ballots = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &ballotR{}
			}
			foreign.R.Poll = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PollID {
				local.R.Ballots = append(local.R.Ballots, foreign)
				if foreign.R == nil {
					foreign.R = &ballotR{}
				}
				foreign.R.Poll = local
				break
			}
		}
	}

	return nil
}

// AddBallots adds the given related objects to the existing relationships
// of the poll, optionally inserting them as new records.
// Appends related to o.R.Ballots.
// Sets related.R.Poll appropriately.
func (o *Poll) AddBallots(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Ballot) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PollID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"ballots\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"poll_id"}),
				strmangle.WhereClause("\"", "\"", 2, ballotPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PollID = o.ID
		}
	}

	if o.R == nil {
		o.R = &pollR{
			Ballots: related,
		}
	} else {
		o.R.Ballots = append(o.R.Ballots, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &ballotR{
				Poll: o,
			}
		} else {
			rel.R.Poll = o
		}
	}
	return nil
}

// Polls retrieves all the records using an executor.
func Polls(mods ...qm.QueryMod) pollQuery {
	mods = append(mods, qm.From("\"polls\""))
	return pollQuery{NewQuery(mods...)}
}

// FindPoll retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPoll(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Poll, error) {
	pollObj := &Poll{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"polls\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, pollObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from polls")
	}

	return pollObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Poll) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no polls provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(pollColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pollInsertCacheMut.RLock()
	cache, cached := pollInsertCache[key]
	pollInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pollAllColumns,
			pollColumnsWithDefault,
			pollColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pollType, pollMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pollType, pollMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"polls\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"polls\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into polls")
	}

	if !cached {
		pollInsertCacheMut.Lock()
		pollInsertCache[key] = cache
		pollInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Poll.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Poll) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	pollUpdateCacheMut.RLock()
	cache, cached := pollUpdateCache[key]
	pollUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pollAllColumns,
			pollPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update polls, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"polls\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pollPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pollType, pollMapping, append(wl, pollPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update polls row")
	}

	if !cached {
		pollUpdateCacheMut.Lock()
		pollUpdateCache[key] = cache
		pollUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q pollQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for polls")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PollSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pollPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"polls\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pollPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in poll slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Poll) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no polls provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(pollColumnsWithDefault, o)

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

	pollUpsertCacheMut.RLock()
	cache, cached := pollUpsertCache[key]
	pollUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pollAllColumns,
			pollColumnsWithDefault,
			pollColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			pollAllColumns,
			pollPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert polls, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pollPrimaryKeyColumns))
			copy(conflict, pollPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"polls\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(pollType, pollMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pollType, pollMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert polls")
	}

	if !cached {
		pollUpsertCacheMut.Lock()
		pollUpsertCache[key] = cache
		pollUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Poll record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Poll) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no Poll provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pollPrimaryKeyMapping)
	sql := "DELETE FROM \"polls\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from polls")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q pollQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no pollQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from polls")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PollSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pollPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"polls\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pollPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from poll slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Poll) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPoll(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PollSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PollSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pollPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"polls\".* FROM \"polls\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pollPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PollSlice")
	}

	*o = slice

	return nil
}

// PollExists checks if the Poll row exists.
func PollExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"polls\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if polls exists")
	}

	return exists, nil
}
