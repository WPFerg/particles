package main

import (
	"log"
	"math"

	"github.com/wpferg/particles/primitives"
)

func getComparisonParticles(subcubes *primitives.Subcubes, x, y, z int) *[]primitives.Particle {
	var particles []primitives.Particle
	CUBE_RANGE := 3
	CUBES_PER_AXIS := float64(len(subcubes.Cubes))

	// +1 for end indices to make them inclusive when slicing
	startX := int(math.Max(0.0, math.Min(float64(x-CUBE_RANGE), CUBES_PER_AXIS)))
	endX := int(math.Max(0.0, math.Min(float64(x+CUBE_RANGE+1), CUBES_PER_AXIS)))
	startY := int(math.Max(0.0, math.Min(float64(y-CUBE_RANGE), CUBES_PER_AXIS)))
	endY := int(math.Max(0.0, math.Min(float64(y+CUBE_RANGE+1), CUBES_PER_AXIS)))
	startZ := int(math.Max(0.0, math.Min(float64(z-CUBE_RANGE+1), CUBES_PER_AXIS)))
	endZ := int(math.Max(0.0, math.Min(float64(z+CUBE_RANGE), CUBES_PER_AXIS)))

	for xIndex := range subcubes.Cubes[startX:endX] {
		yZSlice := subcubes.Cubes[startX+xIndex]

		for yIndex := range yZSlice[startY:endY] {
			zSlice := yZSlice[startY+yIndex]

			for _, subcube := range zSlice[startZ:endZ] {
				particles = append(particles, subcube.Particles...)
			}
		}
	}

	return &particles
}

func simulateSubcube(subcubes *primitives.Subcubes, x, y, z int) []primitives.Particle {
	subcube := subcubes.Cubes[x][y][z]
	newPositions := make([]primitives.Particle, len(subcube.Particles))
	comparisonParticles := getComparisonParticles(subcubes, x, y, z)

	// particle is copied by value, so can mutate
	for i, particle := range subcube.Particles {
		particle.Tick(1e-6, comparisonParticles)
		newPositions[i] = particle
	}
	return newPositions
}

func simulateSubcubes(subcubes *primitives.Subcubes) []primitives.Particle {
	newPositions := make([]primitives.Particle, 0)

	for x := range subcubes.Cubes {
		for y := range subcubes.Cubes[x] {
			for z := range subcubes.Cubes[x][y] {
				newPositions = append(newPositions, simulateSubcube(subcubes, x, y, z)...)
			}
		}
	}

	return newPositions
}

func Simulate(cube *primitives.Cube, subcubes *primitives.Subcubes, particles *[]primitives.Particle) {
	MAX_ITERATIONS := 3000

	log.Println("Starting simulation.")

	currentIteration := *particles

	for i := 0; i < MAX_ITERATIONS; i++ {
		var nextIteration []primitives.Particle

		if i%50 == 0 {
			log.Println("Running iteration", i)
		}

		nextIteration = simulateSubcubes(subcubes)

		(*cube).Collide(&nextIteration)
		currentIteration = nextIteration

		SaveFile(i, currentIteration)

		subcubes.UpdateParticlePositions(&nextIteration)
	}
}
