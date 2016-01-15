package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	// api := slack.New("xoxb-14316137920-HnYQrHOS7pOtkLKEDRh1i30w")
	api := slack.New("xoxp-8035809824-8036215361-17780810228-8a32658b4c")
	params := slack.FileUploadParameters{
		Title: "New Markdown Example",
		Filetype: "post",
		Filename: "example.md",
		// File: "example.md",
		Content:  "# hello world!\n*italics*\n**bold**\n`code`\n",
		InitialComment: "whoa, programmatic post creation!",
		Channels:  []string{"#hubotix"},
	}
	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%+v\n", file)

	/*
	err = api.DeleteFile(file.ID)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("File %s deleted successfully.\n", file.Name)
	*/
}
