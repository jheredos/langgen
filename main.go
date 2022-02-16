package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	close := ConnectToDB()
	defer close()

	router := httprouter.New()

	router.GET("/", CreateNewLanguage)
	router.GET("/phonology/:id", GetInventory)
	router.POST("/phonology/consonants", UpdateConsonantInventory)
	router.POST("/phonology/vowels", UpdateVowelInventory)

	router.POST("/phonotactics/consonant-hierarchy", UpdateConsonantHierarchy)
	router.POST("/phonotactics/nucleus-hierarchy", UpdateNucleusHierarchy)
	// router.POST("/phonotactics/rules", CreatePhonotacticRules)
	// router.POST("/phonotactics/allophonies", CreateAllophonies)

	router.GET("/lexicon/new-words/:id", GetNewWords)

	router.GET("/ping", Ping)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "https://*.herokuapp.com/"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
	}).Handler(router)

	port := ":" + os.Getenv("PORT")
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(port, handler)
}
