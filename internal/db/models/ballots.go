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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Ballot is an object representing the database table.
type Ballot struct {
	ID        int64            `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time        `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time        `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	PollID    int64            `boil:"poll_id" json:"poll_id" toml:"poll_id" yaml:"poll_id"`
	Cookie    null.String      `boil:"cookie" json:"cookie,omitempty" toml:"cookie" yaml:"cookie,omitempty"`
	IPAddr    null.String      `boil:"ip_addr" json:"ip_addr,omitempty" toml:"ip_addr" yaml:"ip_addr,omitempty"`
	Votes     types.Int64Array `boil:"votes" json:"votes,omitempty" toml:"votes" yaml:"votes,omitempty"`

	R *ballotR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ballotL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BallotColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	PollID    string
	Cookie    string
	IPAddr    string
	Votes     string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	PollID:    "poll_id",
	Cookie:    "cookie",
	IPAddr:    "ip_addr",
	Votes:     "votes",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
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

type whereHelpertypes_Int64Array struct{ field string }

func (w whereHelpertypes_Int64Array) EQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_Int64Array) NEQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_Int64Array) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_Int64Array) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpertypes_Int64Array) LT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Int64Array) LTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Int64Array) GT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Int64Array) GTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var BallotWhere = struct {
	ID        whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	PollID    whereHelperint64
	Cookie    whereHelpernull_String
	IPAddr    whereHelpernull_String
	Votes     whereHelpertypes_Int64Array
}{
	ID:        whereHelperint64{field: "\"ballots\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"ballots\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"ballots\".\"updated_at\""},
	PollID:    whereHelperint64{field: "\"ballots\".\"poll_id\""},
	Cookie:    whereHelpernull_String{field: "\"ballots\".\"cookie\""},
	IPAddr:    whereHelpernull_String{field: "\"ballots\".\"ip_addr\""},
	Votes:     whereHelpertypes_Int64Array{field: "\"ballots\".\"votes\""},
}

// BallotRels is where relationship names are stored.
var BallotRels = struct {
	Poll string
}{
	Poll: "Poll",
}

// ballotR is where relationships are stored.
type ballotR struct {
	Poll *Poll `boil:"Poll" json:"Poll" toml:"Poll" yaml:"Poll"`
}

// NewStruct creates a new relationship struct
func (*ballotR) NewStruct() *ballotR {
	return &ballotR{}
}

// ballotL is where Load methods for each relationship are stored.
type ballotL struct{}

var (
	ballotAllColumns            = []string{"id", "created_at", "updated_at", "poll_id", "cookie", "ip_addr", "votes"}
	ballotColumnsWithoutDefault = []string{"poll_id", "cookie", "ip_addr", "votes"}
	ballotColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	ballotPrimaryKeyColumns     = []string{"id"}
)

type (
	// BallotSlice is an alias for a slice of pointers to Ballot.
	// This should generally be used opposed to []Ballot.
	BallotSlice []*Ballot

	ballotQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ballotType                 = reflect.TypeOf(&Ballot{})
	ballotMapping              = queries.MakeStructMapping(ballotType)
	ballotPrimaryKeyMapping, _ = queries.BindMapping(ballotType, ballotMapping, ballotPrimaryKeyColumns)
	ballotInsertCacheMut       sync.RWMutex
	ballotInsertCache          = make(map[string]insertCache)
	ballotUpdateCacheMut       sync.RWMutex
	ballotUpdateCache          = make(map[string]updateCache)
	ballotUpsertCacheMut       sync.RWMutex
	ballotUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single ballot record from the query.
func (q ballotQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Ballot, error) {
	o := &Ballot{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ballots")
	}

	return o, nil
}

// All returns all Ballot records from the query.
func (q ballotQuery) All(ctx context.Context, exec boil.ContextExecutor) (BallotSlice, error) {
	var o []*Ballot

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Ballot slice")
	}

	return o, nil
}

// Count returns the count of all Ballot records in the query.
func (q ballotQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ballots rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ballotQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ballots exists")
	}

	return count > 0, nil
}

// Poll pointed to by the foreign key.
func (o *Ballot) Poll(mods ...qm.QueryMod) pollQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PollID),
	}

	queryMods = append(queryMods, mods...)

	query := Polls(queryMods...)
	queries.SetFrom(query.Query, "\"polls\"")

	return query
}

// LoadPoll allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ballotL) LoadPoll(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBallot interface{}, mods queries.Applicator) error {
	var slice []*Ballot
	var object *Ballot

	if singular {
		object = maybeBallot.(*Ballot)
	} else {
		slice = *maybeBallot.(*[]*Ballot)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ballotR{}
		}
		args = append(args, object.PollID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ballotR{}
			}

			for _, a := range args {
				if a == obj.PollID {
					continue Outer
				}
			}

			args = append(args, obj.PollID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`polls`),
		qm.WhereIn(`polls.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Poll")
	}

	var resultSlice []*Poll
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Poll")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for polls")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for polls")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Poll = foreign
		if foreign.R == nil {
			foreign.R = &pollR{}
		}
		foreign.R.Ballots = append(foreign.R.Ballots, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PollID == foreign.ID {
				local.R.Poll = foreign
				if foreign.R == nil {
					foreign.R = &pollR{}
				}
				foreign.R.Ballots = append(foreign.R.Ballots, local)
				break
			}
		}
	}

	return nil
}

