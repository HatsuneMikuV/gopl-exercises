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
    {"Go", "Bob", "Form Roots down", 2012, length("3m38s")},
    {"Go", "Moby", "Moby", 1992, length("3m37s")},
    {"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
    {"Ready 2 Go", "Martin Solveing", "Smash", 2011, length("4m24s")},
}

type multiSortTrack []*Track

type multiSort struct {
    t    []*Track
    less func(x, y *Track) bool
}

func (x multiSort) Len() int {
    return len(x.t)
}

func (x multiSort) Less(i, j int) bool {
    return x.less(x.t[i], x.t[j])
}

func (x multiSort) Swap(i, j int) {
    x.t[i], x.t[j] = x.t[j], x.t[i]
}

var clickRecords []string

func less(x, y *Track) bool {
    for _, click := range clickRecords {
        if click == "Title" {
            if x.Title == y.Title {
                continue
            }
            return x.Title < y.Title
        }
        if click == "Year" {
            if x.Year == y.Year {
                continue
            }
            return x.Year < y.Year
        }
        if click == "Artist" {
            if x.Artist == y.Artist {
                continue
            }
            return x.Artist < y.Artist
        }
    }
    return false
}

func printTrack() {
    for _, track := range tracks {
        fmt.Println(track.Title, track.Year, track.Artist)
    }
    fmt.Println()
}

func main() {
    // 模拟点击记录
    clickRecords = append(clickRecords, "Title")
    clickRecords = append(clickRecords, "Year")
    clickRecords = append(clickRecords, "Artist")
    printTrack()
    sort.Sort(multiSort{tracks, less})
    printTrack()
}