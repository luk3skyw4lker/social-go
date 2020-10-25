package models_test

import (
	"log"
	"testing"

	"github.com/luk3skyw4lker/social-go/api/models"

	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllPosts(t *testing.T) {
	err := refreshUserAndPostTable()

	if err != nil {
		log.Fatalf("Error on refreshing users and posts table: %v", err)
	}

	_, _, err = seedUsersAndPosts()

	if err != nil {
		log.Fatalf("Error on seeding posts and users: %v", err)
	}

	posts, err := postInstance.FindAllPosts(server.DB)

	if err != nil {
		t.Errorf("Error on find all posts: %v", err)

		return
	}

	assert.Equal(t, len(*posts), 2)
}

func TestSavePost(t *testing.T) {
	err := refreshUserAndPostTable()

	if err != nil {
		log.Fatalf("Error on refreshing users and posts table: %v", err)
	}

	user, err := seedOneUser()

	post := models.Post{
		Title:    "Testing",
		Content:  "Testing",
		AuthorID: user.ID,
	}

	createdPost, err := post.Create(server.DB)

	if err != nil {
		t.Errorf("Error on creating post: %v", err)

		return
	}

	assert.Equal(t, createdPost.AuthorID, post.AuthorID)
	assert.Equal(t, createdPost.Content, post.Content)
	assert.Equal(t, createdPost.Title, post.Title)
	assert.Equal(t, createdPost.ID, post.ID)
}

func TestFindPostById(t *testing.T) {
	err := refreshUserAndPostTable()

	if err != nil {
		log.Fatalf("Error on refreshing users and posts table: %v", err)
	}

	seededPost, err := seedOneUserAndOnePost()

	if err != nil {
		log.Fatalf("Error on seeding posts and users: %v", err)
	}

	post, err := postInstance.FindPostByID(server.DB, seededPost.ID)

	if err != nil {
		t.Errorf("Error on finding post by ID: %v", err)

		return
	}

	assert.Equal(t, seededPost.Content, post.Content)
	assert.Equal(t, seededPost.Title, post.Title)
	assert.Equal(t, seededPost.ID, post.ID)
}

func TestUpdateAPost(t *testing.T) {
	err := refreshUserAndPostTable()

	if err != nil {
		log.Fatalf("Error on refreshing users and posts table: %v", err)
	}

	post, err := seedOneUserAndOnePost()

	if err != nil {
		log.Fatalf("Error on seeding users and posts table: %v", err)
	}

	updatePost := models.Post{
		ID:       post.ID,
		Title:    "new title",
		Content:  "new content",
		AuthorID: post.AuthorID,
	}

	updatedPost, err := updatePost.UpdatePost(server.DB)

	if err != nil {
		t.Errorf("Error on update post: %v", err)
	}

	assert.Equal(t, updatedPost.AuthorID, updatePost.AuthorID)
	assert.Equal(t, updatedPost.Content, updatePost.Content)
	assert.Equal(t, updatedPost.Title, updatePost.Title)
	assert.Equal(t, updatedPost.ID, updatePost.ID)
}

func TestDeleteAPost(t *testing.T) {
	err := refreshUserAndPostTable()

	if err != nil {
		log.Fatalf("Error on refreshing users and posts table: %v", err)
	}

	post, err := seedOneUserAndOnePost()

	if err != nil {
		log.Fatalf("Error on seeding users and posts table: %v", err)
	}

	isDeleted, err := postInstance.DeletePost(server.DB, post.ID, post.AuthorID)

	if err != nil {
		t.Errorf("Error on delete post: %v", err)

		return
	}

	assert.Equal(t, int64(isDeleted), 1)
}
