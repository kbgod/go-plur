# go-plur
A simple Go package to pluralize words.
> This package using Generic types, so it requires Go 1.18 or later.

## Installation
```bash
go get github.com/kbgod/go-plur
```

## Usage
```go
func main() {
	one := "яблоко"
	two := "яблока"
	many := "яблук"

	v := 112
	fmt.Println(plur.Text(v, one, two, many)) // яблук
	fmt.Println(plur.NumberText(v, one, two, many)) // 112 яблук
}
```