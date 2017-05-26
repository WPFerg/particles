package primitives

import (
	"math"
)

const k = float64(9e9)

type Particle struct {
	Position Point
	Vector   Vector
}

func findDistance(x, y, z float64) float64 {
	return math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2)
}

func (p *Particle) Tick(duration float64, otherParticles *[]Particle) {
	for _, otherParticle := range *otherParticles {
		xDiff, yDiff, zDiff := math.Abs(p.Position.X-otherParticle.Position.X),
			math.Abs(p.Position.Y-otherParticle.Position.Y),
			math.Abs(p.Position.Z-otherParticle.Position.Z)
		combinedDifferences := xDiff + yDiff + zDiff

		force := k / findDistance(xDiff, yDiff, zDiff)

		p.Vector.X += duration * force * xDiff / combinedDifferences
		p.Vector.Y += duration * force * yDiff / combinedDifferences
		p.Vector.Z += duration * force * zDiff / combinedDifferences
	}
}
