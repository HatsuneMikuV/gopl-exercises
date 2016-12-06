package main

import (
    "time"
    "sort"
    "fmt"
)

type Track struct {
    Title  string
    Artist string
    Album  string
    Year   int
    Length time.Duration
}

func length(s string) time.Duration {
    d, err := time.ParseDuration(s)
    if err != nil {
        panic(s)
    }
    return d
}

var tracks = []*Track{
    {"Go", "Delilah", "Form Roots up", 2012, length("3m38s")},
    {"Go", "Moby", "Moby", 1992, length("3m37s")},
    {"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
    {"Ready 2 Go", "Martin Solveing", "Smash", 2011, length("4m24s")},
}

type byArtist []*Track

func (x byArtist) Len() int {
    return len(x)
}

func (x byArtist) Less(i, j int) bool {
    return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
    x[i], x[j] = x[j], x[i]
}

func printTrack()  {
    for _, track := range byArtist(tracks) {
        fmt.Println(track.Artist)
    }
    fmt.Println()
}

func main() {
    printTrack()
    sort.Sort(byArtist(tracks))
    printTrack()
}
