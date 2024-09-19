package eval_test

import (
	"os"
	"testing"

	"github.com/dicedb/dice/internal/logger"
	dstore "github.com/dicedb/dice/internal/store"
)

func TestMain(m *testing.M) {
	logr := logger.New(logger.Opts{WithTimestamp: false})
	logger.SetDefault(logr)

	store := dstore.NewStore(nil)
	store.ResetStore()

	exitCode := m.Run()

	os.Exit(exitCode)
}
