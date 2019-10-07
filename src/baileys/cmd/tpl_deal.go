package main

import (
	"io/ioutil"
	"text/template"

	"baileys/entity"
	"baileys/util"
)

// 读取文件夹将其中的文件制作成tpl模板
func ReadDirGetTemplate(path string, tplFunc template.FuncMap) (tplList []*entity.TplModel, err error) {
	tplList = make([]*entity.TplModel, 0)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return tplList, err
	}

	for _, f := range files {
		tpl, err := CreateTemplate(path+f.Name(), tplFunc)
		if err != nil {
			return tplList, err
		}

		filename := util.GetOnlyFilename(f.Name())

		tplList = append(tplList, &entity.TplModel{
			Tpl:            tpl,
			Filename:       filename,
			FilenameExt:    f.Name(),
			OutputPath:     genRootPath + filename + "/",
			FilenameSuffix: "_" + filename + ".go",
		})
	}

	return tplList, err
}

// 创建模板
func CreateTemplate(tplPath string, tplFunc template.FuncMap) (tpl *template.Template, err error) {
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
