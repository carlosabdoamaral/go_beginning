package main

import "fmt"

type post struct {
	title   string
	content string
	author  string
	likes   int
}

func main() {
	post := configPost()

	println("")

	var command int

	for true {
		// 0 - LIKE
		// 1 - CHANGE CONTENT

		println("------------------------------------------------------------")
		println("What do you want to do?")
		println("0 - Like the post")
		println("1 - Change the content")
		println("2 - See post details")
		println("------------------------------------------------------------")

		fmt.Scanln(&command)

		if command == 0 {
			post.likes += 1
			println("")
			println("Okay! We make it")
			println("Likes count: ", post.likes)
			println("")
		} else if command == 1 {
			println("")
			println("Okay! Set the new content:")

			newContent := ""
			fmt.Scanln(&newContent)

			if newContent != "" {
				post.content = newContent
				println("")
				println("Perfect! You change the content")
				println("")
			} else {
				println("")
				println("Ops! Looks like you set an empty content")
				println("")
			}
		} else if command == 2 {
			println("")
			println("POST DETAILS:")
			println("Title: ", post.title)
			println("Content: ", post.content)
			println("Author: ", post.author)
			println("Likes: ", post.likes)
			println("")
		}
	}
}

func configPost() post {
	println("Lets configure the inital post")

	post := post{
		title:   "",
		content: "",
		author:  "",
		likes:   0,
	}

	println("Set a title")
	fmt.Scanln(&post.title)

	println("Set the content")
	fmt.Scanln(&post.content)

	println("Set an author")
	fmt.Scanln(&post.author)

	println("Set the likes")
	fmt.Scanln(&post.likes)

	return post
}
