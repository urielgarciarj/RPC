package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1) Registrar un alumno, materia, calificación")
		fmt.Println("2) Promedio por alumno")
		fmt.Println("3) Promedio por materia")
		fmt.Println("4) Promedio general de alumnos")
		fmt.Println("0) Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:

			//m := make(map[string]string)
			var name string
			var assignature string
			var grade string

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Nombre: ")
			name, _ = reader.ReadString('\n')
			//fmt.Scanln(&name)

			reader = bufio.NewReader(os.Stdin)
			fmt.Print("Materia: ")
			assignature, _ = reader.ReadString('\n')

			fmt.Print("Calificación:")
			fmt.Scanln(&grade)

			//m[subject] = grade
			slice := []string{name, assignature, grade}
			var result string
			err = c.Call("Server.Register", slice, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.Register =", result)
			}

		case 2:
			var name string

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Introduzca nombre de alumno: ")
			name, _ = reader.ReadString('\n')

			var result string
			err = c.Call("Server.StudentGrade", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.StudentGrade", name, "=", result)
			}
		case 3:
			var op string
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("¿De que materia quiere el promedio? ")
			op, _ = reader.ReadString('\n')
			var result string
			err = c.Call("Server.Assignature", op, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.Assignature", op, "=", result)
			}
		case 4:
			var b bool
			b = true
			var result string
			err = c.Call("Server.EveryoneGrade", b, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.EveryoneGrade", b, "=", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
