package utils

import (
	"fmt"
	"time"
)

// Retry 函数接受一个要执行的操作（operation），最大尝试次数（maxRetries）和两次尝试之间的等待时间（delay）
func Retry(operation func() error, maxRetries int, delay time.Duration) {
	for retries := 0; retries < maxRetries; retries++ {
		err := operation()
		if err == nil { // 如果操作成功，返回nil
			return
		}

		// 如果操作失败，打印错误并等待指定的延迟时间后重试
		fmt.Printf("Attempt %d failed; error: %v\n", retries+1, err)
		time.Sleep(delay)
	}
}
