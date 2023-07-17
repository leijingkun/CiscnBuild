package utils

//返回数组里随机一个元素
import "math/rand"

func randomFrom(arr []string) string {
	n := len(arr)

	// 生成0到n-1之间的随机数字
	idx := rand.Intn(n)

	return arr[idx]
}
