package father

import "fmt"

func init() {
	fmt.Println("father package initialized!!")
}

type Father struct {
	name string
}

func (f Father) Data(name string) string {
	f.name = " Father : " + name
	return f.name
}
