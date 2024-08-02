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

func getChannelStats(k string, channelId string) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// w.Write([]byte("Hello, World!"))
		ctx := context.Background()
		yts, err := youtube.NewService(ctx, option.WithAPIKey(k))
		if err != nil {
			fmt.Println("failed to create a service")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		call := yts.Channels.List([]string{"snippet,contentDetails,statistics"})
		resp, err := call.Id(channelId).Do()
		if err != nil {
			fmt.Println("failed to fetch channel stats")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(resp.Items[0].Snippet.Title)

		yt := YoutubeStats{}
		if len(resp.Items) > 0 {
			val := resp.Items[0]
			yt = YoutubeStats{
				Subscribers: int(val.Statistics.SubscriberCount),
				ChannelName: val.Snippet.Title,
				Views:       int(val.Statistics.ViewCount),
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(yt); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// log.Fatal("error while encoding the response")
			return
		}
	})
}
