package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		program string
		want    *Program
	}{
		{
			program: "some stuff without a comment.",
			want:    &Program{},
		},
		{
			program: "# only a comment.",
			want: &Program{
				[]*Comment{
					&Comment{"# only a comment."},
				},
			},
		},
		{
			program: `
# Comment 1, just a line of stuff.

echo "abc the thing"

function dostuff() {
    echo 'doing stuff'
}

a=$(ls)  #list the files.
`,
			want: &Program{
				[]*Comment{
					&Comment{"# Comment 1, just a line of stuff."},
					&Comment{"#list the files."},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.program, func(t *testing.T) {
			assert := assert.New(t)

			got, gotErr := Parse("test", []byte(test.program))

			assert.NoError(gotErr)
			assert.Equal(test.want, got)
		})
	}
}
