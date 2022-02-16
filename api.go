package main

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jheredos/langgen/phonology"
	"github.com/jheredos/langgen/phonotactics"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// Pool represents the database connection pool
var Pool *sql.DB

// ConnectToDB establishes a new connection with the database
func ConnectToDB() func() {
	uri := os.Getenv("DATABASE_URL")
	var err error
	Pool, err = sql.Open("postgres", uri) // using := causes a null pointer error for Pool???

	if err != nil {
		panic(err)
	}

	return func() {
		Pool.Close()
	}
}

// GetLanguageID checks the user's cookies to see if they have an already
// existing language. If not, it creates a new one. The second return val
// denotes whether the id already exists or not
func GetLanguageID(w http.ResponseWriter, r *http.Request) (string, bool) {
	c, _ := r.Cookie("lang_id")
	var id string
	var alreadyExists bool
	if c == nil {
		fmt.Println("Found no cookie, making one...")
		alreadyExists = false
		id = uuid.NewV4().String()
		// http.SetCookie(w, &http.Cookie{
		// 	Name:     "lang_id",
		// 	Value:    id,
		// 	HttpOnly: true,
		// })
	} else {
		alreadyExists = true
		id = c.Value
	}
	return id, alreadyExists
}

func CreateNewLanguage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := uuid.NewV4().String()
	fmt.Println("CreateNewLanguage: ", id)

	data, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Ping is for testing
func Ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Pong!")
	fmt.Println(GetLanguageID(w, r))
	data, _ := json.Marshal("Pong!")
	// http.SetCookie(w, &http.Cookie{
	// 	Name:  "test_cookie",
	// 	Value: "Mmmm cookie",
	// 	// HttpOnly: true,
	// })
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// UnmarshalBinary decodes a gob ([]byte) into the data type provided in the second arg, which must be a pointer
func UnmarshalBinary(source []byte, destination interface{}) error {
	buf := bytes.NewBuffer(source)
	return gob.NewDecoder(buf).Decode(destination)
}

// MarshalBinary encodes the value in source to gob
func MarshalBinary(source interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	err := gob.NewEncoder(&buf).Encode(source)
	return buf.Bytes(), err
}

