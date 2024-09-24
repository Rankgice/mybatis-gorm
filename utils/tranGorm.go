package utils

import (
	"gorm.io/gorm"
)

// matchTemplate 匹配字符串中形如 #{xxx} 的值，
// 并且确保这些值不被单引号或双引号包裹。
// target返回的是符合gorm raw()的原生预编译sql
// result返回的是需要的参数列表
// 时间复杂度 O(n)
func matchTemplate(input string) (target string, result []string) {
	target = ""
	nextEnd := 0
	result = make([]string, 0, 3)
	inputRune := []rune(input)
	inputLen := len(inputRune)
	l := 0
	for i := 0; i < inputLen; i++ {
		curRune := inputRune[i]
		switch curRune {
		case '\'':
			if inputRune[l] == '\'' && l != i {
				l = i + 1
			} else if inputRune[l] != '"' && inputRune[l] != '{' {
				l = i
			}
		case '"':
			if inputRune[l] == '"' && l != i {
				l = i + 1
			} else if inputRune[l] != '\'' && inputRune[l] != '{' {
				l = i
			}
		case '{':
			if inputRune[l] != '\'' && inputRune[l] != '"' && i > 0 && inputRune[i-1] == '#' {
				l = i
			}
		case '}':
			if inputRune[l] == '{' && l > 0 && inputRune[l-1] == '#' {
				target += string(inputRune[nextEnd:l-1]) + " ? "
				nextEnd = i + 1
				result = append(result, string(inputRune[l+1:i]))
				l = i + 1
			}
		}
	}
	target += string(inputRune[nextEnd:inputLen])
	return target, result
}

// TranSql 封装示例
// sql为符合mybatis风格的命名参数sql
// args为命名参数map
func TranSql(db *gorm.DB, sql string, args map[string]interface{}) *gorm.DB {
	targetSql, result := matchTemplate(sql)
	values := make([]interface{}, 0, len(result))
	for _, key := range result {
		values = append(values, args[key])
	}
	return db.Raw(targetSql, values...)
}
