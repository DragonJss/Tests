package run

import (
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/config"
	"github.com/hectorgimenez/koolo/internal/context"
)

type Crypt struct {
	ctx *context.Status
}

func NewCrypt() *Crypt {
	return &Crypt{
		ctx: context.Get(),
	}
}

func (a Crypt) Name() string {
	return string(config.CryptRun)
}

func (a Crypt) Run() error {

	// Define a defaut filter
	monsterFilter := data.MonsterAnyFilter()

	// Update filter if we selected to clear only elites
	if a.ctx.CharacterCfg.Game.Crypt.FocusOnElitePacks {
		monsterFilter = data.MonsterEliteFilter()
	}

	// Use the waypoint
	err := action.WayPoint(area.ColdPlains)
	if err != nil {
		return err
	}

	// Move to the BurialGrounds
	if err = action.MoveToArea(area.BurialGrounds); err != nil {
		return err
	}

	// Move to the Crypt
	if err = action.MoveToArea(area.Crypt); err != nil {
		return err
	}

	// Clear the area
	return action.ClearCurrentLevel(a.ctx.CharacterCfg.Game.Crypt.OpenChests, monsterFilter)

	// Open a TP If we're the leader
	action.OpenTPIfLeader()

	// Move to the Crypt
	if err = action.MoveToArea(area.Crypt); err != nil {
		return err
	}

	// Open a TP If we're the leader
	action.OpenTPIfLeader()

	// Clear the area
	return action.ClearCurrentLevel(a.ctx.CharacterCfg.Game.Crypt.OpenChests, monsterFilter)
}
