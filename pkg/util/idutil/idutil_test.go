package idutil

import (
	"strings"
	"testing"

	"github.com/google/uuid"
)

// TestCreateUUID 测试 CreateUUID 函数
func TestCreateUUID(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		name     string
		prefix   string
		expected string
	}{
		{
			name:     "EmptyPrefix",
			prefix:   "",
			expected: "-" + uuid.NewString(),
		},
		{
			name:     "WithPrefix",
			prefix:   "prefix",
			expected: "prefix-" + uuid.NewString(),
		},
	}

	// 遍历测试用例并执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 调用被测试的函数
			actual := CreateUUID(tt.prefix, nil)

			// 检查前缀是否正确
			if !strings.HasPrefix(actual, tt.prefix+"-") {
				t.Errorf("CreateUUID(%s) = %s, want prefix %s-", tt.prefix, actual, tt.prefix)
			}

			// 检查 UUID 是否有效
			u, err := uuid.Parse(strings.TrimPrefix(actual, tt.prefix+"-"))
			if err != nil {
				t.Errorf("CreateUUID(%s) returned invalid UUID: %v", tt.prefix, err)
			}
			if u == uuid.Nil {
				t.Errorf("CreateUUID(%s) returned a nil UUID", tt.prefix)
			}
		})
	}
}
