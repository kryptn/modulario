package http

import (
	"log"
	"math/rand"

	"github.com/kryptn/modulario/proto"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (e *Engine) CreatePost(user User, request proto.CreatePostRequest) (Post, error) {
	post := Post{
		Key:         makeKey(10),
		UserID:      user.ID,
		DeciderType: request.DeciderType,
		Links: func() (links []Link) {
			for _, raw_link := range request.Links {
				links = append(links, Link{Url: raw_link.Url, Condition: Conditional{Threshold: raw_link.Threshold}})
			}
			return
		}(),
	}

	for _, link := range post.Links {
		post.Conditionals = append(post.Conditionals, link.Condition)
	}

	err := e.db.Create(&post).Error
	if err != nil {
		log.Printf("Error on create: %s", err)
		return Post{}, err
	}

	return post, nil
}

func (e *Engine) DeletePost(key string) (err error) {
	err = e.db.Where(&Post{Key: key}).Delete(&Post{}).Error
	return
}

func (e *Engine) GetPost(key string) (post Post, err error) {
	err = e.db.Where(&Post{Key: key}).First(&post).Error
	if err != nil {
		log.Printf("Error on GetPost: %s", err)
		return Post{}, err
	}
	return post, nil
}

func (e *Engine) GetLinks(post *Post) (err error) {
	err = e.db.Model(&post).Related(&post.Links).Error
	return
}

func (e *Engine) GetPostLinksdec(key string) (*Post, error) {
	post, err := e.GetPost(key)
	if err == nil {
		e.GetLinks(&post)
		if err == nil {
			return &post, nil
		}
	}
	return &Post{}, err
}

func (e *Engine) GetPostLinks(key string) (*Post, error) {
	var post Post
	err := e.db.Preload("Links.Condition").Where(&Post{Key: key}).First(&post).Error
	return &post, err
}

func (e *Engine) VisitKey(key string) (Link, error) {
	post, err := e.GetPost(key)
	if err != nil {
		return Link{}, err
	}
	return e.VisitPost(&post)
}

func (e *Engine) VisitPost(post *Post) (link Link, err error) {
	if post.Links == nil {
		err := e.GetLinks(post)
		if err != nil {
			return Link{}, err
		}
	}

	return e.BuildDecider(*post)(), nil
}

func (e *Engine) TotalPostVisits(post *Post) (visits uint) {
	if post.Links == nil {
		e.GetLinks(post)
	}
	for _, link := range post.Links {
		visits += link.Accesses
	}
	return
}

func makeKey(n int) string {
	key := make([]rune, n)
	for i := range key {
		key[i] = runes[rand.Intn(len(runes))]
	}
	return string(key)
}
