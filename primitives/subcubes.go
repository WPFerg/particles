package primitives

import (
	"math"
)

func clamp(value float64) float64 {
	return math.Max(0, math.Min(value, 1))
}

func getIntendedSubcube(particle *Particle, subcubeAxisSize float64) (int32, int32, int32) {
	return int32(math.Floor(particle.Position.X / subcubeAxisSize)),
		int32(math.Floor(particle.Position.Y / subcubeAxisSize)),
		int32(math.Floor(particle.Position.Z / subcubeAxisSize))
}

type Subcubes struct {
	Cubes           [][][]Subcube
	Particles       []Particle
	subcubeAxisSize float64
}

func (s *Subcubes) clearCubes() {
	for x := range s.Cubes {
		for y := range s.Cubes[x] {
			for z := range s.Cubes[x][y] {
				s.Cubes[x][y][z].Particles = make([]Particle, 0)
			}
		}
	}
}

func (s *Subcubes) updateNeighbouringParticles() {
	for x := range s.Cubes {
		for y := range s.Cubes[x] {
			for z := range s.Cubes[x][y] {
				s.Cubes[x][y][z].RelevantParticles = *s.getNeighbouringParticles(x, y, z)
			}
		}
	}
}

func (s *Subcubes) getNeighbouringParticles(x, y, z int) *[]Particle {
	var particles []Particle
	CUBE_RANGE := 3
	CUBES_PER_AXIS := float64(len(s.Cubes))

	// +1 for end indices to make them inclusive when slicing
	startX := int(math.Max(0.0, math.Min(float64(x-CUBE_RANGE), CUBES_PER_AXIS)))
	endX := int(math.Max(0.0, math.Min(float64(x+CUBE_RANGE+1), CUBES_PER_AXIS)))
	startY := int(math.Max(0.0, math.Min(float64(y-CUBE_RANGE), CUBES_PER_AXIS)))
	endY := int(math.Max(0.0, math.Min(float64(y+CUBE_RANGE+1), CUBES_PER_AXIS)))
	startZ := int(math.Max(0.0, math.Min(float64(z-CUBE_RANGE), CUBES_PER_AXIS)))
	endZ := int(math.Max(0.0, math.Min(float64(z+CUBE_RANGE+1), CUBES_PER_AXIS)))

	for xIndex := range s.Cubes[startX:endX] {
		yZSlice := s.Cubes[startX+xIndex]

		for yIndex := range yZSlice[startY:endY] {
			zSlice := yZSlice[startY+yIndex]

			for _, subcube := range zSlice[startZ:endZ] {
				particles = append(particles, subcube.Particles...)
			}
		}
	}

	return &particles
}

func (s *Subcubes) UpdateParticlePositions(particles *[]Particle) {
	s.clearCubes()
	s.Particles = *particles

	for i := range s.Particles {
		x, y, z := getIntendedSubcube(&s.Particles[i], s.subcubeAxisSize)
		intendedSubcube := &s.Cubes[x][y][z]
		intendedSubcube.Particles = append(intendedSubcube.Particles, s.Particles[i])
	}

	s.updateNeighbouringParticles()
}

func GenerateSubcubes(subcubeWidth float64) Subcubes {
	numberOfSubcubes := int64(math.Ceil(1 / subcubeWidth))
	xSubCubes := make([][][]Subcube, numberOfSubcubes)

	for xIndex := 0; xIndex < len(xSubCubes); xIndex++ {
		ySubCubes := make([][]Subcube, numberOfSubcubes)
		startX := clamp(float64(xIndex) * subcubeWidth)
		endX := clamp(float64(xIndex+1) * subcubeWidth)

		for yIndex := 0; yIndex < len(ySubCubes); yIndex++ {
			zSubcubes := make([]Subcube, numberOfSubcubes)
			startY := clamp(float64(yIndex) * subcubeWidth)
			endY := clamp(float64(yIndex+1) * subcubeWidth)

			for subcubeIndex := range zSubcubes {
				zSubcubes[subcubeIndex] = Subcube{
					startX: startX,
					endX:   endX,
					startY: startY,
					endY:   endY,
					startZ: clamp(float64(subcubeIndex) * subcubeWidth),
					endZ:   clamp(float64(subcubeIndex+1) * subcubeWidth),
				}
			}

			ySubCubes[yIndex] = zSubcubes
		}
		xSubCubes[xIndex] = ySubCubes
	}

	return Subcubes{
		Cubes:           xSubCubes,
		subcubeAxisSize: subcubeWidth,
	}
}
