package utils

import (
	"strings"
	"testing"
)

func TestGenPasswd(t *testing.T) {
	password := "TestPassword123!"

	// 测试密码哈希生成
	hash, err := GenPasswd(password)
	if err != nil {
		t.Fatalf("生成密码哈希失败: %v", err)
	}

	// 验证哈希格式
	if !strings.HasPrefix(hash, "$argon2id$") {
		t.Errorf("哈希格式不正确，期望以 $argon2id$ 开头，实际: %s", hash)
	}

	// 验证哈希长度不为空
	if len(hash) == 0 {
		t.Error("生成的哈希不能为空")
	}

	t.Logf("生成的哈希: %s", hash)
}

func TestComparePasswd(t *testing.T) {
	password := "TestPassword123!"
	wrongPassword := "WrongPassword123!"

	// 生成哈希
	hash, err := GenPasswd(password)
	if err != nil {
		t.Fatalf("生成密码哈希失败: %v", err)
	}

	// 测试正确密码验证
	err = ComparePasswd(hash, password)
	if err != nil {
		t.Errorf("正确密码验证失败: %v", err)
	}

	// 测试错误密码验证
	err = ComparePasswd(hash, wrongPassword)
	if err == nil {
		t.Error("错误密码验证应该失败，但返回了成功")
	}
}

func TestGenPasswdWithParams(t *testing.T) {
	password := "TestPassword123!"

	// 自定义参数
	customParams := &Argon2Params{
		Memory:      32 * 1024, // 32 MB
		Iterations:  2,         // 2 次迭代
		Parallelism: 1,         // 1 个并行线程
		SaltLength:  16,        // 16 字节盐值
		KeyLength:   32,        // 32 字节密钥
	}

	hash, err := GenPasswdWithParams(password, customParams)
	if err != nil {
		t.Fatalf("使用自定义参数生成密码哈希失败: %v", err)
	}

	// 验证可以正确验证
	err = ComparePasswd(hash, password)
	if err != nil {
		t.Errorf("自定义参数密码验证失败: %v", err)
	}

	// 验证参数是否正确编码
	if !strings.Contains(hash, "m=32768") {
		t.Errorf("哈希中应包含自定义内存参数 m=32768")
	}
	if !strings.Contains(hash, "t=2") {
		t.Errorf("哈希中应包含自定义迭代参数 t=2")
	}
	if !strings.Contains(hash, "p=1") {
		t.Errorf("哈希中应包含自定义并行参数 p=1")
	}
}

func TestParseHash(t *testing.T) {
	password := "TestPassword123!"

	// 生成哈希
	hash, err := GenPasswd(password)
	if err != nil {
		t.Fatalf("生成密码哈希失败: %v", err)
	}

	// 解析哈希
	params, salt, hashBytes, err := parseHash(hash)
	if err != nil {
		t.Fatalf("解析哈希失败: %v", err)
	}

	// 验证参数
	if params.Memory != DefaultArgon2Params.Memory {
		t.Errorf("解析的内存参数不匹配，期望: %d，实际: %d", DefaultArgon2Params.Memory, params.Memory)
	}
	if params.Iterations != DefaultArgon2Params.Iterations {
		t.Errorf("解析的迭代参数不匹配，期望: %d，实际: %d", DefaultArgon2Params.Iterations, params.Iterations)
	}
	if params.Parallelism != DefaultArgon2Params.Parallelism {
		t.Errorf("解析的并行参数不匹配，期望: %d，实际: %d", DefaultArgon2Params.Parallelism, params.Parallelism)
	}

	// 验证盐值和哈希不为空
	if len(salt) == 0 {
		t.Error("解析的盐值不能为空")
	}
	if len(hashBytes) == 0 {
		t.Error("解析的哈希不能为空")
	}
}

func TestParseHashInvalidFormat(t *testing.T) {
	invalidHashes := []string{
		"invalid",
		"$argon2id$v=19$m=65536",
		"$bcrypt$v=19$m=65536,t=3,p=2$salt$hash",
		"$argon2id$v=18$m=65536,t=3,p=2$salt$hash",
	}

	for _, invalidHash := range invalidHashes {
		_, _, _, err := parseHash(invalidHash)
		if err == nil {
			t.Errorf("无效哈希 '%s' 应该解析失败，但返回了成功", invalidHash)
		}
	}
}

func TestValidatePasswordStrength(t *testing.T) {
	testCases := []struct {
		password string
		valid    bool
		desc     string
	}{
		{"Password123!", true, "有效的强密码"},
		{"password123!", false, "缺少大写字母"},
		{"PASSWORD123!", false, "缺少小写字母"},
		{"Password!", false, "缺少数字"},
		{"Password123", false, "缺少特殊字符"},
		{"Pass1!", false, "密码太短"},
		{"", false, "空密码"},
		{"Complex@Pass123", true, "另一个有效的强密码"},
	}

	for _, tc := range testCases {
		err := ValidatePasswordStrength(tc.password)
		if tc.valid && err != nil {
			t.Errorf("密码 '%s' (%s) 应该有效，但验证失败: %v", tc.password, tc.desc, err)
		}
		if !tc.valid && err == nil {
			t.Errorf("密码 '%s' (%s) 应该无效，但验证通过", tc.password, tc.desc)
		}
	}
}

func TestHashUniqueness(t *testing.T) {
	password := "TestPassword123!"

	// 生成多个哈希，验证它们是不同的（因为使用了随机盐值）
	hash1, err := GenPasswd(password)
	if err != nil {
		t.Fatalf("生成第一个哈希失败: %v", err)
	}

	hash2, err := GenPasswd(password)
	if err != nil {
		t.Fatalf("生成第二个哈希失败: %v", err)
	}

	if hash1 == hash2 {
		t.Error("相同密码的哈希应该不同（由于随机盐值）")
	}

	// 但两个哈希都应该能验证相同的密码
	if err := ComparePasswd(hash1, password); err != nil {
		t.Errorf("第一个哈希验证失败: %v", err)
	}

	if err := ComparePasswd(hash2, password); err != nil {
		t.Errorf("第二个哈希验证失败: %v", err)
	}
}

func BenchmarkGenPasswd(b *testing.B) {
	password := "TestPassword123!"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GenPasswd(password)
		if err != nil {
			b.Fatalf("生成密码哈希失败: %v", err)
		}
	}
}

func BenchmarkComparePasswd(b *testing.B) {
	password := "TestPassword123!"
	hash, err := GenPasswd(password)
	if err != nil {
		b.Fatalf("生成密码哈希失败: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := ComparePasswd(hash, password)
		if err != nil {
			b.Fatalf("密码验证失败: %v", err)
		}
	}
}
