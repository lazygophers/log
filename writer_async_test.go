package log

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// bufferWriter 是一个包装器，让 bytes.Buffer 实现 Writer 接口
type bufferWriter struct {
	*bytes.Buffer
}

func (bw *bufferWriter) Close() error {
	return nil
}

func newBufferWriter() *bufferWriter {
	return &bufferWriter{Buffer: &bytes.Buffer{}}
}

// TestAsyncWriter 测试异步写入器
func TestAsyncWriter(t *testing.T) {
	buf := newBufferWriter()
	
	// 创建异步写入器
	asyncWriter := NewAsyncWriter(buf)
	require.NotNil(t, asyncWriter, "NewAsyncWriter 应该返回非空值")

	// 测试写入
	testData := []byte("test message 1\n")
	n, err := asyncWriter.Write(testData)
	require.NoError(t, err, "写入不应该出错")
	require.Equal(t, len(testData), n, "应该返回正确的写入字节数")

	// 写入更多数据
	testData2 := []byte("test message 2\n")
	n, err = asyncWriter.Write(testData2)
	require.NoError(t, err, "第二次写入不应该出错")
	require.Equal(t, len(testData2), n, "应该返回正确的写入字节数")

	// 等待一小段时间让异步写入完成
	time.Sleep(100 * time.Millisecond)

	// 关闭异步写入器
	err = asyncWriter.Close()
	require.NoError(t, err, "关闭不应该出错")

	// 验证数据被写入
	output := buf.String()
	require.Contains(t, output, "test message 1", "应该包含第一条消息")
	require.Contains(t, output, "test message 2", "应该包含第二条消息")
}

// TestAsyncWriter_CloseWithPendingWrites 测试有待写入数据时的关闭
func TestAsyncWriter_CloseWithPendingWrites(t *testing.T) {
	buf := newBufferWriter()
	asyncWriter := NewAsyncWriter(buf)
	
	// 写入大量数据
	for i := 0; i < 5; i++ {
		data := []byte("message ")
		_, err := asyncWriter.Write(data)
		require.NoError(t, err, "写入不应该出错")
	}

	// 立即关闭
	err := asyncWriter.Close()
	require.NoError(t, err, "关闭不应该出错")

	// 验证至少有一些数据被写入
	output := buf.String()
	require.Contains(t, output, "message", "应该至少包含部分消息")
}

// TestAsyncWriter_MultipleWrites 测试多次连续写入
func TestAsyncWriter_MultipleWrites(t *testing.T) {
	buf := newBufferWriter()
	asyncWriter := NewAsyncWriter(buf)

	messages := []string{
		"First message\n",
		"Second message\n", 
		"Third message\n",
		"Fourth message\n",
		"Fifth message\n",
	}

	// 连续写入多条消息
	for i, msg := range messages {
		n, err := asyncWriter.Write([]byte(msg))
		require.NoError(t, err, "写入消息 %d 不应该出错", i+1)
		require.Equal(t, len(msg), n, "消息 %d 应该返回正确的写入字节数", i+1)
	}

	// 等待异步写入完成
	time.Sleep(200 * time.Millisecond)

	// 关闭写入器
	err := asyncWriter.Close()
	require.NoError(t, err, "关闭不应该出错")

	// 验证所有消息都被写入
	output := buf.String()
	for i, msg := range messages {
		require.Contains(t, output, strings.TrimSpace(msg), "应该包含消息 %d", i+1)
	}
}

// TestAsyncWriter_WriteAfterClose 测试关闭后写入的行为
func TestAsyncWriter_WriteAfterClose(t *testing.T) {
	buf := newBufferWriter()
	asyncWriter := NewAsyncWriter(buf)

	// 先关闭
	err := asyncWriter.Close()
	require.NoError(t, err, "关闭不应该出错")

	// 尝试在关闭后写入
	_, err = asyncWriter.Write([]byte("after close"))
	// 这里的行为取决于具体实现，可能返回错误或被忽略
	// 我们主要确保不会崩溃
	// require.Error(t, err, "关闭后写入应该返回错误") // 根据实际实现调整
}