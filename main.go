package main

import (
	"log"
	"math/rand"

	"github.com/wpferg/particles/primitives"
)

func createParticles(count int) []primitives.Particle {
	particles := make([]primitives.Particle, count)

	for i := 0; i < count; i++ {
		particles[i] = primitives.Particle{
			Position: primitives.Point{rand.Float64(), rand.Float64(), rand.Float64()},
			Vector:   primitives.Vector{float64(0), float64(0), float64(0)},
			Mass:     0.1,
		}
	}

	particles[0].Mass = 10

	return particles
}

func main() {
	log.Println("Particle Simulationator")
	log.Println("Creating a cube")

	cube := primitives.Cube{
		PointA: primitives.Point{0, 0, 0},
		PointB: primitives.Point{1, 1, 1},
	}

	log.Println("Creating subcube partitioner")
	subcubes := primitives.GenerateSubcubes(0.2)

	log.Println("Creating particles.")
	particles := createParticles(500)

	log.Println("Partitioning particles into subcubes")
	subcubes.UpdateParticlePositions(&particles)

	Simulate(&cube, &subcubes, &particles)

	log.Println("Done")
}
