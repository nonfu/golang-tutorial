package animal

type Pet struct {
    name string
}

func NewPet(name string) Pet {
    return Pet{name: name}
}

func (p Pet) GetName() string  {
    return p.name
}
