package domain

import "time"

type Job struct {
	ID               string
	OutputBucketPath string
	Status           string
	Video            *Video
	Error            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

/*
Variável Video *Video
Ponteiro, apontando para dizer que todasa as vezes que for mexido na estrutura desse video,
estando dentro do job, vai estar mexdendo com o video
ex: método func exampleVideoPointer
*/
func exampleVideoPointer() {
	v := Video{}
	j := Job{}

	// sempre que alterar o j.Video, o v também será alterado e vice-versa.
	// estão apontando para o mesmo local da memória
	j.Video = &v
}
