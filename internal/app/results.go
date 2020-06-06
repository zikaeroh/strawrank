package app

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zikaeroh/strawrank/internal/ctxlog"
	"github.com/zikaeroh/strawrank/internal/db/models"
	"github.com/zikaeroh/strawrank/internal/polling"
	"github.com/zikaeroh/strawrank/internal/polling/fptp"
	"github.com/zikaeroh/strawrank/internal/polling/irv"
	"github.com/zikaeroh/strawrank/internal/templates"
	"go.uber.org/zap"
)

func (a *App) handleResults(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := ctxlog.FromContext(ctx)

	pollID := getPollIDs(r)[0]

	poll, err := models.Polls(models.PollWhere.ID.EQ(pollID), qm.Load(models.PollRels.Ballots)).One(ctx, a.db)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}

		logger.Error("error finding poll", zap.Error(err))
		a.internalServerError(w, err)
		return
	}

	ballots := make([]polling.Ballot, len(poll.R.Ballots))

	for i, b := range poll.R.Ballots {
		ballots[i] = polling.NewBallot(b.Votes...)
	}

	irvResult, irvRounds := irv.TallyWithRounds(ballots)
	fptpResult := fptp.Tally(ballots)

	irvView, err := resultToView(irvResult, poll.Choices)
	if err != nil {
		logger.Error("error creating IRV view", zap.Error(err))
		a.internalServerError(w, err)
		return
	}

	irvRoundViews, err := resultsToViews(irvRounds, poll.Choices)
	if err != nil {
		logger.Error("error creating IRV round view", zap.Error(err))
		a.internalServerError(w, err)
		return
	}

	fptpView, err := resultToView(fptpResult, poll.Choices)
	if err != nil {
		logger.Error("error creating FPTP view", zap.Error(err))
		a.internalServerError(w, err)
		return
	}

	templates.WritePageTemplate(w, &templates.ResultsPage{
		Question:  poll.Question,
		Choices:   poll.Choices,
		IRV:       irvView,
		IRVRounds: irvRoundViews,
		FPTP:      fptpView,
	})
}

func resultToView(result polling.Result, choices []string) (templates.ResultView, error) {
	view := templates.ResultView{
		Rows: make([]templates.ResultRow, len(result.Ranking)),
	}

	for i, c := range result.Ranking {
		view.Rows[i] = templates.ResultRow{
			Name:  choices[c.ID],
			Count: c.Count,
		}
	}

	limit, other, shouldLimit := chartLimit(view.Rows)

	var chart []templates.ResultRow

	if shouldLimit {
		chart = make([]templates.ResultRow, limit, limit+1)
		copy(chart, view.Rows[:limit])
		chart = append(chart, templates.ResultRow{
			Name:  "...",
			Count: other,
		})
	} else {
		chart = view.Rows
	}

	var err error
	view.ChartData, err = rowsToChart(chart)

	if err != nil {
		return templates.ResultView{}, err
	}

	return view, nil
}

func resultsToViews(results []polling.Result, choices []string) ([]templates.ResultView, error) {
	views := make([]templates.ResultView, len(results))

	for i, r := range results {
		var err error
		views[i], err = resultToView(r, choices)
		if err != nil {
			return nil, err
		}
	}

	return views, nil
}

func rowsToChart(rows []templates.ResultRow) ([]byte, error) {
	type chartDataset struct {
		Label           string   `json:"label"`
		Data            []int    `json:"data"`
		BackgroundColor []string `json:"backgroundColor"`
	}

	type chartData struct {
		Labels   []string       `json:"labels"`
		Datasets []chartDataset `json:"datasets"`
	}

	data := chartData{
		Labels: make([]string, len(rows)),
		Datasets: []chartDataset{
			{
				Label:           "Votes",
				Data:            make([]int, len(rows)),
				BackgroundColor: make([]string, len(rows)),
			},
		},
	}

	for i, row := range rows {
		data.Labels[i] = row.Name
		data.Datasets[0].Data[i] = row.Count

		switch {
		case i == 0:
			data.Datasets[0].BackgroundColor[i] = "#00bc8c"
		case row.Name == "...":
			data.Datasets[0].BackgroundColor[i] = "#444"
		default:
			data.Datasets[0].BackgroundColor[i] = "#375a7f"
		}
	}

	return json.Marshal(data)
}

func chartLimit(rr []templates.ResultRow) (int, int, bool) {
	const min = 5

	if len(rr) <= min {
		return 0, 0, false
	}

	sum := 0

	i := len(rr) - 1
	for ; i > min; i-- {
		c := rr[i]
		prev := rr[i-1]

		sum += c.Count

		if sum >= prev.Count {
			sum -= c.Count
			break
		}
	}

	return i, sum, true
}
