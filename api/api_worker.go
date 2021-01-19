package api

import (
	"context"

	"github.com/google/uuid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"

	"github.com/filecoin-project/lotus/build"

	sectorstorage "github.com/filecoin-project/lotus/extern/sector-storage"
)

type WorkerAPI interface {
	Version(context.Context) (build.Version, error)
	// TODO: Info() (name, ...) ?

	TaskTypes(context.Context) (map[sealtasks.TaskType]struct{}, error) // TaskType -> Weight
	Paths(context.Context) ([]stores.StoragePath, error)
	Info(context.Context) (storiface.WorkerInfo, error)

	storiface.WorkerCalls

	// Storage / Other
	Remove(ctx context.Context, sector abi.SectorID) error

	StorageAddLocal(ctx context.Context, path string) error

	Fetch(context.Context, abi.SectorID, stores.SectorFileType, stores.PathType, stores.AcquireMode) error

	Closing(context.Context) (<-chan struct{}, error)

	AllowableRange(ctx context.Context, task sealtasks.TaskType) (bool, error)

	AddRange(ctx context.Context, task sealtasks.TaskType, addType int) error

	GetWorkerInfo(ctx context.Context) sectorstorage.WorkerInfo

	AddStore(ctx context.Context, ID abi.SectorID, taskType sealtasks.TaskType) error

	DeleteStore(ctx context.Context, ID abi.SectorID, taskType sealtasks.TaskType) error

	SetWorkerParams(ctx context.Context, key string, val string) error

	GetWorkerGroup(ctx context.Context) string
	Session(context.Context) (uuid.UUID, error)
}
