package data

import (
	"log"
	"math/rand"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (e *Engine) CreatePost(user User, raw_links []string) (Post, error) {
	post := Post{
		Key:    makeKey(10),
		UserID: user.ID,
		Links: func() (links []Link) {
			for _, raw_link := range raw_links {
				links = append(links, Link{Url: raw_link})
			}
			return
		}(),
	}

	err := e.db.Create(&post).Error
	if err != nil {
		log.Printf("Create Error yo: %s", err)
		return Post{}, err
	}

	//err = e.db.Commit().Error; if err != nil {
	//	log.Printf("Commit Error yo: %s", err)
	//	return Post{}, err
	//}
	return post, nil
}

func (e *Engine) DeletePost(key string) (err error) {
	err = e.db.Where(&Post{Key: key}).Delete(&Post{}).Error
	return
}

func (e *Engine) GetPost(key string) (post Post, err error) {
	err = e.db.Where(&Post{Key: key}).First(&post).Error
	if err != nil {
		log.Printf("havin an issue it seems in getpost: %s", err)
		return Post{}, err
	}
	return post, nil
}

func (e *Engine) GetLinks(post *Post) (err error) {
	err = e.db.Model(&post).Related(&post.Links).Error
	return
}

func (e *Engine) GetPostLinks(key string) (*Post, error) {
	post, err := e.GetPost(key)
	if err == nil {
		e.GetLinks(&post)
		if err == nil {
			return &post, nil
		}
	}
	return &Post{}, err
}

func (e *Engine) VisitPost(post *Post) (link Link, err error) {
	if post.Links == nil {
		err := e.GetLinks(post)
		if err != nil {
			return Link{}, err
		}
	}
	link = post.Links[rand.Intn(len(post.Links))]
	link.Accesses += 1
	e.db.Save(&link)
	return link, err
}

func makeKey(n int) string {
	key := make([]rune, n)
	for i := range key {
		key[i] = runes[rand.Intn(len(runes))]
	}
	return string(key)
}
