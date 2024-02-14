package jobs

import (
	"context"

	"github.com/riverqueue/river"
	"github.com/ugent-library/projects-service/indexes"
	"github.com/ugent-library/projects-service/repositories"
)

type ReindexProjectsArgs struct{}

func (ReindexProjectsArgs) Kind() string { return "reindexProjects" }

func (ReindexProjectsArgs) InsertOpts() river.InsertOpts {
	return river.InsertOpts{
		UniqueOpts: river.UniqueOpts{
			ByQueue: true,
		},
	}
}

type ReindexProjectsWorker struct {
	river.WorkerDefaults[ReindexProjectsArgs]
	repo  *repositories.Repo
	index *indexes.Index
}

func NewReindexProjectsWorker(repo *repositories.Repo, index *indexes.Index) *ReindexProjectsWorker {
	return &ReindexProjectsWorker{repo: repo, index: index}
}

func (w *ReindexProjectsWorker) Work(ctx context.Context, job *river.Job[ReindexProjectsArgs]) error {
	return w.index.ReindexProjects(ctx, w.repo.EachProject)
}
