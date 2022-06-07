package hub

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"

	"backend/instances/models"
	"backend/utils/logger"
)

var Hub = (func() *hub {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		logger.Log(logrus.Info, "Closing hub...")
		os.Exit(0)
	}()

	logger.Log(logrus.Info, "LOADED IN-MEMORY HUB")

	return &hub{
		Increment:     0,
		IncrementLock: sync.Mutex{},
		Channels:      map[uint64]*models.Channel{},
	}
})()

type hub struct {
	Increment     uint64                     `json:"increment"`
	IncrementLock sync.Mutex                 `json:"increment_lock"`
	Channels      map[uint64]*models.Channel `json:"channels"`
}

func (r *hub) GetIncrementValue() uint64 {
	r.IncrementLock.Lock()
	r.Increment++
	increment := r.Increment
	r.IncrementLock.Unlock()
	return increment
}
