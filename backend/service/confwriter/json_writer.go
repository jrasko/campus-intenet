package confwriter

import (
	"backend/model"
	"encoding/json"
	"net/http"
	"os"
)

type JsonWriter struct {
	filename string
}

func New(filename string) JsonWriter {
	return JsonWriter{
		filename: filename,
	}
}

type reservation struct {
	IP  string `json:"ip-address"`
	Mac string `json:"hw-address"`
}

func (dw JsonWriter) WhitelistUsers(member []model.MemberConfig) error {
	reservations := mapToReservationUser(member)

	f, err := os.OpenFile(dw.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return model.Error(http.StatusInternalServerError, err.Error(), "could not open output file")
	}

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(reservations)
	if err != nil {
		return model.Error(http.StatusInternalServerError, err.Error(), "error on writing output file")
	}

	err = f.Close()
	if err != nil {
		return err
	}

	reloadConfig()
	return nil
}

func mapToReservationUser(configs []model.MemberConfig) []reservation {
	reservations := make([]reservation, 0, len(configs))
	for _, config := range configs {
		reservations = append(reservations, reservation{
			IP:  config.IP,
			Mac: config.Mac,
		})
	}
	return reservations
}
