package utils

import (
	"fmt"
	"strings"
)

// 준비된 SQL 문장과 파라미터를 받아 쿼리 문자열을 생성하는 함수
func GenerateQueryString(preparedStatement string, params ...string) (string, error) {
	var parmaList []interface{}

	for _, arg := range params {
		parmaList = append(parmaList, arg)
	}
	var queryBuilder strings.Builder
	placeholderCount := strings.Count(preparedStatement, "?")
	if placeholderCount != len(parmaList) {
		return "", fmt.Errorf("플레이스홀더 수(%d)와 파라미터 수(%d)가 일치하지 않습니다", placeholderCount, len(parmaList))
	}

	paramIndex := 0
	for _, char := range preparedStatement {
		if char == '?' {
			// 플레이스홀더를 파라미터로 대체
			param := parmaList[paramIndex]
			paramIndex++

			// 파라미터를 적절히 포매팅
			paramStr, err := formatParameter(param)
			if err != nil {
				return "", err
			}

			queryBuilder.WriteString(paramStr)
		} else {
			queryBuilder.WriteRune(char)
		}
	}

	return queryBuilder.String(), nil
}

// 파라미터를 적절히 포매팅하여 문자열로 변환하는 함수
func formatParameter(param interface{}) (string, error) {
	switch v := param.(type) {
	case string:
		// 작은따옴표 이스케이프 처리
		escaped := strings.ReplaceAll(v, "'", "''")
		return fmt.Sprintf("'%s'", escaped), nil
	case int, int64, float64:
		return fmt.Sprintf("%v", v), nil
	default:
		return "", fmt.Errorf("지원하지 않는 파라미터 타입: %T", param)
	}
}
