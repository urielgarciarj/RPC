package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"strings"
)

type Server struct{}

var science = map[string]int{}
var math = map[string]int{}
var history = map[string]int{}
var students = map[string]int{}
var grade int
var students2 []string
var dataStructure = "Estructura de Datos"

func (this *Server) Register(s []string, reply *string) error {
	verify := strings.Contains(s[1], "Estructura de Datos")
	if verify == true {
		_, isvalpresent2 := science[s[0]]
		if isvalpresent2 == false {
			grade, _ = strconv.Atoi(s[2])
			fmt.Println(grade)
			science[s[0]] = grade
			students[s[0]], _ = strconv.Atoi(s[2])
			fmt.Println(science)
			*reply = "Alumno registrado!"
		} else {
			*reply = "Error, este estudiante ya tiene calificación en esta materia"
		}
	}

	verify = strings.Contains(s[1], "Algoritmia")
	if verify == true {
		_, isvalpresent2 := math[s[0]]
		if isvalpresent2 == false {
			grade, _ = strconv.Atoi(s[2])
			math[s[0]] = grade
			students[s[0]], _ = strconv.Atoi(s[2])
			fmt.Println(math)
			*reply = "Alumno registrado!"
		} else {
			*reply = "Error, este estudiante ya tiene calificación en esta materia"
		}

	}

	verify = strings.Contains(s[1], "Sistemas Distribuidos")
	if verify == true {
		_, isvalpresent2 := history[s[0]]
		if isvalpresent2 == false {
			grade, _ = strconv.Atoi(s[2])
			history[s[0]] = grade
			students[s[0]], _ = strconv.Atoi(s[2])
			fmt.Println(history)
			*reply = "Alumno registrado!"
		} else {
			*reply = "Error, este estudiante ya tiene calificación en esta materia"
		}
	}

	return nil
}

func (this *Server) StudentGrade(name string, reply *string) error {
	*reply = "Hello "
	var aux float64
	var average float64
	var i float64
	for k, v := range science {
		fmt.Println("key: ", k, "value: ", v)
		if k == name {
			aux = aux + float64(v)
			i++
		}
		//aux = aux + float64(v)
	}

	for k, v := range math {
		fmt.Println("key: ", k, "value: ", v)
		if k == name {
			aux = aux + float64(v)
			i++
		}
		//aux = aux + float64(v)
	}

	for k, v := range history {
		fmt.Println("key: ", k, "value: ", v)
		if k == name {
			aux = aux + float64(v)
			i++
		}
		//aux = aux + float64(v)
	}
	average = aux / i               //int to float
	s := fmt.Sprintf("%f", average) // float to string
	*reply = "Promedio de la materia: " + s
	return nil
}

func (this *Server) Assignature(assignature string, reply *string) error {
	var aux float64
	var average float64
	verify := strings.Contains(assignature, "Estructura de Datos")
	if verify == true {
		for k, v := range science {
			fmt.Println("key: ", k, "value: ", v)
			aux = aux + float64(v)
		}
		fmt.Println(aux)
		average = aux / float64(len(science)) //int to float
		s := fmt.Sprintf("%f", average)       // float to string
		*reply = "Promedio de la materia: " + s
	}

	verify = strings.Contains(assignature, "Algoritmia")
	if verify == true {
		for k, v := range math {
			fmt.Println("key: ", k, "value: ", v)
			aux = aux + float64(v)
		}
		fmt.Println(aux)
		average = aux / float64(len(math)) //int to float
		s := fmt.Sprintf("%f", average)    // float to string
		*reply = "Promedio de la materia: " + s
	}

	verify = strings.Contains(assignature, "Sistemas Distribuidos")
	if verify == true {
		for k, v := range history {
			fmt.Println("key: ", k, "value: ", v)
			aux = aux + float64(v)
		}
		fmt.Println(aux)
		average = aux / float64(len(history)) //int to float
		s := fmt.Sprintf("%f", average)       // float to string
		*reply = "Promedio de la materia: " + s
	}
	return nil
}

func (this *Server) EveryoneGrade(b bool, reply *string) error {
	*reply = "Hello "
	var aux float64
	var average float64
	var i float64
	for k, v := range science {
		fmt.Println("key: ", k, "value: ", v)
		aux = aux + float64(v)
		i++
		//aux = aux + float64(v)
	}

	for k, v := range math {
		fmt.Println("key: ", k, "value: ", v)
		aux = aux + float64(v)
		i++
		//aux = aux + float64(v)
	}

	for k, v := range history {
		fmt.Println("key: ", k, "value: ", v)
		aux = aux + float64(v)
		i++
		//aux = aux + float64(v)
	}
	average = aux / i               //int to float
	s := fmt.Sprintf("%f", average) // float to string
	*reply = "Promedio general de todas las materias: " + s
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
