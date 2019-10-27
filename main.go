package gomodtest

import "fmt"

func Hi(name string, lang string, time string) string {
	return fmt.Sprintf("Hi, %s![v2 version lang:%s time:%s]", name, lang, time)
}
