package bowling

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestGameFixture(t *testing.T) {
	should.Run(&GameFixture{T: t})
}

type GameFixture struct {
	*testing.T
	game *game
}

func (this *GameFixture) Setup() {
	this.game = new(game)
}
func (this *GameFixture) assertScore(expected int) {
	this.Helper()
	should.So(this, this.game.calculateScore(), should.Equal, expected)
}
func (this *GameFixture) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		this.game.recordRoll(pins)
	}
}
func (this *GameFixture) rollSeveral(throws ...int) {
	for _, throw := range throws {
		this.game.recordRoll(throw)
	}
}
func (this *GameFixture) TestGutterGame() {
	this.rollMany(20, 0)
	this.assertScore(0)
}
func (this *GameFixture) TestAllOnes() {
	this.rollMany(20, 1)
	this.assertScore(20)
}
func (this *GameFixture) TestSpare() {
	this.rollSeveral(5, 5, 3)
	this.assertScore(16)
}
func (this *GameFixture) TestStrike() {
	this.rollSeveral(10, 3, 4)
	this.assertScore(24)
}
func (this *GameFixture) TestPerfectGame() {
	this.rollMany(12, 10)
	this.assertScore(300)
}
