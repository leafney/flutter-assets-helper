package main

import (
	"fmt"
	"github.com/leafney/flutter-assets-helper/utils"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// 得到所有的ios图标文件
func loadFiles() []IosGroup {

	pwd, _ := os.Getwd()
	// 待处理文件所在目录
	pendingPwd := "./tmp"

	basePath := path.Join(pwd, pendingPwd)
	fmt.Println(basePath)

	files, err := ioutil.ReadDir(pendingPwd)
	if err != nil {
		fmt.Println(err)
	}

	resList := make([]IosGroup, 0)
	//resList := make([]string, 0)
	for _, f := range files {
		fileName := f.Name()
		fileIsDir := f.IsDir()
		fmt.Println(fileName, fileIsDir)

		if !fileIsDir {
			// 处理文件
			//path.Base(fileName)
			// 获取文件扩展名及文件名称
			fileSuffix := path.Ext(fileName)
			fileNameOnly := strings.TrimSuffix(fileName, fileSuffix)
			isHaveAt := strings.Contains(fileNameOnly, "@")

			// 如果是png文件且名称中不存在@的则添加
			if strings.ToLower(fileSuffix) == ".png" && !isHaveAt {
				iosPng := IosGroup{
					Name:         fileNameOnly,
					fileBasePath: basePath,
				}
				resList = append(resList, iosPng)

				//resList = append(resList, fileNameOnly)
			}
		} else {
			//	处理目录
		}
	}

	fmt.Println("结果： \n", resList)
	return resList
}

type IosGroup struct {
	Name         string `json:"name"` // 文件名，无扩展名
	fileBasePath string
	//At1x map[string]string `json:"at_1x"`
	//At2x map[string]string `json:"at_2x"`
	//At3x map[string]string `json:"at_3x"`
}

// 将ios图标文件拷贝到目标目录下的指定目录中
func copyIosFiles(pngList []IosGroup) {

	pwd, _ := os.Getwd()
	// 目标根目录
	targetPwd := "./test"
	fileExt := ".png"

	targetBasePwd := path.Join(pwd, targetPwd)

	for _, png := range pngList {
		fmt.Println("开始处理： ", png.Name)

		for i := 1; i < 4; i++ {

			theFileName := png.Name
			newFileName := png.Name
			newFilePwd := "."
			if i > 1 {
				theFileName = fmt.Sprintf("%s@%dx", theFileName, i)
				newFilePwd = fmt.Sprintf("%d.0x", i)
			}
			theFilePath := path.Join(png.fileBasePath, fmt.Sprintf("%s%s", theFileName, fileExt))
			newFileBasePath := path.Join(targetBasePwd, newFilePwd)
			//	判断源文件是否存在
			theFileExist := utils.Exists(theFilePath)

			fmt.Println("newFileName: ", newFileName, " newFileBasePath: ", newFileBasePath)
			fmt.Println("theFilePath: ", theFilePath, " theFileExist: ", theFileExist)

			if theFileExist {

				//	判断目标目录是否存在
				if !utils.Exists(newFileBasePath) {
					if err := os.MkdirAll(newFileBasePath, 0711); err != nil {
						// Todo err
						fmt.Println("创建目标目录异常：", err)
					}
				}

				//	拷贝文件
				newFilePath := path.Join(newFileBasePath, fmt.Sprintf("%s%s", newFileName, fileExt))
				if err := copyFile(theFilePath, newFilePath); err != nil {
					// Todo err
					fmt.Println("文件拷贝异常：", err)
				}
			}
		}
	}

	fmt.Println("处理完毕")
}

// 拷贝文件
func copyFile(sourcePath, destPath string) error {
	inputF, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputF, err := os.Create(destPath)
	if err != nil {
		inputF.Close()
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer outputF.Close()

	_, err = io.Copy(outputF, inputF)
	inputF.Close()
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	return nil
}
