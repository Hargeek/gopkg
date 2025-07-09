# 密码加密工具 (Argon2id)

本包提供了基于 Argon2id 算法的安全密码哈希功能，用于替代传统的 bcrypt 算法。Argon2id 是当前推荐的密码哈希算法，提供更强的安全性。

## 功能特性

- **安全性强**: 使用 Argon2id 算法，抗彩虹表、字典攻击和暴力破解
- **防时序攻击**: 使用常数时间比较防止时序攻击
- **参数可配置**: 支持自定义内存使用、迭代次数等参数
- **密码强度验证**: 内置密码复杂度检查
- **兼容性好**: 标准格式输出，便于存储和迁移

## 主要函数

### GenPasswd(passwd string) (string, error)
使用默认参数生成密码哈希。

### ComparePasswd(hashPasswd, passwd string) error
验证密码是否匹配存储的哈希。

### ValidatePasswordStrength(passwd string) error
验证密码强度是否符合要求。

### GenPasswdWithParams(passwd string, params *Argon2Params) (string, error)
使用自定义参数生成密码哈希。

## 默认参数

```go
var DefaultArgon2Params = &Argon2Params{
    Memory:      64 * 1024, // 64 MB 内存使用
    Iterations:  3,         // 3 次迭代
    Parallelism: 2,         // 2 个并行线程
    SaltLength:  16,        // 16 字节盐值
    KeyLength:   32,        // 32 字节密钥
}
```

## 快速开始

```go
import "gitlab.bee.to/sre/opsis-platform/opsis-account-service/pkg/utils"

// 1. 验证密码强度
if err := utils.ValidatePasswordStrength(password); err != nil {
    return fmt.Errorf("密码强度不足: %w", err)
}

// 2. 生成密码哈希
hashedPassword, err := utils.GenPasswd(password)
if err != nil {
    return fmt.Errorf("密码哈希生成失败: %w", err)
}

// 3. 验证密码
if err := utils.ComparePasswd(hashedPassword, inputPassword); err != nil {
    return errors.New("密码不匹配")
}
```

## 密码强度要求

- 最少 8 个字符
- 包含至少一个大写字母
- 包含至少一个小写字母
- 包含至少一个数字
- 包含至少一个特殊字符 (!@#$%^&*()_+-=[]{}|;:,.<>?)

## 哈希格式

生成的哈希采用标准格式：
```
$argon2id$v=19$m=65536,t=3,p=2$base64_salt$base64_hash
```

其中：
- `argon2id`: 算法标识
- `v=19`: Argon2 版本号
- `m=65536`: 内存使用量 (KB)
- `t=3`: 迭代次数
- `p=2`: 并行度
- `base64_salt`: Base64 编码的盐值
- `base64_hash`: Base64 编码的哈希值

## 性能考虑

在 Apple M4 处理器上的性能测试结果：

```
BenchmarkGenPasswd-10         19    66309877 ns/op    67117916 B/op    75 allocs/op
BenchmarkComparePasswd-10     18    61365509 ns/op    67117450 B/op    75 allocs/op
```

- 生成哈希约需 66ms
- 验证密码约需 61ms

## 安全建议

1. **参数调优**: 根据服务器性能调整参数，确保响应时间在可接受范围内
2. **定期更新**: 定期重新评估参数设置，随着硬件升级适当增加安全强度
3. **错误处理**: 不要在错误信息中暴露具体的失败原因
4. **限制重试**: 实施登录失败次数限制，防止暴力破解

## 从 bcrypt 迁移

如果您需要从 bcrypt 迁移到 Argon2id，建议采用以下策略：

1. 新用户直接使用 Argon2id
2. 现有用户登录时验证 bcrypt，如果成功则更新为 Argon2id
3. 逐步淘汰 bcrypt 支持

## 示例

详细使用示例请参考 `examples/argon2id_usage.go` 文件。

## 测试

运行测试：
```bash
go test ./pkg/utils -v
```

运行性能测试：
```bash
go test ./pkg/utils -bench=. -benchmem
``` 