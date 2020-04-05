package models

import (
	"fmt"
	u "github.com/mstfymrtc/go-posts-api/utils"
	"time"
)

type Post struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Content   string     `json:"content"`
	/*gorm:"type:integer REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE"*/
	AuthorID uint `json:"author_id";`
	Author   User `json:"author"`
}

func (post *Post) Validate() (map[string]interface{}, bool) {
	if post.Content == "" {
		return u.Message(false, "Post content cannot be empty!"), false
	}
	if post.AuthorID <= 0 {
		return u.Message(false, "User not recognized!"), false
	}

	return u.Message(true, "Post validation successful"), true
}

func (post *Post) Create() (map[string]interface{}) {
	//post not valid
	if resp, ok := post.Validate(); !ok {
		return resp
	}

	//create post
	GetDB().Create(post)
	/*attach post author to post struct*/
	err := GetDB().Debug().Model(&User{}).Where("id=?", post.AuthorID).Take(&post.Author).Error
	if err != nil {
		return u.Message(false, "An error occured!")
	}
	post.Author.Password = ""
	post.Author.CreatedAt = time.Time{}
	post.Author.UpdatedAt = time.Time{}
	post.Author.DeletedAt = nil
	/////////////////////////////////
	resp := u.Message(true, "Post created successfully!")
	resp["post"] = post
	return resp

}

func GetPost(id uint) (*Post) {
	post := &Post{}
	err := GetDB().Table("posts").Where("id=?", id).First(post).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	//get post user
	err = GetDB().Debug().Model(&User{}).Where("id=?", post.AuthorID).Take(&post.Author).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	post.Author.Password = ""
	post.Author.CreatedAt = time.Time{}
	post.Author.UpdatedAt = time.Time{}
	post.Author.DeletedAt = nil
	return post
}
func DeletePost(id uint) (map[string]interface{}) {
	post := &Post{}
	err := GetDB().Table("posts").Where("id=?", id).Delete(post).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return u.Message(true, "Post deleted successfully.")
}

func GetPosts() ([]*Post) {
	posts := make([]*Post, 0)
	err := GetDB().Table("posts").Order("created_at desc").Find(&posts).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return posts
}
