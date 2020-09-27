package deal

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"

	"golang.org/x/tools/imports"

	"github.com/LinkinStars/baileys/internal/entity"
	"github.com/LinkinStars/baileys/internal/util"
)

// CreateGenPathList 创建目录
func CreateGenPathList(tplList []*entity.TplModel) (err error) {
	for _, tpl := range tplList {
		err := util.CreateDirIfNotExist(tpl.OutputPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// GenFile 创建文件
func GenFile(tpl *template.Template, outputPath, filename string, dataModel interface{}) (err error) {
	resBytes := bytes.NewBufferString("")
	err = tpl.Execute(resBytes, dataModel)
	if err != nil {
		return
	}

	w, err := os.Create(outputPath + filename)
	if err != nil {
		return
	}
	defer w.Close()

	tplContent, err := ioutil.ReadAll(resBytes)
	if err != nil {
		return
	}

	_, err = w.Write(tplContent)
	if err != nil {
		return
	}
	return nil
}

// FormatAndImport 格式化并导包
func FormatAndImport(filename string) (err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	res, err := imports.Process(filename, file, nil)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(filename, res, 0644)
	return err
}
