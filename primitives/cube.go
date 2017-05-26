package primitives

import (
	"math"
)

type Cube struct {
	PointA, PointB Point
}

func flip(collisionAxisA, collisionAxisB, particlePosition, particleVelocity float64) (float64, float64) {
	max := math.Max(collisionAxisA, collisionAxisB)
	min := math.Min(collisionAxisA, collisionAxisB)
	var breachingAxis float64

	if particlePosition > max {
		breachingAxis = max
	}

	if particlePosition < min {
		breachingAxis = min
	}

	return breachingAxis - particlePosition, -particleVelocity
}

func collides(collisionAxisA, collisionAxisB, particlePosition float64) bool {
	max := math.Max(collisionAxisA, collisionAxisB)
	min := math.Min(collisionAxisA, collisionAxisB)
	return particlePosition < min || particlePosition > max
}

func (cube *Cube) runCollider(particles *[]Particle) *[]Particle {
	for i, particle := range *particles {
		if collides(cube.PointA.X, cube.PointB.X, particle.Position.X) {
			particle.Position.X, particle.Vector.X = flip(cube.PointA.X, cube.PointB.X, particle.Position.X, particle.Vector.X)
		}

		if collides(cube.PointA.Y, cube.PointB.Y, particle.Position.Y) {
			particle.Position.Y, particle.Vector.Y = flip(cube.PointA.Y, cube.PointB.Y, particle.Position.Y, particle.Vector.Y)
		}

		if collides(cube.PointA.Z, cube.PointB.Z, particle.Position.Z) {
			particle.Position.Z, particle.Vector.Z = flip(cube.PointA.Z, cube.PointB.Z, particle.Position.Z, particle.Vector.Z)
		}

		(*particles)[i] = particle
	}

	return particles
}

func (cube *Cube) Collide(particles []Particle) []Particle {
	return *(cube.runCollider(&particles))
}
