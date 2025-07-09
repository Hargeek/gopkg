package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

// Argon2id 参数配置
type Argon2Params struct {
	Memory      uint32 // 内存使用量 (KB)
	Iterations  uint32 // 迭代次数
	Parallelism uint8  // 并行度
	SaltLength  uint32 // 盐值长度
	KeyLength   uint32 // 密钥长度
}

// 默认 Argon2id 参数（推荐配置）
var DefaultArgon2Params = &Argon2Params{
	Memory:      64 * 1024, // 64 MB
	Iterations:  3,         // 3 次迭代
	Parallelism: 2,         // 2 个并行线程
	SaltLength:  16,        // 16 字节盐值
	KeyLength:   32,        // 32 字节密钥
}

// GenPasswd 使用 Argon2id 算法生成密码哈希
// passwd: 明文密码
// 返回: 编码后的哈希字符串，格式为 $argon2id$v=19$m=65536,t=3,p=2$salt$hash
func GenPasswd(passwd string) (string, error) {
	return GenPasswdWithParams(passwd, DefaultArgon2Params)
}

// GenPasswdWithParams 使用指定参数生成密码哈希
func GenPasswdWithParams(passwd string, params *Argon2Params) (string, error) {
	// 生成随机盐值
	salt := make([]byte, params.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("生成盐值失败: %w", err)
	}

	// 使用 Argon2id 生成密钥
	hash := argon2.IDKey([]byte(passwd), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	// 编码为标准格式
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// 格式: $argon2id$v=19$m=memory,t=iterations,p=parallelism$salt$hash
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, params.Memory, params.Iterations, params.Parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

// ComparePasswd 验证密码是否匹配
// hashPasswd: 存储的哈希密码
// passwd: 明文密码
// 返回: 如果密码匹配返回 nil，否则返回错误
func ComparePasswd(hashPasswd, passwd string) error {
	// 解析哈希字符串
	params, salt, hash, err := parseHash(hashPasswd)
	if err != nil {
		return fmt.Errorf("解析哈希失败: %w", err)
	}

	// 使用相同参数和盐值对输入密码进行哈希
	inputHash := argon2.IDKey([]byte(passwd), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	// 使用常数时间比较防止时序攻击
	if subtle.ConstantTimeCompare(hash, inputHash) == 1 {
		return nil
	}

	return errors.New("密码不匹配")
}

// parseHash 解析 Argon2id 哈希字符串
func parseHash(encodedHash string) (*Argon2Params, []byte, []byte, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return nil, nil, nil, errors.New("无效的哈希格式")
	}

	if parts[1] != "argon2id" {
		return nil, nil, nil, errors.New("不支持的算法类型")
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return nil, nil, nil, fmt.Errorf("解析版本失败: %w", err)
	}

	if version != argon2.Version {
		return nil, nil, nil, errors.New("不支持的 Argon2 版本")
	}

	params := &Argon2Params{}
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism); err != nil {
		return nil, nil, nil, fmt.Errorf("解析参数失败: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("解析盐值失败: %w", err)
	}
	params.SaltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("解析哈希失败: %w", err)
	}
	params.KeyLength = uint32(len(hash))

	return params, salt, hash, nil
}

// ValidatePasswordStrength 验证密码强度
func ValidatePasswordStrength(passwd string) error {
	if len(passwd) < 8 {
		return errors.New("密码长度至少为8位")
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range passwd {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case strings.ContainsRune("!@#$%^&*()_+-=[]{}|;:,.<>?", char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("密码必须包含至少一个大写字母")
	}
	if !hasLower {
		return errors.New("密码必须包含至少一个小写字母")
	}
	if !hasNumber {
		return errors.New("密码必须包含至少一个数字")
	}
	if !hasSpecial {
		return errors.New("密码必须包含至少一个特殊字符")
	}

	return nil
}
