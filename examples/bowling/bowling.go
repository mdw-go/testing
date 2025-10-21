package bowling

type game struct {
	score  int
	throw  int
	throws [21]int
}

func (this *game) recordRoll(pins int) {
	this.throws[this.throw] = pins
	this.throw++
}

func (this *game) calculateScore() int {
	this.throw = 0
	this.score = 0
	for range 10 {
		this.score += this.scoreCurrentFrame()
		this.throw += this.advanceFrame()
	}
	return this.score
}
func (this *game) scoreCurrentFrame() int {
	if this.currentFrameIsStrike() {
		return this.scoreStrikeFrame()
	} else if this.currentFrameIsSpare() {
		return this.scoreSpareFrame()
	} else {
		return this.scoreRegularFrame()
	}
}
func (this *game) currentFrameIsStrike() bool {
	return this.pins(0) == 10
}
func (this *game) currentFrameIsSpare() bool {
	return this.frameScore() == 10
}
func (this *game) scoreStrikeFrame() int {
	return 10 + this.pins(1) + this.pins(2)
}
func (this *game) scoreSpareFrame() int {
	return 10 + this.pins(2)
}
func (this *game) scoreRegularFrame() int {
	return this.frameScore()
}
func (this *game) frameScore() int {
	return this.pins(0) + this.pins(1)
}
func (this *game) pins(offset int) int {
	return this.throws[this.throw+offset]
}
func (this *game) advanceFrame() int {
	if this.currentFrameIsStrike() {
		return 1
	} else {
		return 2
	}
}
