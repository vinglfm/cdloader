package main

import (
	"flag"
	"fmt"
	"prj/cdloader/auth"
	"prj/cdloader/request"
	"prj/cdloader/storage"
	"strings"
)

func main() {
	courseUrl := flag.String("url", "", "Downloading url")
	email := flag.String("email", "", "User account email")
	password := flag.String("password", "", "User account password")
	path := flag.String("path", "", "Path to folder for storing downlaoded videos")
	startsFrom := flag.Int("startIndex", 1, "Video number to start from")
	flag.Parse()
	if *courseUrl == "" {
		err := fmt.Errorf("%s\n", "No course url specified, set course url -url=course_url")
		fmt.Println(err.Error())
		return
	}

	token, err := auth.Authenticate(*email, *password)
	if err != nil {
		fmt.Println("Can't authenticate:", err)
		return
	}

	videos, err := request.GetVideoUrls(*courseUrl, token)
	if err != nil {
		fmt.Println("Error retrieving videos:", err)
		return
	}

	if *path == "" {
		splittedUrl := strings.Split(*courseUrl, "/")
		*path = splittedUrl[len(splittedUrl)-1]
	}
	storage.SaveAll(videos, *path, *startsFrom-1)
}
