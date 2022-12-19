package persistance

import (
	"gorm.io/gorm"

	"trash_bot/domain/model/comment"
	"trash_bot/domain/repository"
	"trash_bot/infrastructure/dto"
)

type commentPersistance struct {
	Conn *gorm.DB
}

func NewCommentPersistance(conn *gorm.DB) repository.CommentRepository {
	return &commentPersistance{Conn: conn}
}

// 1件の取得
func (cp *commentPersistance) GetComment(id string) (result *comment.Comment, err error) {
	var comment dto.Comment
	if result := cp.Conn.Where("comment_id = ?", id).First(&comment); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_comment, err := dto.AdaptComment(&comment)
	if err != nil {
		return nil, err
	}

	return result_comment, nil
}

// 一覧の取得
func (cp *commentPersistance) GetComments() (result []*comment.Comment, err error) {
	var comments []*dto.Comment
	if result := cp.Conn.Find(&comments); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_comments, err := dto.AdaptComments(comments)
	if err != nil {
		return nil, err
	}
	return result_comments, nil
}

// 登録
func (cp *commentPersistance) InsertComment(c *comment.Comment) error {
	converted_comment := dto.ConvertComment(c)

	if result := cp.Conn.Create(converted_comment); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

// 更新
func (cp *commentPersistance) UpdateComment(c *comment.Comment) error {
	converted_comment := dto.ConvertComment(c)
	
	if result := cp.Conn.Where("comment_id = ?", converted_comment.CommentId).Updates(converted_comment); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

// 削除
func (cp *commentPersistance) DeleteComment(c *comment.Comment) error {
	converted_comment := dto.ConvertComment(c)

	if result := cp.Conn.Where("comment_id = ?", converted_comment.CommentId).Delete(c); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}