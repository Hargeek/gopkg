package main

import (
	"fmt"
	"github.com/hargeek/gopkg/encrypt/argon2id"
	"log"
)

func main() {
	// 示例密码
	password := "MySecure@Password123"

	// 1. 密码强度验证
	fmt.Println("=== 密码强度验证 ===")
	if err := utils.ValidatePasswordStrength(password); err != nil {
		log.Printf("密码强度不足: %v", err)
		return
	}
	fmt.Println("✓ 密码强度验证通过")

	// 2. 生成密码哈希
	fmt.Println("\n=== 生成密码哈希 ===")
	hashedPassword, err := utils.GenPasswd(password)
	if err != nil {
		log.Fatalf("生成密码哈希失败: %v", err)
	}
	fmt.Printf("原始密码: %s\n", password)
	fmt.Printf("哈希密码: %s\n", hashedPassword)

	// 3. 验证密码
	fmt.Println("\n=== 验证密码 ===")
	err = utils.ComparePasswd(hashedPassword, password)
	if err != nil {
		log.Printf("密码验证失败: %v", err)
	} else {
		fmt.Println("✓ 密码验证成功")
	}

	// 4. 错误密码验证
	wrongPassword := "WrongPassword123"
	err = utils.ComparePasswd(hashedPassword, wrongPassword)
	if err != nil {
		fmt.Printf("✓ 错误密码验证失败（符合预期）: %v\n", err)
	} else {
		fmt.Println("✗ 错误密码验证成功（不符合预期）")
	}

	// 5. 使用自定义参数
	fmt.Println("\n=== 使用自定义参数 ===")
	customParams := &utils.Argon2Params{
		Memory:      32 * 1024, // 32 MB （较低的内存使用）
		Iterations:  2,         // 2 次迭代（较快）
		Parallelism: 1,         // 1 个并行线程
		SaltLength:  16,        // 16 字节盐值
		KeyLength:   32,        // 32 字节密钥
	}

	customHash, err := utils.GenPasswdWithParams(password, customParams)
	if err != nil {
		log.Fatalf("使用自定义参数生成哈希失败: %v", err)
	}
	fmt.Printf("自定义参数哈希: %s\n", customHash)

	// 验证自定义参数哈希
	err = utils.ComparePasswd(customHash, password)
	if err != nil {
		log.Printf("自定义参数密码验证失败: %v", err)
	} else {
		fmt.Println("✓ 自定义参数密码验证成功")
	}

	// 6. 演示相同密码产生不同哈希（由于随机盐值）
	fmt.Println("\n=== 哈希唯一性演示 ===")
	hash1, _ := utils.GenPasswd(password)
	hash2, _ := utils.GenPasswd(password)
	fmt.Printf("哈希1: %s\n", hash1)
	fmt.Printf("哈希2: %s\n", hash2)
	if hash1 != hash2 {
		fmt.Println("✓ 相同密码产生不同哈希（由于随机盐值）")
	}

	// 7. 密码强度验证示例
	fmt.Println("\n=== 密码强度验证示例 ===")
	testPasswords := []string{
		"weak",                // 太弱
		"NoNumbers!",          // 缺少数字
		"nonumbers123",        // 缺少大写字母和特殊字符
		"StrongPassword123!",  // 强密码
		"Another@Strong1Pass", // 另一个强密码
	}

	for _, testPwd := range testPasswords {
		err := utils.ValidatePasswordStrength(testPwd)
		if err != nil {
			fmt.Printf("密码 '%s': ✗ %v\n", testPwd, err)
		} else {
			fmt.Printf("密码 '%s': ✓ 强度足够\n", testPwd)
		}
	}
}

/*
在实际项目中的使用示例：

1. 用户注册时：
func RegisterUser(username, password string) error {
    // 验证密码强度
    if err := utils.ValidatePasswordStrength(password); err != nil {
        return fmt.Errorf("密码强度不足: %w", err)
    }

    // 生成哈希
    hashedPassword, err := utils.GenPasswd(password)
    if err != nil {
        return fmt.Errorf("密码哈希生成失败: %w", err)
    }

    // 保存到数据库
    user := &User{
        Username: username,
        Password: hashedPassword,
    }
    return db.Create(user).Error
}

2. 用户登录时：
func LoginUser(username, password string) (*User, error) {
    // 从数据库获取用户
    var user User
    err := db.Where("username = ?", username).First(&user).Error
    if err != nil {
        return nil, err
    }

    // 验证密码
    if err := utils.ComparePasswd(user.Password, password); err != nil {
        return nil, errors.New("用户名或密码错误")
    }

    return &user, nil
}

3. 密码修改时：
func ChangePassword(userID uint, oldPassword, newPassword string) error {
    // 获取用户
    var user User
    err := db.First(&user, userID).Error
    if err != nil {
        return err
    }

    // 验证旧密码
    if err := utils.ComparePasswd(user.Password, oldPassword); err != nil {
        return errors.New("旧密码错误")
    }

    // 验证新密码强度
    if err := utils.ValidatePasswordStrength(newPassword); err != nil {
        return fmt.Errorf("新密码强度不足: %w", err)
    }

    // 生成新哈希
    hashedPassword, err := utils.GenPasswd(newPassword)
    if err != nil {
        return fmt.Errorf("新密码哈希生成失败: %w", err)
    }

    // 更新数据库
    return db.Model(&user).Update("password", hashedPassword).Error
}
*/
