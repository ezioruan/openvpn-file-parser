package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type OpenVPNConfig struct {
	OVPN  string
	lines []string
	CA    string
	Cert  string
	Key   string
}

func NewOpenVPNConfig(lines []string) (OpenVPNConfig, error) {
	o := OpenVPNConfig{lines: lines}
	o.OVPN = strings.Join(lines, "")
	err := o.process()
	return o, err
}

func (o OpenVPNConfig) process() error {
	tags := []string{"ca", "cert", "key"}
	linesWithTags := make(map[string][]string)
	tagLines := make([]string, 0)
	tagStarts := false
	isTagStartLine := false
	for _, line := range o.lines {
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
			content := strings.Join(lines, "")
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
	fmt.Printf("CA ---------------- \n  %s %s %s", o.CA, o.Cert, o.Key)
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
