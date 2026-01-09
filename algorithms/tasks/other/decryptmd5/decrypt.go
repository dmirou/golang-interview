// У нас есть база данных с паролями пользователей, пароли захешированы
// (функция hashPassword), а так же известен набор символов которые могут быть
// использованы в паролях (переменная alphabet). Наша задача реализовать функцию
// RecoverPassword так, чтобы она восстанавливала пароль по известному хэшу и
// TestRecoverPassword завершился успешно.
package decryptmd5

import (
	"bytes"
	"crypto/md5"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}
var maxLength = 7

func RecoverPassword(h []byte) string {

	var str string

	pass, _ := decode(str, h)

	return string(pass)
}

func decode(str string, h []byte) (string, bool) {
	hash := md5.Sum([]byte(str))
	if bytes.Equal(hash[:], h) {
		return str, true
	}

	if len(str) > maxLength {
		return "", false
	}

	for _, r := range alphabet {
		pass, ok := decode(str+string(r), h)
		if ok {
			return pass, ok
		}
	}

	return "", false
}

func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
