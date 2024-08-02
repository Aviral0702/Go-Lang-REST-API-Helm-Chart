package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeStats struct {
	Subscribers    int    `json:"subscribers"`
	ChannelName    string `json:"channelname"`
	MinutesWatched int    `json:"minuteswatched"`
	Views          int    `json:"views"`
}

func getChannelStats(k string) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// w.Write([]byte("Hello, World!"))
		yt := YoutubeStats{
			Subscribers:    50,
			ChannelName:    "Codehakase",
			MinutesWatched: 1000,
			Views:          5000,
		}
		ctx := context.Background()
		yts, err := youtube.NewService(ctx, option.WithAPIKey(k))
		if err != nil {
			fmt.Println("failed to create a service")
		}
		call := yts.Channels.List([]string{"snippet,contentDetails,statistics"})
		resp, err := call.ForUsername().Do()
		fmt.Println("Response ", resp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(yt); err != nil {
			panic(err)
		}
	})
}
