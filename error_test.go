package errors_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xaionaro-go/errors"
)

func TestWrap(t *testing.T) {
	assert.Equal(t, io.EOF, errors.Wrap(io.EOF).GetErr())
	assert.Equal(t, io.EOF, errors.UndefinedError.Wrap(io.EOF).GetErr())
	assert.Equal(t, errors.ProtocolMismatch, errors.ProtocolMismatch.Wrap(io.EOF).GetErr())
	assert.Equal(t, io.EOF, errors.Wrap(io.EOF, errors.ProtocolMismatch).Deepest().GetErr())
	assert.Equal(t, errors.ProtocolMismatch, errors.Wrap(io.EOF, errors.ProtocolMismatch).GetErr())
	assert.Equal(t, io.EOF, errors.Wrap(io.EOF, errors.UndefinedError).GetErr())

	{
		err := errors.UndefinedError.Wrap(io.EOF)
		err = errors.Wrap(err, errors.ProtocolMismatch)
		err = errors.New(`test`, `an argument`).Wrap(err)
		err.SetFormat(`%fE`)
		assert.Equal(t, `test`, err.Error())
		assert.Equal(t, errors.ProtocolMismatch, err.GetWrappedError().GetErr())
		assert.Equal(t, io.EOF, err.Deepest().GetErr())
	}

	assert.Equal(t, nil, errors.Wrap(nil))
	assert.False(t, strings.Index(errors.Wrap(io.EOF).Error(), "error.go") >= 0)
}
