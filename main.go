package main

import (
	"fmt"
)

//estructura imagen y su respectivo metodo mostrar
type Imagen struct {
	titulo  string
	formato string
	canales int64
}

func (im *Imagen) mostrar() {
	fmt.Println("Imagen")
	fmt.Println("Titulo: ", im.titulo)
	fmt.Println("Formato: ", im.formato)
	fmt.Println("Canales: ", im.canales)
}

//estructura audio y su respectivo metodo mostrar
type Audio struct {
	titulo   string
	formato  string
	duracion int64
}

func (im *Audio) mostrar() {
	fmt.Println("Audio")
	fmt.Println("Titulo: ", im.titulo)
	fmt.Println("Formato: ", im.formato)
	fmt.Println("Duracion: ", im.duracion)
}

//estructura video y su respectivo metodo mostrar
type Video struct {
	titulo  string
	formato string
	frames  int64
}

func (im *Video) mostrar() {
	fmt.Println("Video")
	fmt.Println("Titulo: ", im.titulo)
	fmt.Println("Formato: ", im.formato)
	fmt.Println("Frames: ", im.frames)
}

// interface multimedia
type Multimedia interface {
	mostrar()
}

//estructura con slice de la interfrace multimedia
type ContenidoWeb struct {
	contenido []Multimedia
}

//funcion para llamar al metodo mostar de cada elemento multimedia
func (cw *ContenidoWeb) mostrar() {
	for _, multi := range cw.contenido {
		multi.mostrar()
	}
}

//funcion para capturar los datos
func capturaDatos(text string) (titulo, formato string, cdf int64) {
	fmt.Print("Titulo: ")
	fmt.Scan(&titulo)
	fmt.Print("Formato: ")
	fmt.Scan(&formato)
	fmt.Print(text, ": ")
	fmt.Scan(&cdf)
	return titulo, formato, cdf
}
func main() {
	var opc uint64
	cw := ContenidoWeb{contenido: []Multimedia{}}
	i := 1
	for i < 100 {
		//menu
		fmt.Println("Selecione que desea capturar")
		fmt.Println("1.- Imagen")
		fmt.Println("2.- Audio")
		fmt.Println("3.- Video")
		fmt.Println("4.- Mostrar")
		fmt.Println("5.- Salir")
		fmt.Print("Ingrese el numero de la opcion: ")
		fmt.Scan(&opc)

		switch opc {
		case 1:
			fmt.Println("Imagen")
			titulo, formato, CDF := capturaDatos("Canales")
			im01 := Imagen{titulo, formato, CDF}
			cw.contenido = append(cw.contenido, &im01)
		case 2:
			fmt.Println("Audio")
			titulo, formato, CDF := capturaDatos("Duracion")
			aud01 := Audio{titulo, formato, CDF}
			cw.contenido = append(cw.contenido, &aud01)
		case 3:
			fmt.Println("Video")
			titulo, formato, CDF := capturaDatos("Frames")
			vid01 := Video{titulo, formato, CDF}
			cw.contenido = append(cw.contenido, &vid01)
		case 4:
			cw.mostrar()
		case 5: //terminar el ciclo
			i = 150
		default:
			fmt.Println("Error!")
		}
		i = i + 1

		fmt.Println("____________________________________________")
	}
}
