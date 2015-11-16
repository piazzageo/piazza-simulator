package registry

import (
	"github.com/mpgerlek/piazza-simulator/piazza"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"encoding/json"
)

//---------------------------------------------------------------------

type RegistryService struct {
	table *serviceTable		`json:"table"`
	proxy *RegistryProxy
}

//---------------------------------------------------------------------

func (registry *RegistryService) postService(w http.ResponseWriter, r *http.Request) {
	log.Printf("postService")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entry, err := piazza.NewServiceFromBytes(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err2 := registry.table.Add(entry)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	var bytes []byte
	bytes, err = entry.ToBytes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bytes)

	log.Printf("added entry for %s, id:%v (cnt:%d)", entry.Type, entry.Id, registry.table.Len())
}


func (registry *RegistryService) getService(w http.ResponseWriter, r *http.Request) {
	//log.Printf("getService")
	var buf []byte
	buf, err := registry.table.ToBytes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(buf)
}

func (registry *RegistryService) getServiceItem(w http.ResponseWriter, r *http.Request) {
	//log.Printf("getServiceItem")
	var err error
	var ok bool

	key := mux.Vars(r)["key"]

	var list *piazza.ServiceList

	if idInt, err := strconv.ParseInt(key, 10, 64); err == nil {
		id := piazza.PiazzaId(idInt)
		if list, ok = registry.table.LookupById(id); !ok {
			http.Error(w, "id not found", http.StatusBadRequest)
			return
		}
	} else {
		stype := piazza.ServiceTypeFromString(key)
		if stype == piazza.InvalidService {
			list = registry.table.LookupAll()
		} else if list, ok = registry.table.LookupByType(stype); !ok {
			http.Error(w, "service type not found", http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")

	buf, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(buf)
}

func NewRegistryService(registryHost string, comments string) error {
	var err error

	log.Printf("registry started: %s", registryHost)

	registry := new(RegistryService)
	registry.table = NewServiceTable()

	registry.proxy, err = NewRegistryProxy(registryHost)
	if err != nil {
		log.Fatal(err)
		return err
	}
	registry.proxy.Register(registryHost, piazza.RegistryService, comments)

	// TODO: DELETE

	r := mux.NewRouter()
	r.HandleFunc("/service", registry.postService).Methods("POST")
	r.HandleFunc("/service", registry.getService).Methods("GET")
	r.HandleFunc("/service/{key}", registry.getServiceItem).Methods("GET")

	server := &http.Server{Addr: registryHost, Handler: r}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// not reached
	return nil
}

/**func (*RegistryProxy) Lookup(stype piazza.ServiceType) (value string, err error) {
	entry, ok := table.LookupByType(stype)
	if !ok {
		return "", errors.New(fmt.Sprintf("%v not found in registry", stype))
	}

	return entry.Host, nil
}**/
