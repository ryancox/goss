package resource

import (
	"github.com/aelsabbahy/goss/system"
	"github.com/aelsabbahy/goss/util"
)

type Service struct {
	Service string `json:"-"`
	Enabled bool   `json:"enabled"`
	Running bool   `json:"running"`
}

func (s *Service) ID() string      { return s.Service }
func (s *Service) SetID(id string) { s.Service = id }

func (s *Service) Validate(sys *system.System) []TestResult {
	sysservice := sys.NewService(s.Service, sys, util.Config{})

	var results []TestResult

	results = append(results, ValidateValue(s, "enabled", s.Enabled, sysservice.Enabled))
	results = append(results, ValidateValue(s, "running", s.Running, sysservice.Running))

	return results
}

func NewService(sysService system.Service, config util.Config) (*Service, error) {
	service := sysService.Service()
	enabled, _ := sysService.Enabled()
	running, _ := sysService.Running()
	return &Service{
		Service: service,
		Enabled: enabled.(bool),
		Running: running.(bool),
	}, nil
}
