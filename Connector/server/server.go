package server

import (
	. "../indexi"
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

var globalAvailableDrives []string = nil

func AllListHandler(x http.ResponseWriter, b *http.Request){
	enableACAO(&x)
	list := GetMusicList(globalAvailableDrives)
	json := list.ToJson()
	x.Write(stringToByteSlice(json))
}

func RefreshedListHandler (x http.ResponseWriter , b *http.Request) {
	enableACAO(&x)
	list := GetRefreshedMusicList(globalAvailableDrives)
	json := list.ToJson()
	x.Write(stringToByteSlice(json))
}

func stringToByteSlice (convertable string) []byte {
	return bytes.Trim([]byte(convertable), "\x00")
}

type Command struct {
	Com string `json:"com"`
	Fun string `json:"fun"`
}

type Commands []Command

func PrintHelp (x http.ResponseWriter , b *http.Request) {
	var commands Commands
	enableACAO(&x)

	command := Command{Com: "/limited_list/[NUM]", Fun: "gives json of [NUM] elements"}
	commands = append(commands, command)
	command = Command{Com: "/refreshed", Fun: "gives json of refreshed list"}
	commands = append(commands, command)
	command = Command{Com: "/list", Fun: "gives json of all the list items"}
	commands = append(commands, command)
	command = Command{Com: "/search/[STRING]", Fun: "search from stuff"}
	commands = append(commands, command)

	commandsJson, _ := json.Marshal(commands)
	x.Write(stringToByteSlice(string(commandsJson)))
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
	list := GetMusicList(globalAvailableDrives)

	if err == nil && sizei < list.Search(searchQuery).Len() {
		x.Write( stringToByteSlice( list.Search(searchQuery)[0 : sizei].ToJson()))
	} else {
		x.Write( stringToByteSlice( list.Search(searchQuery).ToJson()))
	}
}

func ListLimited(x http.ResponseWriter , b * http.Request) {
	enableACAO(&x)

	size := b.FormValue("size")

	list := GetMusicList(globalAvailableDrives)
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
	frag := b.FormValue("file")
	frag , _ = url.QueryUnescape(frag)
	file , err := os.Open(frag)
	if err != nil {
		item := MusicMoreDetail{Artist: "", Title: "", Album: "", Year: 0, Genre: "",}
		json := item.ToJson()
		json = strings.Replace(json , " " , "" , -1)
		x.Write(stringToByteSlice(json))
		return
	}

	m , err := tag.ReadFrom(file)
	if err != nil {
		log.Fatal(err)
		item := MusicMoreDetail{Artist: "", Title: "", Album: "", Year: 0, Genre: "",}
		json := item.ToJson()
		json = strings.Replace(json , " " , "" , -1)
		x.Write(stringToByteSlice(json))
		return
	}

	item := MusicMoreDetail{Artist: m.Artist(), Title: m.Title(), Album: m.Album(), Year: m.Year(), Genre: m.Genre(),}

	json := item.ToJson()
	json = strings.Replace(json , " " , "" , -1)
	x.Write(stringToByteSlice(json))
}

func StartServer (c []string) {
	globalAvailableDrives = c
	r := mux.NewRouter()
	r.HandleFunc("/help", PrintHelp)
	r.HandleFunc("/list" , AllListHandler)
	r.HandleFunc("/refreshed" , RefreshedListHandler)
	r.HandleFunc("/list_limited" , ListLimited)
	r.HandleFunc("/search" , SearchMusicName)
	r.HandleFunc("/item_detail" , GetMusicDetails)
	log.Fatal(http.ListenAndServe("127.0.0.1:8002" , r))
}
