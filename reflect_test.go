package dockergen

import "testing"

func TestDeepGetNoPath(t *testing.T) {
	item := RuntimeContainer{}
	value := deepGet(item, "")
	if _, ok := value.(RuntimeContainer); !ok {
		t.Fail()
	}

	var returned RuntimeContainer
	returned = value.(RuntimeContainer)
	if !returned.Equals(item) {
		t.Fail()
	}
}

func TestDeepGetSimple(t *testing.T) {
	item := RuntimeContainer{
		ID: "expected",
	}
	value := deepGet(item, "ID")
	if _, ok := value.(string); !ok {
		t.Errorf("expected: %#v. got: %#v", "expected", value)
	}

	if value != "expected" {
		t.Errorf("expected: %s. got: %s", "expected", value)
	}
}

func TestDeepGetSimpleDotPrefix(t *testing.T) {
	item := RuntimeContainer{
		ID: "expected",
	}
	value := deepGet(item, "...ID")
	if _, ok := value.(string); !ok {
		t.Errorf("expected: %#v. got: %#v", "expected", value)
	}

	if value != "expected" {
		t.Errorf("expected: %s. got: %s", "expected", value)
	}
}

func TestDeepGetMap(t *testing.T) {
	item := RuntimeContainer{
		Env: map[string]string{
			"key": "value",
		},
	}
	value := deepGet(item, "Env.key")
	if _, ok := value.(string); !ok {
		t.Errorf("expected: %#v. got: %#v", "value", value)
	}

	if value != "value" {
		t.Errorf("expected: %s. got: %s", "value", value)
	}
}

func TestDeepGetMapWithDots(t *testing.T) {
	item := RuntimeContainer{
		Labels: map[string]string{
			"some.dot.value": "value",
		},
	}
	value := deepGet(item, "Labels.some\\.dot\\.value")
	if _, ok := value.(string); !ok {
		t.Errorf("expected: %#v. got: %#v", "value", value)
	}

	if value != "value" {
		t.Errorf("expected: %s. got: %s", "value", value)
	}
}

func TestDeepGetMapWithDotsThenMore(t *testing.T) {
	item := RuntimeContainer{
		Volumes: map[string]Volume{
			"some.dot.value": Volume{
				Path: "value",
			},
		},
	}
	value := deepGet(item, "Volumes.some\\.dot\\.value.Path")
	if _, ok := value.(string); !ok {
		t.Errorf("expected: %#v. got: %#v", "value", value)
	}

	if value != "value" {
		t.Errorf("expected: %s. got: %s", "value", value)
	}
}
