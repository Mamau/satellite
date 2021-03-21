package libs

import (
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestGetClientConfig(t *testing.T) {
	fp := GetPwd() + "/testdata/starter"
	result := GetClientConfig(fp)

	if result != fp+".yaml" {
		t.Errorf("file %s is not exist", fp)
	}

	fp = GetPwd() + "/testdata/starter_not_exists"
	result = GetClientConfig(fp)
	if result != "" {
		t.Errorf("file %s not exists and return non empty string", fp)
	}
}

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
	fpath := GetPwd() + "/testdata/testFile.txt"
	exist := FileExists(fpath)
	if exist == false {
		t.Errorf("file %s is not exists", fpath)
	}
}
