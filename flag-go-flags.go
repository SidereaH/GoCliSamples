package main 
import (
	"fmt"
	flags "github.com/jessevdk/go-flags" // ← Импорт пакета go-flags с присвоением псевдонима flags
)
var opts struct {
	Name string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Spanish bool `short:"s" long:"spanish" description:"Use Spanish Language"`
}
// Структура с определениями флагов
func main() {
 	flags.Parse(&opts)// ←  Извлечение значений флагов в структуру
	if opts.Spanish == true {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
