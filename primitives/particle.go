package primitives

import (
	"math"
)

const k = float64(1)

type Particle struct {
	Position Point
	Vector   Vector
	Mass     float64
}

func sanitiseZeroValues(v float64) float64 {
	if v == 0 {
		return 1
	}
	return v
}

func findDistance(x, y, z float64) float64 {
	return sanitiseZeroValues(math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2))
}

func (p *Particle) Tick(duration float64, otherParticles *[]Particle) {
	// p.Vector.X *= 0.9
	// p.Vector.Y *= 0.9
	// p.Vector.Z *= 0.9
	for _, otherParticle := range *otherParticles {
		xDiff, yDiff, zDiff := p.Position.X-otherParticle.Position.X,
			p.Position.Y-otherParticle.Position.Y,
			p.Position.Z-otherParticle.Position.Z
		combinedDifferences := sanitiseZeroValues(math.Abs(xDiff) + math.Abs(yDiff) + math.Abs(zDiff))

		force := k * otherParticle.Mass / findDistance(xDiff, yDiff, zDiff)

		if force < 1e4 {
			p.Vector.X += force * xDiff / combinedDifferences
			p.Vector.Y += force * yDiff / combinedDifferences
			p.Vector.Z += force * zDiff / combinedDifferences
		}

	}

	p.Position.X += duration * p.Vector.X
	p.Position.Y += duration * p.Vector.Y
	p.Position.Z += duration * p.Vector.Z
}
