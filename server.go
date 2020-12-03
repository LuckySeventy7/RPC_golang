package main

import (
	//"errors"
	"fmt"
	"net"
	"net/rpc"
)
type Alumno struct{
	Nombre string
	Materia string
	Cali float64
}

type Server struct{
	Materias map[string]map[string]float64
	Alumnos map[string]map[string]float64
}

//func (this *Server) Hello(alumnos map[string]map[string]float64,  reply *string) error {
func (this *Server) Agregar(al Alumno,  reply *string) error {
	
	if _, ok := this.Alumnos[al.Nombre]; ok {//si alumno ya existe
		
		if _, ok := this.Alumnos[al.Nombre][al.Materia]; ok {
			*reply = "Error, Alumno con esta materia ya se fue registrado"
		}else{
			this.Alumnos[al.Nombre][al.Materia] = al.Cali //alumno ya registradi pero agregadno otra materia
			*reply = "Informacion actualizada"
		}
	}else{//si alumno es registrado por primer vez
	
			grade := make(map[string]float64)
			grade[al.Materia] = al.Cali
			this.Alumnos[al.Nombre] = grade
			*reply = "Informacion actualizada"
		
	}
	
	
	return nil
}

func (this *Server) PromAlumno(nom string, reply *string) error {
	prom:=0.0
	sum:=0.0
	if _, ok := this.Alumnos[nom]; !ok {
		*reply ="Alumno No encontrado"
	}else{
	
		for i:= range this.Alumnos[nom]{
			sum+= this.Alumnos[nom][i]
		}
		prom = sum/float64(len(this.Alumnos[nom]))
		*reply = "Promedio del Almuno: " + fmt.Sprintf("%f",prom)
	}
	
	return nil
}
func (this *Server) PromGen( aux float64,  reply *string) error {
	prom:=0.0
	sum:=0.0
	sum2:=0.0
	if len(this.Alumnos) ==0 {
		fmt.Println("Lista vacia")
	}else{
	
		for i:= range this.Alumnos{
			sum=0.0
			for j:= range this.Alumnos[i]{
				sum+= this.Alumnos[i][j]
				
			}
			sum2+=sum/float64(len(this.Alumnos[i]))
		}


	}
	//fmt.Println(sum2)
	//fmt.Println(len(this.Alumnos[nom]))
	prom = sum2/float64(len(this.Alumnos))
	*reply = "Promedio General de todo los alumnos: " + fmt.Sprintf("%f",prom)
	return nil
}

func (this *Server) PromMat(mat string,  reply *string) error {
	prom:=0.0
	sum:=0.0
	le:=0.0
	if len(this.Alumnos) ==0 {
		*reply = "Lista vacia"
	}else{
	
		for i:= range this.Alumnos{
			//sum=0.0
			for  range this.Alumnos[i]{
				if _, ok := this.Alumnos[i][mat]; ok {
					//fmt.Println(this.Alumnos[i], ",", mat, ",",this.Alumnos[i][mat])
					sum+= this.Alumnos[i][mat]
					le++
					break
				}
			}
		}

		prom = sum/le
		*reply = "Promedio General de todo los alumnos: " + fmt.Sprintf("%f",prom)
	}
	//fmt.Println(le)
	//fmt.Println(sum)
	
	return nil
}


func server() {
	sl:= new(Server)
	sl.Alumnos = make(map[string]map[string]float64)
	sl.Materias = make(map[string]map[string]float64)
	rpc.Register((sl))
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