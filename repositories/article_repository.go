package repositories

import (
	"com.phh/blog/models"
	"github.com/jinzhu/gorm"
)

//repo interface
type ArticleRepo interface {
	//查询文章
	Query(query models.Article) []models.Article
}

//默认构造
func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{
		db: db,
	}
}

//repo struct
type articleRepo struct {
	db *gorm.DB
}

func (a *articleRepo) Query(query models.Article) (results []models.Article) {
	a.db.Where(query).Find(&results)
	return
}
