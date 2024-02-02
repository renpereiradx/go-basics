package main

import (
	pk "curso/hola/mypackage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Funciones
func hello5(saludo string) {
	println(saludo)
}

//	func suma(a int8, b int8) int8 {
//		return a + b
//	}
//
// funcion con retorno multiple
func suma(a int8, b int8) (int8, string) {
	return a + b, "La suma es: "
}

func resta(name string, a, b int8) (string, int8) {
	message := "Hola " + name + " la resta es: "
	return message, a - b
}

// funcion que rotorna tres valores
func retornaTresValores() (int8, string, bool) {
	return 1, "Hola", true
}

// Structs
type user struct {
	nombre string
	edad   int8
}

func main() {
	println("Hello, World! This is my first Go program.")
	// Declaracion de variables y constantes
	// Variables
	var x int = 5
	var y float32 = 2.2
	z := 3
	println(x, y, z)
	// Constantes
	const pi float64 = 3.1415
	const pi2 = 3.14
	pi3 := 3.14159
	println(pi, pi2, pi3)
	// Area de un cuadrado, formula
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	println("El area del cuadrado es: ", areaCuadrado)
	// Valores zero
	var a int
	var b float64
	var c string
	var d bool
	println(a, b, c, d)

	var m int = 50
	var n int = 10

	// suma
	result := m + n
	println("La suma es: ", result)
	// resta
	result = m - n
	println("La resta es: ", result)
	// multiplicacion
	result = m * n
	println("La multiplicacion es: ", result)
	// division
	result = m / n
	println("La division es: ", result)
	// modulo
	result = m % n
	println("El modulo es: ", result)
	// incremento
	m++
	println("El incremento es: ", m)
	// decremento
	m--
	println("El decremento es: ", m)
	// area de un rectangulo
	baseRectangulo := 20
	alturaRectangulo := 10
	areaRectangulo := baseRectangulo * alturaRectangulo
	println("El area del rectangulo es: ", areaRectangulo)
	// area de un triangulo
	baseTriangulo := 20
	alturaTriangulo := 10
	areaTriangulo := (baseTriangulo * alturaTriangulo) / 2
	println("El area del triangulo es: ", areaTriangulo)
	// area de un circulo
	const pi4 float64 = 3.1415
	var radioCirculo float64 = 10
	areaCirculo := pi4 * (radioCirculo * radioCirculo)
	println("El area del circulo es: ", areaCirculo)
	// Tipos de datos primitivos
	// Enteros	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// byte	// alias para uint8
	// rune	// alias para int32
	// Flotantes	// float32, float64
	// Strings	// string
	// Booleanos	// bool
	// Impresion en consola
	// Println
	fmt.Printf("El valor de pi es: %f\n", pi)
	fmt.Printf("El valor de m es: %d y el valor de n es: %d\n", m, n)
	fmt.Printf("El valor de x es %v y el valor de y es %v\n", x, y)
	// Funciones
	hello5("Hola mundo")
	// suma
	resultSuma, message := suma(2, 3)
	println(message, resultSuma)
	// resta
	messageResta, resultResta := resta("Juan", 10, 5)
	println(messageResta, resultResta)
	// retornar tres valores
	_, valor2, valor3 := retornaTresValores()
	println(valor2, valor3)
	// Ciclos
	// For
	for i := 10; i >= 0; i-- {
		println(i)
	}
	// For while
	counter := 0
	for counter < 10 {
		println(counter)
		counter++
	}
	// For forever
	// counterForever := 0
	// for {
	// 	println(counterForever)
	// 	counterForever++
	// }
	// For range
	numbers := []string{"cero", "uno", "dos", "tres"}
	for i, v := range numbers {
		println(i, v)
	}
	// If
	value := 1
	if value == 1 {
		println("Es uno")
	} else {
		println("No es uno")
	}

	i, err := strconv.Atoi("45")
	if err != nil {
		println(err)
	} else {
		println(i)
	}
	// Switch
	modulo := 4 % 2
	switch modulo {
	case 0:
		println("Es par")
	default:
		println("Es impar")
	}
	// Switch sin condicion
	valueSwitch := 50
	switch {
	case valueSwitch > 100:
		println("Es mayor a 100")
	case valueSwitch < 0:
		println("Es menor a 0")
	default:
		println("No condicion")
	}
	// defer, continue y break
	defer println("Se imprimira al final: DEFER")
	// continue
	for i := 0; i < 10; i++ {
		if i == 2 {
			println("Es 2")
			continue
		}
		if i == 8 {
			println("Break")
			break
		}
		println(i)
	}
	// Arreglos
	var arreglo [5]int
	arreglo[0] = 1
	arreglo[1] = 2
	fmt.Println(arreglo)
	// Arreglos con valores
	arreglo2 := [3]int{1, 2, 3}
	fmt.Println(arreglo2)
	// Slices	// Similar a los arreglos pero sin definir el tamaño
	slices1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slices1, len(slices1), cap(slices1))
	// obtener una serie de elementos del slice
	slices2 := slices1[0:3] // [1, 2, 3] 		// desde el indice 0 hasta el indice 3 (excluyente)
	fmt.Println(slices2, len(slices2), cap(slices2))
	// obtener una serie de elementos del slice
	slices3 := slices1[2:] // [3, 4, 5, 6]	// desde el indice 2 hasta el final
	fmt.Println(slices3, len(slices3), cap(slices3))
	// obtener una serie de elementos del slice sin incluir el ultimo elemento
	slices4 := slices1[:3] // [1, 2, 3]		// desde el inicio hasta el indice 3
	fmt.Println(slices4, len(slices4), cap(slices4))
	// obtener una serie de elementos del slice sin incluir el primer elemento
	slices4[2] = 100
	fmt.Println(slices1)
	// append
	newSlice := append(slices1, 7)
	println(newSlice, len(newSlice), cap(newSlice))
	println(newSlice[0], " ", slices1)
	// Recorrido de Slice con range
	for i, v := range newSlice {
		println(i, v)
	}
	// Recorrido de un Array con range
	for i, v := range arreglo {
		println(i, v)
	}
	// Maps
	paises := make(map[string]string)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	detallesPersona := map[string]int{
		"Edad":   30,
		"Altura": 180,
		"Peso":   80,
	}
	fmt.Println(detallesPersona)
	valorPais, ok := paises["Mexico"]
	fmt.Println(valorPais, ok)

	for i, v := range detallesPersona {
		println(i, v)
	}
	// Instanciar un struct
	usuario := user{nombre: "Juan", edad: 30}
	fmt.Println(usuario)
	// Instanciar un struct sin valores
	usuario2 := user{}
	usuario2.nombre = "Pedro"
	fmt.Println(usuario2)
	// Instanciar desde otro paquete
	var Car pk.Car
	Car.Brand = "Ferrari"
	Car.Year = 2020
	fmt.Println(Car)
	// Punteros
	var valuePointer *int = &x
	fmt.Println(valuePointer, *valuePointer)
	// Otra forma de declarar punteros
	valuePointer2 := &y
	fmt.Println(valuePointer2)
	// Punteros receptores
	var machine pk.PC
	machine.Ram = 16
	machine.Disk = 200
	machine.Brand = "MSI"
	machine.MultiplicarRam()
	fmt.Println(machine)
	// Stringer
	fmt.Println(machine.String())
	// Interfaces
	cuadrado := pk.Cuadrado{Base: 10}
	triangulo := pk.Triangulo{Base: 10, Altura: 30}
	var interfaz pk.Figura
	interfaz = cuadrado
	fmt.Println("Interfaz", " ", interfaz.Area())
	fmt.Println("Interfaz", " ", pk.Calcular(interfaz))
	fmt.Println(pk.Calcular(cuadrado))
	fmt.Println(pk.Calcular(triangulo))
	// Goroutines
	// var wg sync.WaitGroup
	// fmt.Println("Hello")
	// wg.Add(1)
	// go pk.Say("World", &wg)
	// // wg.Wait()
	// // wg.Add(1)
	// wg.Add(1)
	// go func(text string, wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	fmt.Println(text)
	// }("Adios", &wg)
	// go func(text string) {
	// 	// defer wg.Done()
	// 	fmt.Println(text)
	// }("Adios2")
	// wg.Wait()
	// fmt.Println("End")
	// Channels
	channel := make(chan string, 1)
	channel2 := make(chan int, 1)
	fmt.Println("Hello")
	go pk.SaySomething("World", channel)
	channelResult := <-channel
	fmt.Println(channelResult)
	channel2 <- 1
	fmt.Println(<-channel2)
	// Select
	channel3 := make(chan string, 2)
	channel4 := make(chan string, 2)
	channel3 <- "Hello"
	channel3 <- "World"
	channel4 <- "Adios"
	channel4 <- "Mundo"
	close(channel3)
	close(channel4)
	// Necesitaras suspender el programa por el for infinito
	// for {
	//  No se puede usar el default en un for infinito
	// 	time.Sleep(1 * time.Second)
	// 	select {
	// 	case message := <-channel3:
	// 		fmt.Println(message)
	// 	case message2 := <-channel4:
	// 		fmt.Println(message2)
	// 	default:
	// 		fmt.Println("Default")
	// 	}
	// }
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
