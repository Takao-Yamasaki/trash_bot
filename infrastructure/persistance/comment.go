package persistance

import (
	"gorm.io/gorm"

	"trash_bot/domain/model"
	"trash_bot/domain/repository"
)

type commentPersistance struct {
	Conn *gorm.DB
}

func NewCommentPersistance(conn *gorm.DB, c repository.CommentRepository) *commentPersistance {
	return &commentPersistance{Conn: conn}
}

// 1件の取得
func (co *commentPersistance) GetComment(id int) (result *model.Comment, err error) {
	
	var comment model.Comment
	if result := co.Conn.First(&comment, id); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &comment, nil
}

// 一覧の取得
func (co *commentPersistance) GetComments() (result []model.Comment, err error) {
	
	var cms []model.Comment
	if result := co.Conn.Find(&cms); result.Error != nil {
		err := result.Error
		return nil, err
	}
	return cms, nil
}

// 登録
func (co *commentPersistance) Create(cm model.Comment) error {
	if result := co.Conn.Create(&cm); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

// 更新
func (co *commentPersistance) Update(cm model.Comment) error {
	if result := co.Conn.Save(&cm); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

// 削除
func (co *commentPersistance) Delete(cm model.Comment) error {
	if result := co.Conn.Delete(&cm); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}