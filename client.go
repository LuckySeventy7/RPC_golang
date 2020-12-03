package main

import (
	"fmt"
	"net/rpc"
)
type Alumno struct{
	Nombre string
	Materia string
	Cali float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	var al Alumno

	for {
		fmt.Println("------------------------------------")
		fmt.Println("1) Agregar calificacion")
		fmt.Println("2) Mostrar el promedio de un Alumno")
		fmt.Println("3) Mostrar el promedio general")
		fmt.Println("4) Mostrar el promedio de una materia")
		fmt.Println("0) Exit")
		fmt.Print("- Ingrese Opcion: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var name string
			fmt.Print("Name: ")
			fmt.Scanln(&name)

			var materia string
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)

			var cali float64
			fmt.Print("Calificacion: ")
			fmt.Scanln(&cali)

			al.Nombre = name
			al.Materia = materia
			al.Cali =cali
			
			var result string
			err = c.Call("Server.Agregar", al, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case 2:
			var nom string
			fmt.Print("Nombre: ")
			fmt.Scanln(&nom)

			var result string
			err = c.Call("Server.PromAlumno", nom, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case 3:
			aux:=0.0
			var result string
			err = c.Call("Server.PromGen",  aux, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(  result)
			}
		case 4:
			var mat string
			var result string
			fmt.Print("Materia: ")
			fmt.Scanln(&mat)
			err = c.Call("Server.PromMat",  mat, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}