package libs

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDeleteEmpty(t *testing.T) {
	expectedItems := 3
	data := []string{
		"",
		"item",
		"item2",
		"",
		"item3",
	}
	if result := DeleteEmpty(data); len(result) != expectedItems {
		t.Errorf("Expected itemes is not equal %d", expectedItems)
	}
}

func TestFind(t *testing.T) {
	data := []string{
		"item",
		"item2",
		"item3",
	}
	index, found := Find(data, "item2")
	if index != 1 {
		t.Errorf("assert item2 has index 1, but got index is %d", index)
	}

	if found != true {
		t.Error("found is not true value")
	}

	index, found = Find(data, "item4")
	if index != -1 {
		t.Error("item4 is not exists, but has index not equal -1")
	}

	if found != false {
		t.Error("item4 is not exists, and found must be is false, true got")
	}
}

func getRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func TestInsertToSlice(t *testing.T) {
	index := getRandomInt(0, 3)
	data := []string{"item1", "item3", "item4"}
	result := InsertToSlice(data, "insertingItem", index)

	if result[index] != "insertingItem" {
		t.Errorf("insertingItem has not %d index", index)
	}

	for i, v := range data {
		if i < index {
			if result[i] != v {
				t.Errorf("item %s must have index %d in result", v, i)
			}
		}
		if i >= index {
			if result[i+1] != v {
				t.Errorf("item %s must have index %d in result", v, i)
			}
		}
	}
}

func TestGetPwd(t *testing.T) {
	d, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can't get dir. error: %v\n", err)
	}
	if d != GetPwd() {
		t.Errorf("Get pwd is not same this: %s", d)
	}
}

func TestFileExists(t *testing.T) {
	fpath := GetPwd() + "/testdata/test_file.txt"
	exist := FileExists(fpath)
	if exist == false {
		t.Errorf("file %s is not exists", fpath)
	}
	fpath = GetPwd() + "/not_exist.txt"
	if FileExists(fpath) != false {
		t.Error("should be false when file not exist")
	}
}

func TestReplaceEnvVariables(t *testing.T) {
	setEnvVar("TEST_VAR", "TEST_VAL", t)
	setEnvVar("TEST_VAR2", "TEST_VAL2", t)
	setEnvVar("TEST_VAR3", "TEST_VAL3", t)

	data := []string{"some-string{TEST_VAR}for change", "second{TEST_VAR2}string{TEST_VAR3}"}
	expected := "some-stringTEST_VALfor change secondTEST_VAL2stringTEST_VAL3"
	if result := ReplaceEnvVariables(data); strings.Join(result, " ") != expected {
		t.Errorf("error while replace env, excpected: %q, got %q", expected, strings.Join(result, " "))
	}
}

func TestReplaceEnv(t *testing.T) {
	setEnvVar("TEST_VAR", "TEST_VAL", t)
	target := "some-string{TEST_VAR}for change"
	expected := "some-stringTEST_VALfor change"
	if result := replaceEnv(target, "{TEST_VAR}"); result != expected {
		t.Errorf("error while replace env var, expected: %q, got %q", expected, result)
	}

	target = "no vars"
	expected = "no vars"
	if result := replaceEnv(target, "{NOT_EXISTS_VAR}"); result != expected {
		t.Errorf("error while replace env var, expected: %q, got %q", expected, result)
	}
}

func TestMergeSliceOfString(t *testing.T) {
	data := []string{
		"data", "", "param",
		"data2", "",
		"data3", "param3", "",
	}
	e := "data param data2 data3 param3"
	r := MergeSliceOfString(data)
	if len(r) != 5 {
		t.Errorf("slice must be length 5")
	}

	if e != strings.Join(r, " ") {
		t.Errorf("error merge expect %q\n got %q", e, strings.Join(r, " "))
	}
}

func TestIndexExists(t *testing.T) {
	data := []string{
		"data",
		"data2",
	}
	if isSet := IndexExists(data, 1); isSet != true {
		t.Error("index must be exists")
	}
	if isSet := IndexExists(data, 3); isSet != false {
		t.Error("index must be not exists")
	}
}

func setEnvVar(name, value string, t *testing.T) {
	if err := os.Setenv(name, value); err != nil {
		t.Error("error while setting env variable")
	}
}
