package depedenciestest

//depedecies injectin, menginject sesuatu ke dalam class

type Speeder interface {
	MaxSpeed() int
}

func NewCar(speeder Speeder) *Car {
	return &Car{
		Speeder: speeder,
	}
}

type Car struct {
	Speeder Speeder
}

func (c Car) Speed() int {
	defaultSpeed := 80

	if c.Speeder.MaxSpeed() < 10 {
		return 20
	}

	if defaultSpeed > c.Speeder.MaxSpeed() {
		return c.Speeder.MaxSpeed()
	}
	return defaultSpeed
}

//mesin bawaan pabrik
type DefaultEngine struct{}

func (e DefaultEngine) MaxSpeed() int {
	return 50
}

//turbo
type TurboEngine struct{}

func (e TurboEngine) MaxSpeed() int {
	return 100
}
