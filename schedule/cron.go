package schedule

import (
	"github.com/robfig/cron"
)

type Job interface {
	cron.Job
	Frequency() string
}

type Manager struct {
	cron *cron.Cron
}

func NewManager() *Manager {
	c := cron.New()
	return &Manager{c}
}

func (m *Manager) Register(job ... Job) {
	for _, j := range job {
		m.cron.AddJob(j.Frequency(), j)
	}
}

func (m *Manager) RegisterFunc(frequency string, cmd func()) {
	m.cron.AddFunc(frequency, cmd)
}

func (m *Manager) Start() {
	m.cron.Start()
}

func (m *Manager) Stop() {
	m.cron.Stop()
}

func (m *Manager) Entries() []*cron.Entry {
	return m.cron.Entries()
}
