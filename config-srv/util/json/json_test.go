package json

import "testing"

func TestDotStringToJSON(t *testing.T) {
	t.Log(DotStringToJSON("a.b.c.d.e", "10", map[string]interface{}{}))
}
