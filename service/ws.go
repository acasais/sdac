package service

import (
	"flag"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/acasais/sdac/configs"
)

var (
	endpoint = flag.String("endpoint", os.Getenv("SDAC_EP"), "endpoint sdac")
)

type WebService struct {
	client   *http.Client
	endpoint string
}

var onceWS sync.Once
var onceWSObj *WebService

func NewWS() *WebService {
	onceWS.Do(func() {
		if *endpoint == "" {
			*endpoint = configs.WS_ENDPOINT
		}
		onceWSObj = &WebService{
			client:   &http.Client{Timeout: 5 * time.Second},
			endpoint: *endpoint,
		}
	})
	return onceWSObj
}

func (ws *WebService) GetPricesForDate(yyyymmdd string) ([]string, error) {

	url := ws.endpoint + yyyymmdd + ".1"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := ws.client.Do(request)
	if err != nil {
		return nil, err
	}

	//prices := []model.Price{}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	yyyy := yyyymmdd[0:4]
	mm := yyyymmdd[4:6]
	dd := yyyymmdd[6:]

	//sample
	//body := "MARGINALPDBC;\r\n2025;01;07;1;97.28;64.07;\r\n2025;01;07;2;72.3;60;\r\n2025;01;07;3;67.41;49.9;\r\n2025;01;07;4;38.41;38.41;\r\n2025;01;07;5;19.2;19.2;\r\n2025;01;07;6;38.7;38.7;\r\n2025;01;07;7;60.08;60.08;\r\n2025;01;07;8;97.56;97.56;\r\n2025;01;07;9;115;115;\r\n2025;01;07;10;105.18;105.18;\r\n2025;01;07;11;93.49;93.49;\r\n2025;01;07;12;74.93;74.93;\r\n2025;01;07;13;71.25;71.25;\r\n2025;01;07;14;71.62;71.62;\r\n2025;01;07;15;72.57;72.57;\r\n2025;01;07;16;85;85;\r\n2025;01;07;17;100.64;100.64;\r\n2025;01;07;18;118.34;118.34;\r\n2025;01;07;19;128.87;128.87;\r\n2025;01;07;20;125;125;\r\n2025;01;07;21;113.97;113.97;\r\n2025;01;07;22;110.76;110.76;\r\n2025;01;07;23;98.53;98.53;\r\n2025;01;07;24;89.27;89.27;\r\n*"
	sbody := strings.Replace(string(body), "\r\n", "", -1)
	sbody = strings.Replace(sbody, yyyy+";"+mm+";"+dd+";", "", -1)
	prices := strings.Split(sbody, ";")
	return prices, nil
}
