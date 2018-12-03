package server

import (
	"../indexi"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tag-master"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func AllListHandler(x http.ResponseWriter, b *http.Request){
	enableACAO(&x)
	list := indexi.GetMusicList()
	json := list.ToJson()
	x.Write(stringToByteSlice(json))
}

func RefreshedListHandler (x http.ResponseWriter , b *http.Request) {
	enableACAO(&x)
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
	enableACAO(&x)

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
	enableACAO(&x)
	searchQuery := b.FormValue("q")
	searchQuery , errorq := url.QueryUnescape(searchQuery)

	if errorq != nil {
		print(errorq)
		searchQuery = b.FormValue("q")
	}

	size := b.FormValue("s")
	sizei , err := strconv.Atoi(size)
	list := indexi.GetMusicList()

	if err == nil && sizei < list.Search(searchQuery).Len() {
		x.Write( stringToByteSlice( list.Search(searchQuery)[0 : sizei].ToJson()))
	} else {
		x.Write( stringToByteSlice( list.Search(searchQuery).ToJson()))
	}
}

func ListLimited(x http.ResponseWriter , b * http.Request) {
	enableACAO(&x)

	size := b.FormValue("size")

	list := indexi.GetMusicList()
	sizei , _ := strconv.Atoi(size)

	list = list[0 : sizei]
	json := list.ToJson()
	json = strings.Replace(json , " " , "" , -1)
	x.Write(stringToByteSlice(json))
}

func UseTagGo() {
	file , _ := os.Open("C://t.mp3")
	m , err := tag.ReadFrom(file)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(m.Format()) // The detected format.
	log.Print(m.Title())
}

func enableACAO( rw *http.ResponseWriter) {
	(*rw).Header().Set("Access-Control-Allow-Origin" , "*")
}

func GetMusicDetails (x http.ResponseWriter , b *http.Request) {
	enableACAO(&x)
	frag := strings.Replace(b.URL.Path , "/item_detail/" , "" , -1)
	frag , _ = url.QueryUnescape(frag)
	file , err := os.Open(frag)
	if err != nil {
		log.Fatal(err)
	}

	m , err := tag.ReadFrom(file)
	if err != nil {
		log.Fatal(err)
	}

	item := indexi.MusicMoreDetail{Artist: m.Artist(), Title: m.Title(), Album: m.Album(), Year: m.Year(), Genre: m.Genre(),}

	json := item.ToJson()
	json = strings.Replace(json , " " , "" , -1)
	x.Write(stringToByteSlice(json))
}

func StartServer () {
	r := mux.NewRouter()
	r.HandleFunc("/help", PrintHelp)
	r.HandleFunc("/list" , AllListHandler)
	r.HandleFunc("/refreshed" , RefreshedListHandler)
	r.HandleFunc("/list_limited" , ListLimited)
	r.HandleFunc("/search" , SearchMusicName)
	r.HandleFunc("/item_detail/{file:.}" , GetMusicDetails)
	log.Fatal(http.ListenAndServe("127.0.0.1:8002" , r))
}
