package gun

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			Name:  "Musket",
			Power: 3,
		},
	}
}
