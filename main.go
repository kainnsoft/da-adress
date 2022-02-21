package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	dadata "github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
)

// параметры локации для запроса
var locationsSlice []*suggest.RequestParamsLocation
var tmpl *template.Template

func init() {
	err := os.Setenv("DADATA_API_KEY", "6b022035bf9d381b46bffdcd317b2c5745019f76")
	if err != nil {
		log.Fatalf("error setting api_key: %v", err)
	}

	locationsSlice = make([]*suggest.RequestParamsLocation, 1)
	mapLocations := suggest.RequestParamsLocation{
		RegionFiasID: "88cd27e2-6a8a-4421-9718-719a28a0a088", // Нижегородская обл.
	}
	locationsSlice = append(locationsSlice, &mapLocations)

	// web
	tmpl = template.Must(template.ParseGlob("web/templates/*"))
	if err != nil {
		log.Fatal("Couldn't parse template...")
	}
}

func main() {

	//------------------------------------------------------------------
	// web
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/index/", index())
	http.HandleFunc("/da", indexDa())
	//------------------------------------------------------------------
	http.HandleFunc("/dadata", handleDadataRequest())
	//------------------------------------------------------------------
	// server   --------------------------------------------------------
	//------------------------------------------------------------------
	addr := flag.String("addr", ":8099", "Сетевой адрес веб-сервера Hint")
	flag.Parse()

	srv := &http.Server{
		Addr:    *addr,
		Handler: nil,
	}
	err := srv.ListenAndServe()
	log.Fatalf("main: srv.ListenAndServe() error: %v", err)
}

func handleDadataRequest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		t := r.URL.Query()
		if len(t) < 1 {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("There are no parametrs."))
			return
		}
		if len(t) > 1 {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Too many parametrs."))
			return
		}

		adressQueryStr := strings.TrimSpace(r.URL.Query().Get("adress"))
		if adressQueryStr != "" {

			// request from 1C
			// if strings.Contains(r.Header["User-Agent"][0], "Enterprise") {
			// 	// TODO LRU Cash
			// 	remAddr := r.RemoteAddr
			// 	assertLenRemAddr := len(remAddr) - 6
			// 	remAddr = remAddr[:assertLenRemAddr]
			// 	t, ok := globalDelayMap.getTime(remAddr)
			// 	if !ok {
			// 		globalDelayMap.writeTime(remAddr, time.Now())
			// 		return
			// 	}
			// 	if time.Since(t) < 5*time.Second {
			// 		globalDelayMap.writeTime(remAddr, time.Now()) // рестартуем таймер
			// 		return
			// 	}
			// 	globalDelayMap.delRecord(remAddr)
			// }

			// получаем весь массив сотрудников, который будем возвращать, т.к. параметр обрабатывается как like
			jsonSliceOfAddressSuggestion := adressMarshal(adressQueryStr)
			rw.WriteHeader(http.StatusOK)
			rw.Write(jsonSliceOfAddressSuggestion)
			return
		} else {
			http.NotFound(rw, r)
			return
		}

	}
}

func toDadataRequest(adressQueryStr string) []*suggest.AddressSuggestion {
	api := dadata.NewSuggestApi()
	params := suggest.RequestParams{Query: adressQueryStr, Locations: locationsSlice}
	result, err := api.Address(context.Background(), &params)

	if err != nil {
		fmt.Println("Error = ", err)
	}
	fmt.Println("result = ", result)

	// for i, mapa := range result {
	// 	fmt.Printf("i = %d , result = %v, type = %T \n", i, mapa, *mapa)
	// 	fmt.Println("", mapa.Value)
	// 	fmt.Println("", mapa.Data.RegionFiasID)
	// 	//break
	// }
	return result

	// apiKey, ok := os.LookupEnv("DADATA_API_KEY")
	// if !ok {
	// 	fmt.Println("no api keys in env")
	// }
	// fmt.Println("apiKey = ", apiKey)

}

func adressMarshal(adressQueryStr string) []byte {
	// пока ВСЕ (не один) из базы:
	usersByUserNameSlice := toDadataRequest(adressQueryStr)

	sliceOfByte, err := json.MarshalIndent(usersByUserNameSlice, "", "  ")
	if err != nil {
		fmt.Println("marshal GetUsersByNameLightVersionAttributes error", err)
	}

	return sliceOfByte
}

func index() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(rw, "index_first.gohtml", "123")
		if err != nil {
			log.Fatal("Couldn't execute template index_first.gohtml...")
		}
	}
}

func indexDa() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(rw, "index.gohtml", "123")
		if err != nil {
			log.Fatal("Couldn't execute template index.gohtml...")
		}
	}
}