// GetInventory ...
func GetInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("getInventory")

	id := ps.ByName("id")
	row := Pool.QueryRow(`SELECT consonants, vowels FROM languages WHERE lang_id=$1`, id)
	if row.Err() == sql.ErrNoRows {
		http.Error(w, fmt.Sprintf("No language with id \"%s\" found.", id), http.StatusBadRequest)
		return
	}

	csb, vsb := []byte{}, []byte{}
	err := row.Scan(&csb, &vsb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cs, vs := []phonology.Consonant{}, []phonology.Vowel{}

	err = UnmarshalBinary(csb, &cs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = UnmarshalBinary(vsb, &vs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(&struct {
		Consonants []phonology.Consonant `json:"consonants"`
		Vowels     []phonology.Vowel     `json:"vowels"`
	}{
		Consonants: cs,
		Vowels:     vs,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// UpdateConsonantInventory ...
func UpdateConsonantInventory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("UpdateConsonantInventory")

	var reqData struct {
		ID   string                `json:"id"`
		Data []phonology.Consonant `json:"data"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cs := reqData.Data
	id := reqData.ID

	bs, err := MarshalBinary(cs)

	stmt := `INSERT INTO languages (lang_id, consonants) VALUES ($1, $2) ON CONFLICT (lang_id) DO UPDATE SET consonants=$2 WHERE languages.lang_id=$1;`
	_, err = Pool.Exec(stmt, id, bs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, _ := json.Marshal(fmt.Sprintf("Successfully updated consonant inventory for language %s", id))
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// UpdateVowelInventory ...
func UpdateVowelInventory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("UpdateVowelInventory")

	var reqData struct {
		ID   string            `json:"id"`
		Data []phonology.Vowel `json:"data"`
	}

	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vs := reqData.Data
	id := reqData.ID

	bs, err := MarshalBinary(vs)

	stmt := `INSERT INTO languages (lang_id, vowels) VALUES ($1, $2) ON CONFLICT (lang_id) DO UPDATE SET vowels=$2 WHERE languages.lang_id=$1;`
	_, err = Pool.Exec(stmt, id, bs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, _ := json.Marshal(fmt.Sprintf("Successfully updated vowel inventory for language %s", id))
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// UpdateConsonantHierarchy ...
func UpdateConsonantHierarchy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("UpdateConsonantHierarchy")
	var reqData struct {
		ID   string                          `json:"id"`
		Data phonotactics.ConsonantHierarchy `json:"data"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ch := reqData.Data
	id := reqData.ID

	bs, err := MarshalBinary(ch)

	stmt := `INSERT INTO languages (lang_id, onset_clusters) VALUES ($1, $2) ON CONFLICT (lang_id) DO UPDATE SET onset_clusters=$2 WHERE languages.lang_id=$1;`
	if !ch.Onset {
		stmt = `INSERT INTO languages (lang_id, coda_clusters) VALUES ($1, $2) ON CONFLICT (lang_id) DO UPDATE SET coda_clusters=$2 WHERE languages.lang_id=$1;`

	}
	result, err := Pool.Exec(stmt, id, bs)
	affected, _ := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if affected == 0 {
		http.Error(w, "Failed to update consonant clusters", http.StatusInternalServerError)
	}

	res, _ := json.Marshal(fmt.Sprintf("Successfully updated consonant clusters for language %s", id))
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// UpdateNucleusHierarchy ...
func UpdateNucleusHierarchy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("UpdateNucleusHierarchy")
	var reqData struct {
		ID   string                        `json:"id"`
		Data phonotactics.NucleusHierarchy `json:"data"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nh := reqData.Data
	id := reqData.ID

	bs, err := MarshalBinary(nh)

	stmt := `INSERT INTO languages (lang_id, nucleus_clusters) VALUES ($1, $2) ON CONFLICT (lang_id) DO UPDATE SET nucleus_clusters=$2 WHERE languages.lang_id=$1;`
	res, err := Pool.Exec(stmt, id, bs)
	affected, _ := res.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if affected == 0 {
		http.Error(w, "Failed to update consonant clusters", http.StatusInternalServerError)
	}

	data, _ := json.Marshal(fmt.Sprintf("Successfully updated vowel inventory for language %s", id))
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// func CreatePhonotacticRules(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Println("CreatePhonotacticRules")

// }

// func CreateAllophonies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Println("CreateAllophonies")

// }

// GetNewWords ...
func GetNewWords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("GetNewWords")
	id := ps.ByName("id")

	var onsets phonotactics.ConsonantHierarchy
	var nuclei phonotactics.NucleusHierarchy
	var codas phonotactics.ConsonantHierarchy
	var ob, nb, cb []byte
	row := Pool.QueryRow(`SELECT onset_clusters, nucleus_clusters, coda_clusters FROM languages WHERE lang_id=$1`, id)
	if row.Err() == sql.ErrNoRows {
		http.Error(w, fmt.Sprintf("No language with id \"%s\" found.", id), http.StatusBadRequest)
		return
	}

	err := row.Scan(&ob, &nb, &cb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = UnmarshalBinary(ob, &onsets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = UnmarshalBinary(nb, &nuclei)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = UnmarshalBinary(cb, &codas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	root, err := phonotactics.NewPhonotacticTree(onsets, nuclei, codas)
	root.SetHiatus(phonotactics.NeverRF)
	wordGen := phonotactics.NewWordGenerator(root)

	words := []string{}
	for i := 0; i < 30; i++ {
		length := phonotactics.GetWordLength(phonotactics.MonosyllabicWL, phonotactics.ShortWL, phonotactics.MediumWL)
		words = append(words, wordGen.NewWord(length))
	}

	data, err := json.Marshal(words)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
