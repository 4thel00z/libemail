package gmail

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	src := `<!DOCTYPE html>
<html>
<body>

<h1>{{index . "Header"}}</h1>

<p>{{index . "Paragraph"}}</p>

</body>
</html>`
	expected := `<!DOCTYPE html>
<html>
<body>

<h1>Some Header</h1>

<p>Some Paragraph</p>

</body>
</html>`
	m := map[string]string{
		"Header":    "Some Header",
		"Paragraph": "Some Paragraph",
	}
	out, err := RenderTemplate("test", src, m)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, out)
}
