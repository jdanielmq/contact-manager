package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//revisdaremos pasar un string a numero, capturando el error de la funcion, 1 = exitoso y 2 = error
	//numero 1
	str := "1213"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("el numero parseado es: ", num)
	//numero 2
	str = "1234" // para revisar o probar el archivo, se debe colocar en numerico la variable para que puede seguir la fecuencia
	n, e := strconv.Atoi(str)
	if e != nil {
		fmt.Println("Error: ", e)
		fmt.Println("###############################")
		return
	}
	fmt.Println("el numero pasrseado es: ", n)
	fmt.Println("###############################")

	// vamos a revisar uina division / 0 , validamos si el divisor es 0 creamos un error dentro de la funcion
	num1, err1 := divide(4, 1)
	if err1 != nil {
		fmt.Println("Error de la division: ", err1)
		fmt.Println("###############################")
		return
	}
	fmt.Println("la division se realizo correctamente: ", num1)
	fmt.Println("###############################")

	/* vamos revisar como se utiliza la funcionalidad de DEFER, se ejecuta siempre al final
	   pero si tenemos mas un defer, se apilan y se ejecutan de modo fifo
	*/
	file, err2 := os.Create("hola.txt")
	if err2 != nil {
		fmt.Println("Error: ", err2)
		fmt.Println("###############################")
		return
	}
	defer file.Close()

	_, err2 = file.Write([]byte("Hola, soy juan daniel"))
	if err2 != nil {
		fmt.Println(err2)
		fmt.Println("###############################")
		return
	}

	//vamos revisar como se utiliza el panic y recoved
	dividir(200, 2)
	dividir(300, 3)
	dividir(100, 0)
	dividir(30, 5)
	fmt.Println("###############################")

	//vamos a revisar la session de registros de errores LOG
	log.Print("estamos imprimiendo con print")
	log.Println("estamos imprimiendo con println")
	//log.Fatal("soy un registro de errores") ** esta linea debemos porque aca se corta la secuencia
	log.Print("puedes verme!!!")
	//log.Panic("oye, soy un registro de errores.!!") ** esta linea se comenta porque corta la secuencia y muentra donde esta el error.
	log.Println("estamos imprimiendo despues del log.panic()")
	log.SetPrefix("main()")
	log.Printf("estamos colocando un prefijo a la linea, se colocan en la clase o archivo")
	fmt.Println("###############################")

	//ahora vamos registrar los mensajes en un archivo info.log
	file, err3 := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err3 != nil {
		log.Fatal(err3)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Oye, soy un log.")
	fmt.Println("###############################")

}

func divide(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("No se puede dividir por 0")
	}

	return (num1 / num2), nil
}

func dividir(num1, num2 int) {
	/*funcion anomima para compturar todo los panicos de la funcion y no interrumpe la ejecucion
	y si no existe la funcion de anonima, se ejecuta el panic go se interrumpe la ejecucion */
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(num2)
	fmt.Println(num1 / num2)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por cero.")
	}
}
