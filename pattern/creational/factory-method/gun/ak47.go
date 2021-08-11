package gun

type AK47 struct {
	Gun
}

func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			Name:  "Kalashnikov",
			Power: 4,
		},
	}
}
