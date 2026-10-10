package implemented

import (
	"github.com/dejitarudemon/ignicula/internal/storage"
	"github.com/dejitarudemon/ignicula/internal/storage/implemented/syncmap"
)

var (
	_ storage.Storage = (*syncmap.SyncMap)(nil)
)