// SetPoll of the ballot to the related item.
// Sets o.R.Poll to related.
// Adds o to related.R.Ballots.
func (o *Ballot) SetPoll(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Poll) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ballots\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"poll_id"}),
		strmangle.WhereClause("\"", "\"", 2, ballotPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PollID = related.ID
	if o.R == nil {
		o.R = &ballotR{
			Poll: related,
		}
	} else {
		o.R.Poll = related
	}

	if related.R == nil {
		related.R = &pollR{
			Ballots: BallotSlice{o},
		}
	} else {
		related.R.Ballots = append(related.R.Ballots, o)
	}

	return nil
}

// Ballots retrieves all the records using an executor.
func Ballots(mods ...qm.QueryMod) ballotQuery {
	mods = append(mods, qm.From("\"ballots\""))
	return ballotQuery{NewQuery(mods...)}
}

// FindBallot retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBallot(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Ballot, error) {
	ballotObj := &Ballot{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ballots\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, ballotObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ballots")
	}

	return ballotObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Ballot) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ballots provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(ballotColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ballotInsertCacheMut.RLock()
	cache, cached := ballotInsertCache[key]
	ballotInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ballotAllColumns,
			ballotColumnsWithDefault,
			ballotColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ballotType, ballotMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ballotType, ballotMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ballots\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ballots\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into ballots")
	}

	if !cached {
		ballotInsertCacheMut.Lock()
		ballotInsertCache[key] = cache
		ballotInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Ballot.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Ballot) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	ballotUpdateCacheMut.RLock()
	cache, cached := ballotUpdateCache[key]
	ballotUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ballotAllColumns,
			ballotPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update ballots, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ballots\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ballotPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ballotType, ballotMapping, append(wl, ballotPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update ballots row")
	}

	if !cached {
		ballotUpdateCacheMut.Lock()
		ballotUpdateCache[key] = cache
		ballotUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q ballotQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for ballots")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BallotSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ballotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ballots\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ballotPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in ballot slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Ballot) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ballots provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(ballotColumnsWithDefault, o)

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

	ballotUpsertCacheMut.RLock()
	cache, cached := ballotUpsertCache[key]
	ballotUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ballotAllColumns,
			ballotColumnsWithDefault,
			ballotColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			ballotAllColumns,
			ballotPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert ballots, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ballotPrimaryKeyColumns))
			copy(conflict, ballotPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"ballots\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(ballotType, ballotMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ballotType, ballotMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert ballots")
	}

	if !cached {
		ballotUpsertCacheMut.Lock()
		ballotUpsertCache[key] = cache
		ballotUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Ballot record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Ballot) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("models: no Ballot provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ballotPrimaryKeyMapping)
	sql := "DELETE FROM \"ballots\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from ballots")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q ballotQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("models: no ballotQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from ballots")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BallotSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ballotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ballots\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ballotPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from ballot slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Ballot) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBallot(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BallotSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BallotSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ballotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ballots\".* FROM \"ballots\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ballotPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BallotSlice")
	}

	*o = slice

	return nil
}

// BallotExists checks if the Ballot row exists.
func BallotExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ballots\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ballots exists")
	}

	return exists, nil
}
