package main

import (
	"fmt"
	"os"
)

func main() {

	var blog string

	print("Enter the name of the new article: ")
	fmt.Scan(&blog)

	// Create blog directory
	if err := os.MkdirAll("blog-posts/"+blog, 0777); err != nil {
		fmt.Println(err)
	}
	_, err := os.Create("blog-posts/" + blog + "/" + blog + ".md")
	if err != nil {
		fmt.Println(err)
	}

	// Create code directory
	if err := os.MkdirAll("blog-posts/"+blog+"/code", 0777); err != nil {
		fmt.Println(err)
	}
	file_code, err := os.Create("blog-posts/" + blog + "/code/.gitkeep")
	if err != nil {
		fmt.Println(err)
	}
	defer file_code.Close()

	// Create assets directory
	if err := os.MkdirAll("blog-posts/"+blog+"/assets", 0777); err != nil {
		fmt.Println(err)
	}
	file_assets, err := os.Create("blog-posts/" + blog + "/assets/.gitkeep")
	if err != nil {
		fmt.Println(err)
	}
	defer file_assets.Close()
}
