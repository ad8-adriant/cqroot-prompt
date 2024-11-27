package write_test

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/ad8-adriant/cqroot-prompt/write"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"
)

func TestWithTeaProgramOpts(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	withInput := tea.WithInput(&in)
	withOutput := tea.WithOutput(&out)

	model := write.New(
		"",
		write.WithTeaProgramOpts(withInput, withOutput),
	)

	require.True(t, reflect.ValueOf(withInput) == reflect.ValueOf(model.TeaProgramOpts()[0]))
	require.True(t, reflect.ValueOf(withOutput) == reflect.ValueOf(model.TeaProgramOpts()[1]))
}

func TestWithValidateFunc(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	validateErr := errors.New("validation error")
	validateFunc := func(s string) error {
		if s != "test" {
			return validateErr
		}
		return nil
	}

	in.Write([]byte{byte(tea.KeyCtrlD)})

	model := write.New("", write.WithValidateFunc(validateFunc))

	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(write.Model)
	require.True(t, ok)

	require.Equal(t, m.Error(), validateErr)
}

func TestDefaultValueWithValidateFunc(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	validateErr := errors.New("validation error")
	validateFunc := func(s string) error {
		if s != "test" {
			return validateErr
		}
		return nil
	}

	in.Write([]byte{byte(tea.KeyCtrlD)})

	model := write.New("test", write.WithValidateFunc(validateFunc))

	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(write.Model)
	require.True(t, ok)

	require.Nil(t, m.Error())
}
