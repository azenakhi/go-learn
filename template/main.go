import (
        "log"
        "os"
        "text/template"
)

type Config struct {
        Material string
        Count    uint
}

func main() {
        f, err := os.Create("out/config.xml")
        if err != nil {
                log.Println("create file: ", err)
                return
        }
        config := Config{"wool", 17}
        //tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
        tmpl, err := template.ParseFiles("tmpl/config.tmpl")

        if err != nil {
                panic(err)
        }
        err = tmpl.Execute(f, config)
        if err != nil {
                panic(err)
        }
        f.Close()
}
