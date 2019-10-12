package gitcheck

import (
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// Metrics are a summary of metrics from a git repo over a time window
type Metrics struct {
	Period          TimeWindow
	UniqueCommiters map[string]object.Signature
	Commits         []*object.Commit
	LatestCommit    time.Time
}

// TimeWindow keeps track of a start and end time
type TimeWindow struct {
	Start time.Time
	End   time.Time
}

// GetMetrics gets the metrics of a given git repository url
func GetMetrics(url string, days int) (*Metrics, error) {
	now := time.Now()
	limit := now.AddDate(0, 0, -days)
	metrics := &Metrics{
		Period:          TimeWindow{Start: limit, End: now},
		UniqueCommiters: make(map[string]object.Signature),
	}
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return nil, err
	}
	cIter, err := r.CommitObjects()
	if err != nil {
		return nil, err
	}
	err = cIter.ForEach(func(commit *object.Commit) error {
		if commit.Committer.When.After(limit) {
			metrics.Commits = append(metrics.Commits, commit)
			metrics.UniqueCommiters[commit.Committer.String()] = commit.Committer
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return metrics, nil
}
