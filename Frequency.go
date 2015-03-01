package main

type Frequency struct {
	name  string
	hours int
}

func (f *Frequency) daily() Frequency {
	return Frequency{"daily", 24}
}
func (f *Frequency) weekly() Frequency {
	return Frequency{"weekly", 24 * 7}
}

// todo: fix
func (f *Frequency) monthly() Frequency {
	return Frequency{"monthly", 24 * 7 * 4}
}
