package animal

type Dog struct {
    animal *Animal
    pet Pet
}

func NewDog(animal *Animal, pet Pet) Dog {
    return Dog{animal: animal, pet: pet}
}

func (d Dog) FavorFood() string {
    return d.animal.FavorFood() + "骨头"
}

func (d Dog) Call() string {
    return d.animal.Call() + "汪汪汪"
}

func (d Dog) GetName() string {
    return d.pet.GetName()
}