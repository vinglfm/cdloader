package storage

import (
	"fmt"
	"io"
	"os"
	"prj/cdloader/request"
	"strings"
	"time"
)

func SaveAll(videos []string, folder string, startIndex int) error {
	for idx := startIndex; idx < len(videos); idx++ {
		err := save(videos[idx], folder)
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

func save(video string, folder string) error {
	urlParts := strings.Split(video, "/")
	filePath := folder + "/" + urlParts[len(urlParts)-1]
	fmt.Println(filePath)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	stream, err := request.DownloadVideo(video)
	if err != nil {
		return err
	}
	defer stream.Close()

	io.Copy(out, stream)
	return nil
}
