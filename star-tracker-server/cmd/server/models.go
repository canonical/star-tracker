package main

type Project struct {
	Name                string `json:"name"`
	UpstreamRepository  string `json:"upstream_repository"`
	MirroredRepository  string `json:"mirrored_repository"`
}

type Component struct {
	Name              string `json:"name"`
	Version           string `json:"version"`
	Project           string `json:"project"`
	SourcecraftName   string `json:"sourcecraft_name"`
	SourcecraftTrack  string `json:"sourcecraft_track"`
	License           string `json:"license"`
}
