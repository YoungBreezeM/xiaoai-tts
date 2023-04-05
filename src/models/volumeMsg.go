package models

import "encoding/json"

type VolumeMsg struct {
	Code int64 `json:"code"`
	Info Info  `json:"info"`
}

type Info struct {
	Status         int64          `json:"status"`
	Volume         int64          `json:"volume"`
	LoopType       int64          `json:"loop_type"`
	MediaType      int64          `json:"media_type"`
	PlaySongDetail PlaySongDetail `json:"play_song_detail"`
	TrackList      []string       `json:"track_list"`
}

type PlaySongDetail struct {
	AudioID  string `json:"audio_id"`
	Position int64  `json:"position"`
	Duration int64  `json:"duration"`
}

func (i *Info) String() string {
	v, _ := json.MarshalIndent(i, "", " ")
	return string(v)
}

func (vm *VolumeMsg) String() string {
	v, _ := json.MarshalIndent(vm, "", " ")
	return string(v)
}
