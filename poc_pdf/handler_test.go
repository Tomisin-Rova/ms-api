package poc_pdf

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateFileBuffer(t *testing.T) {

	tests := []struct {
		name string
	}{
		{
			name: "Test success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, result := generateFileBuffer()
			assert.NotNil(t, result)
			assert.NoError(t, err)
		})
	}
}

func Test_getDarkGrayColor(t *testing.T) {
	tests := []struct {
		name string
		want color.Color
	}{
		{
			name: "Test success",
			want: color.Color{
				Red:   55,
				Green: 55,
				Blue:  55,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getDarkGrayColor(), "getDarkGrayColor()")
		})
	}
}
