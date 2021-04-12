package untils

func main() {

}

type Person struct {
	name string "lq"
	age int
}

func (p Person) SetName()string{
	return p.name
}