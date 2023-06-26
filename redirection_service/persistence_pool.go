package redirection_service

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	db "github.com/elbombardi/squrl/db/sqlc"
)

type PersistencePool struct {
	WorkersCount        int
	ShortURLsRepository db.ShortURLsRepository
	ClicksRepository    db.ClicksRepository
	Stopped             bool
	jobChannel          chan *PersistenceJob
	waitGroup           *sync.WaitGroup
}

type PersistenceJob struct {
	ShortUrl *db.ShortUrl
}

func (p *PersistencePool) AddJob(shortUrl *db.ShortUrl) {
	p.jobChannel <- &PersistenceJob{
		ShortUrl: shortUrl,
	}
}

var instance *PersistencePool

func NewPersistencePool(workersCount int,
	ShortURLsRepository db.ShortURLsRepository,
	ClicksRepository db.ClicksRepository) *PersistencePool {
	instance = &PersistencePool{
		WorkersCount:        workersCount,
		ShortURLsRepository: ShortURLsRepository,
		ClicksRepository:    ClicksRepository,
		Stopped:             false,
		waitGroup:           &sync.WaitGroup{},
		jobChannel:          make(chan *PersistenceJob),
	}
	return instance
}

func (p *PersistencePool) Start() {
	p.waitGroup.Add(p.WorkersCount)
	for i := 0; i < p.WorkersCount; i++ {
		go p.worker()
	}
}

func (p *PersistencePool) Stop() {
	log.Println("Stopping Persistence Pool..")
	p.Stopped = true
	close(p.jobChannel)
	p.waitGroup.Wait()
	log.Println("Persistence Pool Stopped..")
}

func (p *PersistencePool) worker() {
	for job := range p.jobChannel {
		p.persistClickInfo(job.ShortUrl)
		if p.Stopped {
			break
		}
	}
	p.waitGroup.Done()
}

func (p *PersistencePool) persistClickInfo(shortUrl *db.ShortUrl) {
	p.ShortURLsRepository.IncrementShortURLClickCount(context.Background(), shortUrl.ID)
	p.ShortURLsRepository.SetShortURLLastClickDate(context.Background(), db.SetShortURLLastClickDateParams{
		ID: shortUrl.ID,
		LastClickDateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if !shortUrl.FirstClickDateTime.Valid {
		p.ShortURLsRepository.SetShortURLFirstClickDate(context.Background(), db.SetShortURLFirstClickDateParams{
			ID: shortUrl.ID,
			FirstClickDateTime: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		})
	}
}

func FinalizePersistencePool() {
	if instance != nil {
		instance.Stop()
	}
}
