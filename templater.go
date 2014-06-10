package main
import "text/template"
import "os"
import "io/ioutil"

func equal(a string, b string)(c bool) {
    return bool (a == b);
}

func main() {

    bytes, _ := ioutil.ReadAll(os.Stdin)
    text := string(bytes)

    funcMap := template.FuncMap{
          "getenv": os.Getenv,
          "equal": equal,
          }

    tmpl, err := template.New("test").Funcs(funcMap).Parse(text)

    if err != nil { panic(err) }
    err = tmpl.Execute(os.Stdout, nil)
    if err != nil { panic(err) }
}
