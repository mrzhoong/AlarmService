package test

import (
	"AlarmService/g"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	g.LoadConfig()
}

func TestLoadCfg(t *testing.T) {
	g.LoadCfg("cfg.json")
}
