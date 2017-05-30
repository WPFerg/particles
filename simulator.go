package main

import (
	"log"

	"github.com/wpferg/particles/primitives"
)

func Simulate(cube *primitives.Cube, particles *[]primitives.Particle) {
	MAX_ITERATIONS := 1000

	log.Println("Starting simulation.")

	currentIteration := *particles

	for i := 0; i < MAX_ITERATIONS; i++ {
		nextIteration := make([]primitives.Particle, len(currentIteration))

		if i%50 == 0 {
			log.Println("Running iteration", i)
		}

		for j, particle := range currentIteration {
			particle.Tick(1e-4, &currentIteration)
			nextIteration[j] = particle
		}

		// (*cube).Collide(&nextIteration)
		currentIteration = nextIteration

		SaveFile(i, currentIteration)
	}
}
