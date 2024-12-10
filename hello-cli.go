package main
import (
	"fmt"
	"os"
	"github.com/urfave/cli"// ← Подключение пакета cli.go
)
func main() {
// Создание нового экземпляра приложения
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print hello world"
	app.Flags = []cli.Flag{
// Настройка флагов
		cli.StringFlag{
			Name: "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
	}
	app.Action = func(c *cli.Context) error {
		// Определение выполняемого действия
		name := c.GlobalString("name")
		fmt.Printf("Hello %s!\n", name)
		return nil
	}
	app.Run(os.Args)// ←  Запуск приложения
}

