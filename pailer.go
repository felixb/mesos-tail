package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
	"os"
)

const (
	PAILER_CHUNK_SIZE = 50000
	PAILER_INTERVAL = 1 * time.Second
	PAILER_STOP_DELAY = 3 * time.Second
)

type StringWriter interface {
	WriteString(string) (int, error)
}

type Pailer struct {
	BaseUrl  string
	BasePath string
	Path     string
	Offset   int
	writer   StringWriter
	running  bool
	ticker   *time.Ticker
	wait     chan bool
}

type update struct {
	Offset int
	Data   string
}

// get a pailer pointing to task's file
func NewPailer(w StringWriter, slave *SlaveState, task *Task, path string) (*Pailer, error) {
	if w == nil {
		return nil, fmt.Errorf("w must not be nil")
	}

	fw := slave.Framekwork(task.FrameworkId)
	if fw == nil {
		return nil, fmt.Errorf("Error finding framework %s in slave state", task.FrameworkId)
	}
	ex := fw.Executor(task.Id)
	if ex == nil {
		return nil, fmt.Errorf("Error finding executor %s in slave state", task.Id)
	}

	return &Pailer{
		BaseUrl:  fmt.Sprintf("%s/files/read.json", slave.Slave().Url()),
		BasePath: ex.Directory,
		Path:     path,
		Offset:   0,
		writer:   w,
		running:  false,
		wait:     make(chan bool, 1),
	}, nil
}

// start the pailer
func (p *Pailer) Start() {
	p.running = true
	p.ticker = time.NewTicker(PAILER_INTERVAL)
	go p.tick()
}

// stop the pailer
func (p *Pailer) Stop() {
	p.running = false
}

// wait for pailer to finish last fetch
func (p *Pailer) Wait() {
	<-p.wait
	p.wait <- true
}

// fetch update via http
func (p *Pailer) fetch() (*update, error) {
	url := fmt.Sprintf("%s?length=%d&offset=%d&path=%s",
		p.BaseUrl,
		PAILER_CHUNK_SIZE,
		p.Offset,
		url.QueryEscape(fmt.Sprintf("%s/%s", p.BasePath, p.Path)))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return p.decode(resp.Body)
}

// decode data to struct
func (p *Pailer) decode(r io.Reader) (*update, error) {
	d := json.NewDecoder(r)
	var u update
	if err := d.Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}

// apply fetched update
func (p *Pailer) update(u *update) {
	p.Offset = u.Offset + len(u.Data)
	p.writer.WriteString(u.Data)
}

// fetch update and apply
func (p *Pailer) fetchAndUpdate() {
	u, err := p.fetch()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Fetching pailer update failed: %s\n", err))
	} else {
		p.update(u)
	}
}

// fetch updates every Xs
func (p *Pailer) tick() {
	for p.running {
		p.fetchAndUpdate()
		<-p.ticker.C
	}
	p.fetchAndUpdate()
	p.wait <- true
}
