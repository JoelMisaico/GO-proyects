// Vamos a realizar un web scrapper
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Datos de la página web a scrapear
var baseURL = "https://www.example.com"
var urls = []string{"/page1", "/page2", "/page3"}

// Estructura de datos para almacenar los resultados del scraping
type Result struct {
	Title string
	URL   string
}

func main() {
	// Crear un archivo CSV para almacenar los resultados del scraping
	file, err := os.Create("scraping_results.csv")
	if err != nil {
		log.Fatal("No se pudo crear el archivo CSV", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Crear un canal para almacenar los resultados del scraping
	resultChan := make(chan Result)

	// Crear un grupo de espera para sincronizar los goroutines
	var wg sync.WaitGroup

	// Crear un goroutine para cada URL a scrapear
	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			// Hacer una solicitud HTTP a la página web que quieres scrapear
			res, err := http.Get(baseURL + url)
			if err != nil {
				log.Println("Error al obtener la página", url, err)
				return
			}
			defer res.Body.Close()

			// Analizar el HTML de la página web utilizando la biblioteca goquery
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Println("Error al analizar el HTML de la página", url, err)
				return
			}

			// Extraer los datos que deseas del HTML utilizando selectores de goquery
			doc.Find("h1").Each(func(i int, s *goquery.Selection) {
				title := strings.TrimSpace(s.Text())
				result := Result{
					Title: title,
					URL:   baseURL + url,
				}
				resultChan <- result
			})
		}(url)
	}

	// Cerrar el canal de resultados cuando se hayan procesado todas las URL
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Escribir los resultados en el archivo CSV
	for result := range resultChan {
		err := writer.Write([]string{result.Title, result.URL})
		if err != nil {
			log.Println("Error al escribir en el archivo CSV", err)
			continue
		}
	}

	fmt.Println("El scraping ha finalizado. Los resultados se han almacenado en el archivo scraping_results.csv")
}
