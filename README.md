# One way publishing of your blog posts from a git repo to dev.to

## First, what is dev.to?

https://dev.to is a free and open source blogging platform for developers.

> dev.to (or just DEV) is a platform where software developers write articles, take part in discussions, and build their professional profiles. We value supportive and constructive dialogue in the pursuit of great code and career growth for all members. The ecosystem spans from beginner to advanced developers, and all are welcome to find their place within our community.

## Why would I want to put all my blog posts on a git repo?

- Don't be afraid to mess up with one of your articles while editing it
- Same good practices as when you're developing (format, commits, saving history, compare, etc)
- Use prettier to format the markdown and all the code
- Let people contribute to your article by creating a PR against it (tired of comments going sideways because of some typos? Just let people know they can make a PR at the end of your blog post)
- Create code examples close to your blog post and make sure they're correct thanks to [Embedme](https://github.com/zakhenry/embedme) (_\*1_)

_\*1: Embedme allows you to write code in actual files rather than your readme, and then from your Readme to make sure that your examples are matching those files._

If you prefer not to use Prettier or Embed me, you can do so by simply removing them but I think it's a nice thing to have!

## How do I connect my git repo to dev.to?

There's a `package.json` file where you can define my git repo URL.

```json
{
  "name": "dev.to",
  "repository": {
    "type": "git",
    "url": "<Paste GitHub repository URL>"
  }
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~  
}
```

## How do I choose which files I want to publish?

There's a `dev-to-git.json` file where you can define an array of blog posts, e.g.

```json
[
  {
    "id": 12345,
    "relativePathToArticle": "./template-posts/template-blog.md"
  }
]
```

## How can I find the ID of my blog post on dev.to?

I write Go lang code to get the ID of my blog post on dev.to. Before running code, you need to set DEV Community API Key to `DEVAPIKEY`

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/itchyny/gojq"
)

func curl() interface{} {
	DEVAPIKEY := os.Getenv("DEVAPIKEY") //Set your DEV Community API Key in your environment variables
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://dev.to/api/articles/me/unpublished", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("api-key", DEVAPIKEY)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func main() {
	// Parse JSON
	query, err := gojq.Parse(".[].id")
	if err != nil {
		log.Fatalln(err)
	}
	input := curl()
	iter := query.Run(input)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			log.Fatalln(err)
		}
		fmt.Printf("%1.0f\n", v)
	}
}
```

## How do I configure every blog post individually?

A blog post has to have a [**front matter**](https://dev.to/p/editor_guide) header. You can find an example in this repository here: https://github.com/maxime1992/dev.to/blob/master/blog-posts/name-of-your-blog-post/name-of-your-blog-post.md

Simple and from there you have control over the following properties: `title`, `published`, `cover_image`, `description`, `tags`, `series` and `canonical_url`.

## How do I add images to my blog posts?

Instead of uploading them manually on dev.to, simply put them within your git repo. Here's an example: `![image](./assets/sample.png)`.

If you've got some plugin to preview your markdown from your IDE, the images will be correctly displayed. Then, on CI, right before they're published, the link will be updated to match the raw file.

## How to setup CI for auto deploying the blog posts?

If you want to use Github Actions, a `.build.yml` file has been already prepared for you.

First, you have to create a token on your dev.to account: https://dev.to/settings/account and set an environment variable on GitHUb called `DEV_TO_GIT_TOKEN` that will have the newly created token as value. ![](https://i.imgur.com/euVacys.png)

## README template

The following is simply a template that you may want to use for your own version of that repository.

# Your blog source
