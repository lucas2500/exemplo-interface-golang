package main

import "fmt"

type Autor struct {
	Nome  string
	Email string
}

func (a Autor) String() string {
	return fmt.Sprint("Nome do autor/Email do autor: ", a.Nome, "/", a.Email)
}

type Livros struct {
	Nome      string
	Categoria string
	Autor     string
}

type Livro struct {
	Livros
	TotalCapitulos int
}

func (l Livro) String() string {
	return fmt.Sprint("Nome do livro/Categoria/Autor/Total de capitulos: ", l.Nome, "/", l.Categoria, "/", l.Autor, "/", l.TotalCapitulos)
}

type HQ struct {
	Livros
	Desenhista string
}

func (h HQ) String() string {
	return fmt.Sprint("Nome da HQ/Categoria/Autor/Desenhista: ", h.Nome, "/", h.Categoria, "/", h.Autor, "/", h.Desenhista)
}

type Stringer interface {
	String() string
}

func main() {

	a := Autor{"Lucas Barbosa", "lucasrafael654@gmail.com"}

	Printar(a)

	b := Livro{Livros: Livros{Nome: "Programando em Go", Categoria: "Programação", Autor: "Lucas Barbosa"}, TotalCapitulos: 10}

	Printar(b)

	c := HQ{Livros: Livros{Nome: "Guerra Civil Marvel Essentials", Categoria: "Ação e aventura", Autor: "Mark Millar"}, Desenhista: "Steve McNiven"}

	Printar(c)
}

func Printar(s Stringer) {
	fmt.Println(s.String())
}
