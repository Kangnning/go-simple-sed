package sed

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/Kangnning/go-simple-sed/config"
)

type Sed struct {
	conf *config.Config
	reg  *regexp.Regexp
}

func New() *Sed {
	return &Sed{}
}

func (s *Sed) Run(conf *config.Config) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Run error: %v\n", r)
		}
	}()

	s.conf = conf
	err := s.checkConf()
	if err != nil {
		return err
	}

	s.reg = regexp.MustCompile(s.conf.Pattern)
	err = s.run()
	return err
}

// 检查配置文件是否正确
func (s *Sed) checkConf() error {
	if len(s.conf.FileName) == 0 {
		return errors.New("文件名不能为空")
	}

	if _, err := os.Stat(s.conf.FileName); err != nil {
		if os.IsNotExist(err) {
			return errors.New("文件不存在")
		}
		return err
	}

	if len(s.conf.Pattern) == 0 {
		return errors.New("匹配字符串不能为空")
	}

	return nil
}

func (s *Sed) run() error {
	// 创建临时文件
	tmpFile, err := os.CreateTemp("./", "tmp_")
	if err != nil {
		return err
	}
	// defer tmpFile.Close()

	// 打开原始文件
	// fmt.Println(s.conf.FileName)
	origFile, err := os.Open(s.conf.FileName)
	if err != nil {
		return err
	}
	// defer origFile.Close()

	reader := bufio.NewReader(origFile)
	writer := bufio.NewWriter(tmpFile)

	// 标记是否已经找到了模式并插入了字符串
	inserted := false

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}

		if err == io.EOF {
			break
		}

		// 如果尚未插入并且当前行包含模式，则插入字符串
		if !inserted && s.reg.MatchString(line) {
			switch s.conf.Act {
			case config.InsertBefore:
				_, err := writer.WriteString(s.conf.DesString + "\n")
				if err != nil {
					return err
				}
				inserted = true
			case config.Replace:
				_, err := writer.WriteString(s.conf.DesString + "\n")
				if err != nil {
					return err
				}
				fallthrough
			case config.Delete:
				inserted = true
				continue
			}
		}

		_, err = writer.WriteString(line)
		if err != nil {
			return err
		}

		// 如果尚未插入并且当前行包含模式，则插入字符串
		if !inserted && s.conf.Act == config.InsertAfter && s.reg.MatchString(line) {
			_, err := writer.WriteString(s.conf.DesString + "\n")
			if err != nil {
				return err
			}
			inserted = true
		}
	}

	// 刷新缓冲区，确保所有数据都被写入临时文件
	err = writer.Flush()
	if err != nil {
		return err
	}

	// 关闭文件，避免重命名时报错
	origFile.Close()
	tmpFile.Close()
	// 使用临时文件替换原始文件
	err = os.Rename(tmpFile.Name(), s.conf.FileName)
	if err != nil {
		return err
	}

	return nil
}
