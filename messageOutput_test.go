package message

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockWrite implements io.Writer and is used to test messageOutput
type MockWriter struct {
	buff []byte
}
// Writes implementation
func (m *MockWriter) Write(p []byte) (n int, err error) {
	m.buff = p

	return len(p), nil
}
// Read implementation
func (m *MockWriter) Read(p []byte) (n int, err error) {
	p = m.buff[0:len(p)-1]

	return len(p), nil
}

// test Writers
func TestInfo(t *testing.T) {
	w := &MockWriter{
		buff: []byte{},
	}

	m := New(3, w, 0)
	m.Info("hi!")
	assert.Equal(t, w.buff[0:len(w.buff)-1], []byte("[INFO] [hi!]"), "")
}

func TestWarn(t *testing.T) {
	w := &MockWriter{
		buff: []byte{},
	}

	m := New(3, w, 0)
	m.Warn("hi!")
	assert.Equal(t, w.buff[0:len(w.buff)-1], []byte("[WARN] [hi!]"), "")
}

func TestError(t *testing.T) {
	w := &MockWriter{
		buff: []byte{},
	}

	m := New(3, w, 0)
	m.Error("hi!")
	assert.Equal(t, w.buff[0:len(w.buff)-1], []byte("[ERROR] [hi!]"), "")
}

func TestDebug(t *testing.T) {
	w := &MockWriter{
		buff: []byte{},
	}

	m := New(3, w, 0)
	m.Debug("hi!")
	assert.Equal(t, w.buff[0:len(w.buff)-1], []byte("[DEBUG] [hi!]"), "")
}