package errors_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xaionaro-go/errors"
)

func TestWrap(t *testing.T) {
	assert.Equal(t, io.EOF, errors.Wrap(io.EOF).Err)
	assert.Equal(t, io.EOF, errors.UndefinedError.Wrap(io.EOF).Err)
	assert.Equal(t, errors.ProtocolMismatch, errors.ProtocolMismatch.Wrap(io.EOF).Err)
	assert.Equal(t, io.EOF, errors.Wrap(io.EOF, errors.ProtocolMismatch).Deepest().Err)
	assert.Equal(t, errors.ProtocolMismatch, errors.Wrap(io.EOF, errors.ProtocolMismatch).Err)
	assert.Equal(t, io.EOF, errors.Wrap(io.EOF, errors.UndefinedError).Err)

	{
		err := errors.UndefinedError.Wrap(io.EOF)
		err = errors.Wrap(err, errors.ProtocolMismatch)
		err = errors.New(`test`, `an argument`).Wrap(err)
		err.Format = `%fE`
		assert.Equal(t, `test`, err.Error())
		assert.Equal(t, errors.ProtocolMismatch, err.WrappedError.Err)
		assert.Equal(t, io.EOF, err.Deepest().Err)
	}
}
