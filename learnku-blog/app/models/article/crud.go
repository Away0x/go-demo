package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.Preload("User").Preload("Category").First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

// GetAll 获取全部文章
func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Preload("User").Preload("Category").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil

	// 上面代码相当于原生 sql 查询的
	/*
		1. 执行查询语句，返回一个结果集
		rows, err := db.Query("SELECT * from articles")
		logger.LogError(err)
		defer rows.Close()

		var articles []Article
		2. 循环读取结果
		for rows.Next() {
			var article Article
			2.1 扫码每一行的结果并赋值到一个 article 对象中
			err := rows.Scan(&article.ID, &article.Title, &article.Body)
			logger.LogError(err)
			2.2 将 article 追加到 articles 的这个数组中
			articles = append(articles, article)
		}

		2.3 检测遍历时是否发生错误
		err = rows.Err()
		logger.LogError(err)
	*/
}

// GetAllWithPage 获取全部文章 (分页)
func GetAllWithPage(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	// 1. 初始化分页实例
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}

// Create 创建文章，通过 article.ID 来判断是否创建成功
func (article *Article) Create() (err error) {
	result := model.DB.Create(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Update 更新文章
func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// Delete 删除文章
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// GetByUserID 获取全部文章
func GetByUserID(uid string) ([]Article, error) {
	var articles []Article
	if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// GetByCategoryID 获取分类相关的文章
func GetByCategoryID(cid string, r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	// 1. 初始化分页实例
	db := model.DB.Model(Article{}).Where("category_id = ?", cid).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}
