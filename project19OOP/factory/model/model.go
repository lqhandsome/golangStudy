package model

func main() {

}

type student struct {
	Name string
	Score float64
}


func GetStuden(name string,score float64) *student {
		return &student{
			name,
			score,
		}
}