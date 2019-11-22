package json

import "testing"

func TestDotStringToJSON(t *testing.T) {
	t.Log(dotStringToJSON("a.b.c.d.e", "10", map[string]interface{}{}))
}
