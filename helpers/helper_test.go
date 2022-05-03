package helpers

import "testing"

func TestSuccessAPIResponse(t *testing.T) {
	data := SuccessAPIResponse("success", 200, nil)
	if data.Message != "success" || data.Code != 200 {
		t.Error("Error to generate success API response")
	}
}

func TestFailedAPIResponse(t *testing.T) {
	data := FailedAPIResponse("failed", 400)
	if data.Message != "failed" || data.Code != 400 {
		t.Error("Error to generate failed API response")
	}
}

func TestReadJsonFile(t *testing.T) {
	_, err := ReadJsonFile("../data/categories.json")
	if err != nil {
		t.Errorf("Error read json file : %s", err.Error())
	}
}

func TestQueryToUint(t *testing.T) {
	testUint := QueryToUint("1")
	if testUint != 1 {
		t.Error("Error convert string to uint")
	}
}
