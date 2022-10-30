package ping

import (
	"fmt"
	"time"

	goconfig "github.com/budimanlai/go-config"
	"github.com/eqto/dbm"
	_ "github.com/eqto/dbm/driver/mysql"
)

type ServicePing struct {
	Config      *goconfig.Config
	Db          *dbm.Connection
	Indentifier string
	Interval    int
}

const (
	YYYYMMDDHHMMSS = "2006-01-02 15:04:05"
)

func (s *ServicePing) Init(configFile string, id string) error {
	if s.Config == nil {
		s.Config = &goconfig.Config{}
		e := s.Config.Open(configFile)
		if e != nil {
			return e
		}
	}

	if s.Db == nil {
		e1 := s.OpenDatabase()
		if e1 != nil {
			return e1
		}
	}

	s.Indentifier = id
	return nil
}

func (s *ServicePing) OpenDatabase() error {
	cn, e := dbm.Connect("mysql", s.Config.GetString("iam.hostname"), s.Config.GetInt("iam.port"),
		s.Config.GetString("iam.username"), s.Config.GetString("iam.password"), s.Config.GetString("iam.database"))
	if e != nil {
		return e
	}
	s.Db = cn
	return nil
}

func (s *ServicePing) Start() {
	model, e := s.Db.Get("SELECT * FROM services WHERE indentifier = ?", s.Indentifier)
	if e != nil {
		fmt.Println("Error:", e)
		return
	}

	if model == nil {
		_, e1 := s.Db.Insert("services", map[string]interface{}{
			`indentifier`:    s.Indentifier,
			`start_datetime`: s.getDate(),
		})
		if e1 != nil {
			fmt.Println("Error:", e1)
			return
		}
	} else {
		_, e2 := s.Db.Exec(`UPDATE services SET start_datetime = ? WHERE indentifier = ?`, s.getDate(), s.Indentifier)
		if e2 != nil {
			fmt.Println("Error:", e2)
			return
		}
	}
}

func (s *ServicePing) Stop() {
	_, e2 := s.Db.Exec(`UPDATE services SET stop_datetime = ? WHERE indentifier = ?`, s.getDate(), s.Indentifier)
	if e2 != nil {
		fmt.Println("Error:", e2)
		return
	}
}

func (s *ServicePing) Update() {
	_, e2 := s.Db.Exec(`UPDATE services SET last_datetime = ? WHERE indentifier = ?`, s.getDate(), s.Indentifier)
	if e2 != nil {
		fmt.Println("Error:", e2)
		return
	}
}

func (s *ServicePing) getDate() string {
	now := time.Now()
	return now.Format(YYYYMMDDHHMMSS)
}
