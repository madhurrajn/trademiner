package main

import (
	"miner/auth"
	"miner/score"
	"testing"
)

func TestRefresh(t *testing.T) {
	auth.Init()

	score.Refresh()
}
