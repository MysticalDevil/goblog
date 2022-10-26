package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/article"
)

// ValidateArticleForm 验证表单，返回 errs 长度等于零通过
func ValidateArticleForm(data article.Article) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{"required", "min_cn:3", "max_cn:40"},
		"body": []string{"required", "min_cn:10"},
	}

	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度许下与 40",
		},
		"body": []string {
			"required:文章内容为必填项",
			"min_cn:内容长度需大于 10",
		},
	}

	opts := govalidator.Options{
		Data: &data,
		Rules: rules,
		TagIdentifier: "valid",
		Messages: messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
