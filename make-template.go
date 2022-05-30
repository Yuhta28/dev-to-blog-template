package main

import (
	"fmt"
	"os"
)

func main() {

	var blog string

	print("Enter the name of the new article: ")
	fmt.Scan(&blog)

	// Create blog directory and article file
	if err := os.MkdirAll("blog-posts/"+blog, 0777); err != nil {
		fmt.Println(err)
	}
	file_article, err := os.Create("blog-posts/" + blog + "/" + blog + ".md")
	if err != nil {
		fmt.Println(err)
	}
	defer file_article.Close()

	// Add blog metadata in article file
	file_article.WriteString("---\ntitle: Title \npublished: false\ndescription: description\ntags: tag1, tag2\n---\n")

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
