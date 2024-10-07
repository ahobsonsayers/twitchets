package main

import (
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalShorthandWithShorthandTag(t *testing.T) {
	var test struct {
		One   string `mapstructure:",shorthand"`
		Two   string
		Three string
	}
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			DecodeHook: ShorthandHookFunc(),
			Result:     &test,
		},
	)
	require.NoError(t, err)

	input := "test"
	err = decoder.Decode(input)
	require.NoError(t, err)

	require.Equal(t, input, test.One)
	require.Empty(t, test.Two)
	require.Empty(t, test.Three)
}

func TestUnmarshalShorthandWithWrongType(t *testing.T) {
	var test struct {
		One   string `mapstructure:",shorthand"`
		Two   string
		Three string
	}
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			DecodeHook: ShorthandHookFunc(),
			Result:     &test,
		},
	)
	require.NoError(t, err)

	input := 1
	err = decoder.Decode(input)
	require.Error(t, err)

	err = decoder.Decode(input)
	require.Error(t, err)
}

func TestUnmarshalFullWithShorthandTag(t *testing.T) {
	var test struct {
		One   string `mapstructure:",shorthand"`
		Two   string
		Three string
	}
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			DecodeHook: ShorthandHookFunc(),
			Result:     &test,
		},
	)
	require.NoError(t, err)

	one := "One"
	two := "Two"
	three := "Three"
	input := map[string]string{
		one:   one,
		two:   two,
		three: three,
	}
	err = decoder.Decode(input)
	require.NoError(t, err)

	require.Equal(t, one, test.One)
	require.Equal(t, two, test.Two)
	require.Equal(t, three, test.Three)
}

func TestUnmarshalFullWithMultipleShorthandTag(t *testing.T) {
	var test struct {
		One   string `mapstructure:",shorthand"`
		Two   string `mapstructure:",shorthand"`
		Three string
	}
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			DecodeHook: ShorthandHookFunc(),
			Result:     &test,
		},
	)
	require.NoError(t, err)

	one := "One"
	two := "Two"
	three := "Three"
	input := map[string]string{
		one:   one,
		two:   two,
		three: three,
	}
	err = decoder.Decode(input)
	require.NoError(t, err)

	require.Equal(t, one, test.One)
	require.Equal(t, two, test.Two)
	require.Equal(t, three, test.Three)
}

func TestUnmarshalWithoutShorthandTag(t *testing.T) {
	var test struct {
		One   string
		Two   string
		Three string
	}
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			DecodeHook: ShorthandHookFunc(),
			Result:     &test,
		},
	)
	require.NoError(t, err)

	one := "One"
	two := "Two"
	three := "Three"
	input := map[string]string{
		one:   one,
		two:   two,
		three: three,
	}
	err = decoder.Decode(input)
	require.NoError(t, err)

	require.Equal(t, one, test.One)
	require.Equal(t, two, test.Two)
	require.Equal(t, three, test.Three)
}
