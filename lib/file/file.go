package file

import (
	"io/ioutil"
	"os"
)

/**
 * @Author: yaoqiang
 * @Description: 读取文件
 * @Date: 2020/8/28 5:09 下午
 * @param null
 * @return:
 **/
func loadFile(filename string) ([]byte, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return buf, nil
}
