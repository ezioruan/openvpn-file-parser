package parser

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type OpenVPNConfig struct {
	lines []string
	CA    string
	Cert  string
	Key   string
}

func NewOpenVPNConfig(lines []string) (OpenVPNConfig, error) {
	o := OpenVPNConfig{lines: lines}
	err := o.process()
	return o, err
}

func (o *OpenVPNConfig) process() error {
	tags := []string{"ca", "cert", "key"}
	linesWithTags := make(map[string][]string)
	tagLines := make([]string, 0)
	tagStarts := false
	isTagStartLine := false
	for _, line := range o.lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		for _, tag := range tags {
			startTag := fmt.Sprintf("<%s>", tag)
			endTag := fmt.Sprintf("</%s>", tag)
			isTagStartLine = strings.TrimSpace(line) == startTag
			if isTagStartLine {
				tagLines = make([]string, 0)
				tagStarts = true
				break
			}
			if strings.TrimSpace(line) == endTag {
				linesWithTags[tag] = tagLines
				tagStarts = false
				break
			}
		}
		if tagStarts && !isTagStartLine {
			tagLines = append(tagLines, line)
		}
	}
	for _, tag := range tags {
		if lines, ok := linesWithTags[tag]; ok {
			content := strings.Join(lines, "\n")
			switch tag {
			case "ca":
				o.CA = content
			case "cert":
				o.Cert = content
			case "key":
				o.Key = content
			}
		} else {
			return fmt.Errorf("Could not get %s context", tag)
		}
	}
	return nil
}

func (o *OpenVPNConfig) SplitFiles(outputDir string) error {
	path, err := filepath.Abs(outputDir)
	if err != nil {
		return err
	}
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	CAPath := filepath.Join(outputDir, "ca.crt")
	clientKeyPath := filepath.Join(outputDir, "client.key")
	clientCertPath := filepath.Join(outputDir, "client.crt")

	err = ioutil.WriteFile(CAPath, []byte(o.CA), 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(clientKeyPath, []byte(o.Key), 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(clientCertPath, []byte(o.Cert), 0644)
	if err != nil {
		return err
	}
	return nil
}

func ParseFromFile(filePath string) (config OpenVPNConfig, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return
	}
	return NewOpenVPNConfig(lines)

}
