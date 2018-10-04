package server

import (
	"../indexi"
	"../mux-master"
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func AllListHandler(x http.ResponseWriter, b *http.Request){
	list := indexi.GetMusicList()
	json := indexi.MusicListToJson(list)
	x.Write(stringToByteSlice(json))
}

func RefreshedListHandler (x http.ResponseWriter , b *http.Request) {
	list := indexi.GetRefreshedMusicList()
	json := indexi.MusicListToJson(list)
	x.Write(stringToByteSlice(json))
}

func stringToByteSlice (convertable string) []byte {
	return bytes.Trim([]byte(convertable), "\x00")
}

func ListLimited(x http.ResponseWriter , b * http.Request) {
	list := indexi.GetMusicList()
	frag := strings.Replace(b.URL.Path , "/list_limited/" , "" , -1)
	limit , err := strconv.Atoi(frag)

	if err != nil {
		log.Panic(err)
	}

	list = list[0 : limit]
	json := indexi.MusicListToJson(list)
	json = strings.Replace(json , " " , "" , -1)
	x.Write(stringToByteSlice(json))
}

func StartServer () {
	r := mux.NewRouter()
	r.HandleFunc("/", AllListHandler)
	r.HandleFunc("/list" , AllListHandler)
	r.HandleFunc("/refreshed" , RefreshedListHandler)
	r.HandleFunc("/list_limited/{size:[0-9]+}" , ListLimited)
	log.Fatal(http.ListenAndServe("127.0.0.1:8002" , r))
}
