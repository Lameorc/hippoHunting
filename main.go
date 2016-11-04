package main

import (
	tl "github.com/JoelOtter/termloop"
)

//ScoreCard ...
type ScoreCard struct {
	entity *tl.Entity
	value  int
}

func (scoreCard *ScoreCard) Draw(screen *tl.Screen) {
	scoreCard.entity.SetCell(0, 0, &tl.Cell{Ch: rune(scoreCard.value)})
	scoreCard.entity.Draw(screen)
}

func (scoreCard *ScoreCard) Tick(event tl.Event) {
}

type Projectile struct {
	direction  Direction
	entity     *tl.Entity
	prevX      int
	prev      int
	originator *tl.Entity
	ticker     byte
}

func (projectile *Projectile) Draw(screen *tl.Screen) {
	projectile.entity.Draw(screen)
}
func (projectile *Projectile) Tick(event tl.Event) {
	projectile.prevX, projectile.prevY = projectile.entity.Position()
	switch projectile.direction {
	case Right:
		projectile.entity.SetPosition(projectile.prevX+1, projectile.prevY)
	case Left:
		projectile.entity.SetPosition(projectile.prevX-1, projectile.prevY)
	case Up:
		projectile.entity.SetPosition(projectile.prevX, projectile.prevY-1)
	case Down:
		projectile.entity.SetPosition(projectile.prevX, projectile.prevY+1)
	}
}

type Direction byte

const (
	Up Direction = iota
	Right
	Down
	Left
)

//Player ...
type Player struct {
	entity       *tl.Entity
	level        *tl.BaseLevel
	prevX        int
	prevY        int
	aimDirection Direction
}

//Draw ...
func (player *Player) Draw(screen *tl.Screen) {

	player.entity.Draw(screen)
}

//Tick ...
func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.entity.Position()
		switch event.Key { // if so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.entity.SetPosition(player.prevX+1, player.prevY)
			player.aimDirection = Right
		case tl.KeyArrowLeft:
			player.entity.SetPosition(player.prevX-1, player.prevY)
			player.aimDirection = Left
		case tl.KeyArrowUp:
			player.entity.SetPosition(player.prevX, player.prevY-1)
			player.aimDirection = Up
		case tl.KeyArrowDown:
			player.entity.SetPosition(player.prevX, player.prevY+1)
			player.aimDirection = Down
		case tl.KeySpace:
			playerX, playerY := player.entity.Position()
			var bullet = new(Projectile)
			bullet.entity = tl.NewEntity(playerX, playerY, 1, 1)
			bullet.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Ch: '.'})
			bullet.originator = player.entity
			level := player.level
			level.AddEntity(bullet)
		}
	}
}
func main() {
	game := tl.NewGame()
	game.SetDebugOn(true)
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	level.AddEntity(
		tl.NewText(
			25,
			0,
			"giant cocroach hunting hippos",
			tl.ColorWhite,
			tl.ColorBlack))

	player := Player{
		entity: tl.NewEntity(10, 10, 1, 1),
	}
	playerScore := ScoreCard{
		entity: tl.NewEntity(0, 1, 1, 1),
	}

	// Set the Character at position (0, 0) on the entity.
	player.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: 'P'})
	playerScore.entity.SetCell(0, 0, &tl.Cell{
		Fg: tl.ColorWhite,
		Ch: '0',
	})
	level.AddEntity(&player)
	level.AddEntity(&playerScore)
	game.Screen().SetLevel(level)

	game.Start()
}
