package main
import (
	"fmt"
	"flag"
)

var name = flag.String("name", "World", "A name to say hello to.")

var russian bool
func init(){
	flag.BoolVar(&russian, "russian", false, "Use Russian language.")
	flag.BoolVar(&russian, "r", false, "Use Russian language.")	
}

func main() {
		flag.Parse() //  Парсинг флагов и установка соответствующих значений в переменных
		if russian == true {
			fmt.Printf("Привет %s!\n", *name)// ←  Доступ к name как к указателю
		} else {
			fmt.Printf("Hello %s!\n", *name)
	}
}

