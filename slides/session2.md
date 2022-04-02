# Go Study Group

## Sesión 2: Structs, métodos e interfaces

### 7 de abril, 2021

![w:398 h: 430 bg right](img/gomic.png)

---

# ¿Qué es un struct?

Un _struct_ es un tipo de datos compuesto de varios valores, que puede ser tratado como una sola unidad.

```go
package store

import "time"

// Record represents information about a record.
type Record struct {
	Name        string
	Artist      string
	ReleaseDate time.Time
}
```

---
# Uso de mayúsculas/minúsculas

Las palabras que comienzan con mayúscula son de acceso público (pueden ser accedidas desde fuera).
A su vez, las que comienzan con minúscula son de acceso privado (no pueden ser accedidas desde fuera).

```go
package store

type Customer struct {} // público

type stock struct {} // privado

type Record struct {
	Name string // público
	copies int  // privado
}
```

---
# Struct _nesting_ (anidamiento)

Se pueden anidar structs dentro de otros, como campos. Esto nos permite componer structs.

```go
package store

type Image struct {
	Url string
}

type Record struct {
	ArtCover Image
}

r := Record{
    Name:        "Animals",
    Artist:      "Pink Floyd",
    ArtCover:    Image{Url: "..."},
    ReleaseDate: date(1977, time.January, 23),
}
```

---
# Structs y funciones

Al definir un struct, introducimos un nuevo tipo de datos. Pueden ser parámetro o valor de retorno de funciones.

```go

func DaysSinceRelease(r Record) float64 {
    return time.Since(r.releaseDate).Hours() / 24
}

func ReleaseNewRecord(name, artist string, artCover Image) Record {
    return Record{
		Name: name,
		Artist: artist,
		ReleaseDate: time.Now(),
		ArtCover: artCover
	}
}
```

---
# Métodos

Go soporta programación orientada a objetos uniendo funciones (métodos) a structs. No soporta clases ni herencia.

```go
func (r Record) DaysSinceRelease() float64 {
    return time.Since(r.ReleaseDate).Hours() / 24
}
r = Record{...}
r.DaysSinceRelease()
```

---
# Recursos

- Go by example: structs https://gobyexample.com/structs
- Kata sobre Structs, métodos e interfaces https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces
- Katas + slides: https://github.com/isaporiti/go-study-group

---
# Consultas y sugerencias

- https://twitter.com/codurance_ES
- ignacio.saporiti@codurance.com
