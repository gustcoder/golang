package main

import (
	"testing"
)

func TestFileExists(t *testing.T) {
	t.Run("testing a existent file", func(t *testing.T) {
		existentFile := "template_mocks/mock.html"
		actual := FileExists(existentFile)

		if !actual {
			t.Errorf("Error. Expected: true. Actual: false")
		}
	})

	t.Run("testing unexistent file", func(t *testing.T) {
		existentFile := "template_mocks/non_existent.html"
		actual := FileExists(existentFile)

		if actual {
			t.Errorf("Error. Expected: false. Actual: true")
		}
	})
}

func TestGetTemplates(t *testing.T) {
	t.Run("should return existent templates", func(t *testing.T) {
		mockTemplates := getTemplates("template_mocks/mock.html")

		if mockTemplates == nil {
			t.Error("Informed file does not exists")
		}

		expectedTemplateName := "mock.html"
		actualTemplateName := mockTemplates.Name()

		if actualTemplateName != expectedTemplateName {
			t.Errorf("Error. Actual: %s... Expected: %s", actualTemplateName, expectedTemplateName)
		}
	})
}
