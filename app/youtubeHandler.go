package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type YoutubeStats struct {
	Subscribers    int    `json:"subscribers"`
	ChannelName    string `json:"channelname"`
	MinutesWatched int    `json:"minuteswatched"`
	Views          int    `json:"views"`
}

func getChannelStats() httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// w.Write([]byte("Hello, World!"))
		yt := YoutubeStats{
			Subscribers:    50,
			ChannelName:    "Codehakase",
			MinutesWatched: 1000,
			Views:          5000,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(yt); err != nil {
			panic(err)
		}
	})
}
