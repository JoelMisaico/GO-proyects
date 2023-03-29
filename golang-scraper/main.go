// --------------------->>>>>>> Web scraper con GO
// Tener instalado Go

//Crear archivo main.go, esta sera el punto de entrada de nuestra aplicación
//Inicializamos go.mod con : go mod init example.com/"web-scraper with golang"
//Usaremos el paquete Colly para contruir el webscraper
// usaremos el: go get github.com/gocolly/colly

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	//importaremos el paquete colly para proporcionarnos los métodos y la funcionalidad para construir nuestro web scraper
	"github.com/gocolly/colly"
)

// estructura de lo que buscamos en html de la pagina, el id y la descripcion
type Fact struct {
	ID          int    `json:"id"`          //tipo int
	Description string `json:"description"` // tipo string
}

// función principal
func main() {
	//Creamos un slice vacio para contener nuestros datos que luego lo estaremos llenando
	//este slice solo podrá tener facts
	allFacts := make([]Fact, 0)

	//Usando el paquete colly vamos a crear un nuevo recopliador y establecer que sus dominios permitidos sean: "factretriever.com", "www.factretriever.com"
	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	//Ahora atravesaremos el DOM, usamos el método OnHTML que toma dos argumentos:
	//El primer argumento es el selector de goQuery
	//El segundo argumento es la función de devolución de llamada que se ejecutará cada vez que nuestro recopilador encuentre un elemento de la lista de facts
	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		//El paquete colly utiliza una biblioteca llamada goQuery para interactuar con el DOM, goQuery es el jQuery pero en golang

		//creamos una variable para almacenar el Id de cada elemento que se itera
		//El ID es de tipo cadena por lo que lo convertiremos a tipo int
		factId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			//si recibimos un error simplemente imprimiremos un mensaje que diga que no pudimos obtener un id
			log.Println("Could not get id")
		}

		//Ahora creamos la variable para almacenar el texto de descripción de cada fact(hecho), en tipo de texto(string)
		factDesc := element.Text

		//Ahora creamos una nueva estructura de facts para cada elemento de la lista que iteramos
		fact := Fact{
			ID:          factId,
			Description: factDesc,
		}

		//Aqui queremos agregar la estructura de fact a la sección de allFatcs
		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	//Ahora aqui le decimos la direccion de la web para hacer el scraper
	collector.Visit("https://www.factretriever.com/rhino-facts")

	//Con esto imprimiremos los datos en el archivo rhinofacts.json que crearemos.
	writeJSON(allFacts)
}

// Ahora creamos una funcion para que podamos guardar los datos que recopilamos en un archivo que podamos usar más tarde
func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ") //el método MarshalIndent devuelve la codificación JSON de los datos y tambien devuelve un error
	if err != nil {
		//ponemos un mensaje si no pudimos crear un archivo JSON
		log.Println("Unable to create json file")
		return
	}

	//usamos el paquete de ioutil para crear un archivo json funtamente con su método WriteFile, donde este último creará el archivo
	_ = ioutil.WriteFile("rhinofacts.json", file, 0644) //el archivo se llamará "rhinofacts.json" con el código de permisos: 0644
	//el valor 0644 especifica que el propietario del archivo tiene permisos de lectura y escritura (6), mientras que otros usuarios tienen permisos de solo lectura (4).
}
