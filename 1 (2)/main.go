package main

func main() {
    // parameter validation
    if len(os.Args) < 2 {
        fmt.Println("please provide an album id")
        return
    }
    albumId, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("album id should be an integer")
        return
    }
    fmt.Printf("album id: %d\n", albumId)

    // get all track list
    tracks, err := getAllTrackList(albumId)
    if err != nil {
        fmt.Printf("error in get all track list, err: %v\n", err)
        return
    }
    fmt.Printf("all track list got, total: %d\n", len(tracks))

    // get audio addresses
    for _, track := range tracks {
        audioAddr, err := getAudioAddress(track.TrackId)
        if err != nil {
            fmt.Printf("error in get audio address, err: %v\n", err)
            break
        }

        // download
        filePath, err := download(audioAddr, track.Title, track.AlbumTitle)
        if err != nil {
            fmt.Printf("error in audo download, err: %v\n", err)
            continue
        }
        fmt.Printf("downloaded! file: %s\n", filePath)
    }
}