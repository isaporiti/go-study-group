# Go Study Group

## Sesión 1: Introducción a Go

### 10 de Marzo, 2021

![w:398 h: 430 bg right](img/gomic.png)

---

# ¿Qué es Go? 🤔

Es un lenguaje de programación compilado. Un software llamado compilador traduce código fuente a código máquina (ejecutable).

```shell
go build main.go -o greeter_app
./greeter_app
# Hello, World!
```

--- 

# ¿Qué es Go? 🤔 (II)

Estáticamente tipado: el compilador comprueba que todos los tipos coinciden con las operaciones que el programa realiza.

```go
package main

func main(){
	var number int
	number = 3
	number += " little pigs" // type error
	...
}
```

--- 

# ¿Qué es Go? 🤔 (III)

Sintácticamente similar a C pero más seguro en gestión de memoria:

- Punteros, pero sin aritmética de punteros
- Garbage Collection

A veces llamado "el C del siglo XXI", otras veces como "C--". 🤷‍♂️

---

# ¿Qué es Go? 🤔 (IV)

Herramientas de desarrollo out-of-the-box:

- Testing (`go test`)
- Gestión de dependencias (`go get ...`)
- Formato de código (`go fmt`)

---
Supongamos que tenemos que comenzar un proyecto con Go, desde cero.

¿Por dónde empezarían?

---
# Paquete de Go: _testing_

Nos provee soporte para ejecutar tests automatizados de otros paquetes Go. Estos tests se ejecutan con el comando `go test`:

```shell
$ ls
greeter.go greeter_test.go

$ go test -v
=== RUN   TestGreet
--- PASS: TestGreet (0.00s)
PASS
ok  	greeter/greeter 0.111s
```

---
# Cómo usar _testing_

`go test` sólo reconoce como tests a _funciones_ que: 

- se encuentren en archivos con nombres que terminen en `_test.go`
- tengan la siguiente forma:

```go
func TestXxx(*testing.T)
```

donde `Xxx` no comienza con una letra minúscula.

---
# ¿Qué es *testing.T?

Parámetro que nos pasa Go en cada ejecución de nuestros tests. Nos da el control para señalizar cuándo fallan nuestros test.
Algunas de las funciones que provee, son:

- `Errorf`: marca al test como fallido con un mensaje que explica por qué falló. No detiene ejecución del test.
- `Fatalf`: igual que `Errorf` pero detiene ejecución del test.
- Para casos en donde un fallo hace que no tenga sentido continuar con más aserciones.

---
# Cómo usar _testing_ II

```shell
$ ls
greet.go greet_test.go
```

```go
package main

import "testing"

func TestGreet(t *testing.T)  {
		got := Greeter("Codurance", English)
		want := "Hello, Codurance!"

	if got != want {
			t.Errorf("want '%s', got '%s'", want, got)
	}
}
```

---
# Tema del día: Test-Driven Development (TDD)

- Proceso iterativo de desarrollo de software.
- Escribimos código solo para hacer pasar tests escritos anteriormente.
- Escribir soluciones simples revela patrones ➡️ al detectarlos podemos refactorizar.
- Refactor: modificación de código ya existente para mejorarlo.

![w:340 h:340 bg right](img/tdd-cicle.png)

---
# ¿Por qué TDD?

- Nos concentramos en unidades de código. Dividir y conquistar.
- Nos da confianza: aquello que desarrollamos está cubierto por los tests.
- Nos ayuda a detectar cambios que rompen otras funcionalidades ➡️ evitamos introducir bugs
- Tests bien escritos pueden describir nuestro código. Facilita onboarding.
