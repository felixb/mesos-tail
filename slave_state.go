package main

type SlaveState struct {
	Id                  string
	Pid                 string
	Hostname            string
	Frameworks          []Framework
	CompletedFrameworks []Framework `json:"compleded_frameworks"`
}

func (state *SlaveState) Slave() *Slave {
	return &Slave{
		Id: state.Id,
		Pid: state.Pid,
		Hostname: state.Hostname,
	}
}

func (state *SlaveState) Framekwork(id string) *Framework {
	for _, fw := range state.Frameworks {
		if fw.Id == id {
			return &fw
		}
	}
	for _, fw := range state.CompletedFrameworks {
		if fw.Id == id {
			return &fw
		}
	}
	return nil
}

func FetchSlaveState(slave *Slave) (*SlaveState, error) {
	var state SlaveState
	err := fetchJson(slave.StateUrl(), &state)
	return &state, err
}
