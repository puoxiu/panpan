package utils

import (
	"math/rand"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
    "time"
)

// HashPassword 对密码进行加密
func HashPassword(password string) (string) {
    // 生成哈希值，cost 是加密的强度，数值越大越安全但也越耗时，一般 10 - 14 是比较合适的范围
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        logx.Error("HashPassword error: ", err)
        return ""
    }
    return string(hashedPassword)
}

// CheckPasswordHash 验证密码是否匹配
func CheckPasswordHash(password, hashedPassword string) bool {
    // 比较密码和哈希值
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}



// GeneratePassword 生成随机密码
func GeneratePassword(length int) string {
    // 定义密码包含的字符集
    charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.!_"
    // 使用新的方式创建一个随机数生成器
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    password := make([]byte, length)
    for i := 0; i < length; i++ {
        password[i] = charset[r.Intn(len(charset))]
    }
    return string(password)
}