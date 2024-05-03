package db

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type DB []interface{}

func NewDB(path string) DB {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var db DB
	json.Unmarshal(data, &db)

	return db
}

func (db DB) Info() {
	fmt.Printf("文档总数：%d\n", len(db))
}

func noListIndent(text string) string {
	var build strings.Builder
	r, _ := regexp.Compile(`^(\s{4})+\-\s.+$`)
	for _, line := range strings.Split(string(text), "\n") {
		if r.MatchString(line) {
			line = line[4:]
		}

		build.WriteString(line)
		build.WriteString("\n")
	}

	ret := build.String()
	return ret[:len(ret)-1]
}

func (db DB) Get(i int) {
	block := db[i].(map[string]interface{})
	content, ok := block["content"].(string)
	if !ok {
		content = ""
	}
	delete(block, "content")
	data, err := yaml.Marshal(block)
	if err != nil {
		panic(err)
	}
	fm := string(data)
	fmt.Printf("---\n%s---\n\n%s", noListIndent(fm), content)
}

func check(data interface{}, sub string) bool {
	switch v := data.(type) {
	case string:
		if strings.Contains(v, sub) {
			return true
		}
	case []string:
		for _, item := range v {
			if strings.Contains(item, sub) {
				return true
			}
		}
	case []interface{}:
		for _, item := range v {
			if check(item, sub) {
				return true
			}
		}
	case map[string]interface{}:
		for _, item := range v {
			if check(item, sub) {
				return true
			}
		}
	}
	return false
}

func (db DB) Find(subs []string) {
	resultCount := 0
	for i := range db {
		block := db[i].(map[string]interface{})

		found := true
		for _, sub := range subs {
			if !check(block, sub) {
				found = false
			}
		}

		if !found {
			continue
		}

		// 打印 title
		if title, ok := block["title"].(string); ok {
			fmt.Printf("%d\t%s\n", i, title)
		} else if content, ok := block["content"].(string); ok && content != "" {
			line := strings.Split(content, "\n")[0]
			fmt.Printf("%d\t%s\n", i, line)
		} else if date, ok := block["date"].(string); ok {
			fmt.Printf("%d\t%s\n", i, date)
		} else {
			err := fmt.Sprintf("db.fnd(%s) 打印 title 结果失败，i: %d", subs, i)
			panic(err)
		}

		resultCount += 1
		if resultCount >= 20 {
			break
		}
	}
}
