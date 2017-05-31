package primitives

type Subcube struct {
	startX, endX, startY, endY, startZ, endZ float64
	Particles                                []Particle
	RelevantParticles                        []Particle // Particles not in the cube, but near it
}

func (s *Subcube) Tick(time float64) *[]Particle {
	newSubcubeParticlePositions := make([]Particle, len(s.Particles))

	// particle is copied by value, so can mutate
	for i, particle := range s.Particles {
		particle.Tick(time, &s.RelevantParticles)
		newSubcubeParticlePositions[i] = particle
	}

	return &newSubcubeParticlePositions
}
