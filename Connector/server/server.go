package server

import (
	"../indexi"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	//"github.com/mikkyang/id3-go"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func AllListHandler(x http.ResponseWriter, b *http.Request){
	enableCors(&x)
	list := indexi.GetMusicList()
	json := list.ToJson()
	x.Write(stringToByteSlice(json))
}

func RefreshedListHandler (x http.ResponseWriter , b *http.Request) {
	enableCors(&x)
	list := indexi.GetRefreshedMusicList()
	json := list.ToJson()
	x.Write(stringToByteSlice(json))
}

func stringToByteSlice (convertable string) []byte {
	return bytes.Trim([]byte(convertable), "\x00")
}

//func testMp3Loading (x string) {
//	music , err := id3.Open(x)
//
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("Music Artist : " , music.Artist())
//	fmt.Println("Music Album : " , music.Album())
//	fmt.Println("Music Genre : " , music.Genre())
//	fmt.Println("Music Year : " , music.Year())
//
//}

type Command struct {
	Com string `json:"com"`
	Fun string `json:"fun"`
}

type Commands []Command

func PrintHelp (x http.ResponseWriter , b *http.Request) {
	var commands Commands
	enableCors(&x)

	command := Command{Com: "/limited_list/[NUM]", Fun: "gives json of [NUM] elements"}
	commands = append(commands , command)
	command = Command{Com: "/refreshed", Fun: "gives json of refreshed list"}
	commands = append(commands , command)
	command = Command{Com: "/list", Fun: "gives json of all the list items"}
	commands = append(commands , command)
	command = Command{Com: "/search/[STRING]", Fun: "search from stuff"}
	commands = append(commands , command)

	commandsJson , _ := json.Marshal(commands)
	x.Write(stringToByteSlice(string(commandsJson)))
	//x.Write(stringToByteSlice("{\n \"/limited_list/[NUM]\" : \"gives json of [NUM] elements\" ,\n \"/refreshed\" : \"gives json of refreshed list\",\n \"/list\" : \"gives json of all the list items\"\n}"))
}

func SearchMusicName (x http.ResponseWriter , b *http.Request) {
	enableCors(&x)
	searchQuery := strings.Replace(b.URL.Path , "/search/" , "" ,-1)
	list := indexi.GetMusicList()
	x.Write( stringToByteSlice( list.Search(searchQuery).ToJson()))
}

func ListLimited(x http.ResponseWriter , b * http.Request) {
	enableCors(&x)
	list := indexi.GetMusicList()
	frag := strings.Replace(b.URL.Path , "/list_limited/" , "" , -1)
	limit , err := strconv.Atoi(frag)

	if err != nil {
		log.Panic(err)
	}

	list = list[0 : limit]
	json := list.ToJson()
	json = strings.Replace(json , " " , "" , -1)
	x.Write(stringToByteSlice(json))
}

func enableCors ( rw *http.ResponseWriter) {
	(*rw).Header().Set("Access-Control-Allow-Origin" , "*")
}

func StartServer () {
	r := mux.NewRouter()
	r.HandleFunc("/help", PrintHelp)
	r.HandleFunc("/list" , AllListHandler)
	r.HandleFunc("/refreshed" , RefreshedListHandler)
	r.HandleFunc("/list_limited/{size:[0-9]+}" , ListLimited)
	//testMp3Loading("C:\\one.mp3")
	r.HandleFunc("/search/{name:[a-zA-Z0-9 ]+}" , SearchMusicName)
	log.Fatal(http.ListenAndServe("127.0.0.1:8002" , r))
}
