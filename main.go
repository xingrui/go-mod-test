package gomodtest

import "fmt"

func Hi(name string, lang string) string {
	return fmt.Sprintf("Hi, %s![v2 version lang:%s]", name, lang)
}
