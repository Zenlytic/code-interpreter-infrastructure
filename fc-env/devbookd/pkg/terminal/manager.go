package terminal

import (
	"fmt"
	"sync"
)

type TerminalManager struct {
	lock    sync.RWMutex
	termMap map[TerminalID]*Terminal
}

func NewTerminalManager() *TerminalManager {
	return &TerminalManager{
		termMap: make(map[TerminalID]*Terminal),
	}
}

func (t *TerminalManager) Remove(id TerminalID) {
	term, ok := t.Get(id)

	if !ok {
		return
	}

	term.Destroy()

	t.lock.Lock()
	defer t.lock.Unlock()

	delete(t.termMap, id)
}

func (m *TerminalManager) Get(id TerminalID) (*Terminal, bool) {
	if id == "" {
		return nil, false
	}

	m.lock.RLock()
	defer m.lock.RUnlock()

	term, ok := m.termMap[id]
	return term, ok
}

func (m *TerminalManager) Add(root string, cols, rows uint16) (*Terminal, error) {
	term, err := NewTerminal(root, cols, rows)
	if err != nil {
		return nil, fmt.Errorf("failed to create new terminal: %s", err)
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	m.termMap[term.ID] = term
	return term, nil
}
