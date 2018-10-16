import (
    "fmt"
    "io/ioutil
)


type Page struct{
    Title string
    Body []byte
}

func ( p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filenname, p.Body, 0600)
}

