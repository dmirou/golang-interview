
```text
// Просто медленная функция
func slowFunc() int64 {
	// slow operations
	time.Sleep(2 * time.Second)
	return time.Now().Unix()
}

// Нужно написать функцию обертку над slowFunc, которая использует контекст,
// не меняя ее сигнатуру. То есть мы завершаем выполнение с ошибкой если контекст завершается
// раньше чем мы получили ответ от slowFunc.
func slowFuncWithContext(ctx context.Context) (int64, error) {
	// TODO implement me
}
```
