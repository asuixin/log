package log

import (
	"sync"

	seelog "github.com/cihub/seelog"
)

type MLog struct {
	sync.RWMutex
	path string
	log  seelog.LoggerInterface
}

func NewMLog(path string) (*MLog, error) {
	logger, err := seelog.LoggerFromConfigAsFile(path)
	if err != nil {
		return nil, err
	}

	return &MLog{
		path: path,
		log:  logger,
	}, nil
}

func (m *MLog) Reload() error {

	logger, err := seelog.LoggerFromConfigAsFile(m.path)
	if err != nil {
		return err
	}

	err = seelog.ReplaceLogger(logger)
	if err != nil {
		return err
	}
	defer seelog.Flush()

	m.Lock()
	m.log = logger
	m.Unlock()
	return nil
}

func (m *MLog) Logger() seelog.LoggerInterface {
	m.RLock()
	logger := m.log
	m.RUnlock()
	return logger
}
