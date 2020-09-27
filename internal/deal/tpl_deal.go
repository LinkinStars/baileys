package deal

import (
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/LinkinStars/baileys/internal/cache"
	"github.com/LinkinStars/baileys/internal/entity"
	"github.com/LinkinStars/baileys/internal/util"
)

// ReadDirGetTemplate 读取文件夹将其中的文件制作成tpl模板
func ReadDirGetTemplate(path string, tplFunc template.FuncMap) (tplList []*entity.TplModel, err error) {
	tplList = make([]*entity.TplModel, 0)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return tplList, err
	}

	for _, f := range files {
		// 根据文件后缀过滤无关文件
		if !strings.Contains(f.Name(), ".tpl") {
			continue
		}
		tpl, err := createTemplate(path+f.Name(), tplFunc)
		if err != nil {
			return tplList, err
		}

		// 获取没有带有后缀的文件名
		filename := util.GetOnlyFilename(f.Name())

		tplList = append(tplList, &entity.TplModel{
			Tpl:             tpl,
			Filename:        filename,
			FilenameWithExt: f.Name(),
			OutputPath:      cache.GenRootPath + filename + "/",
			FilenameSuffix:  "_" + filename + ".go",
		})
	}

	return tplList, err
}

// createTemplate 创建模板
func createTemplate(tplPath string, tplFunc template.FuncMap) (tpl *template.Template, err error) {
	bs, err := ioutil.ReadFile(tplPath)
	if err != nil {
		return
	}

	t := template.New(tplPath).Funcs(tplFunc)
	tpl, err = t.Parse(string(bs))
	if err != nil {
		return
	}

	return
}
