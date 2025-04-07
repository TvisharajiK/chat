package validate

import (
	"testing"

	"text/template"

	"path/filepath"
	"runtime"

	"os"

	"bytes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestValidateHostURL_ValidAndInvalidInputs_001(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"Valid URL", "https://example.com", "https://example.com/", false},
		{"Invalid URL - Relative", "/relative/path", "", true},
		{"Invalid URL - No Hostname", "https:///path", "", true},
		{"Invalid URL - Fragment", "https://example.com#fragment", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateHostURL(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// Test generated using Keploy
func TestValidateHostURL_ParseError_456(t *testing.T) {
	input := ":invalid url"
	_, err := ValidateHostURL(input)
	require.Error(t, err)
}

// Test generated using Keploy
func TestExecuteTemplate_NilParts_789(t *testing.T) {
	tmpl, err := template.New("main").Parse("Hello {{.Name}}")
	require.NoError(t, err)
	params := map[string]interface{}{"Name": "World"}

	content, err := ExecuteTemplate(tmpl, nil, params)
	require.NoError(t, err)
	require.Contains(t, content, "")
	assert.Equal(t, "Hello World", content[""])
}

// Test generated using Keploy
func TestExecuteTemplate_SpecificParts_101(t *testing.T) {
	tmpl, err := template.New("main").Parse(`{{define "subject"}}Subject: {{.Subject}}{{end}}{{define "body"}}Body: {{.Body}}{{end}}`)
	require.NoError(t, err)
	params := map[string]interface{}{"Subject": "Test Subject", "Body": "Test Body"}
	parts := []string{"subject", "body"}

	content, err := ExecuteTemplate(tmpl, parts, params)
	require.NoError(t, err)
	require.Len(t, content, 2)
	assert.Equal(t, "Subject: Test Subject", content["subject"])
	assert.Equal(t, "Body: Test Body", content["body"])
}

// Test generated using Keploy
func TestResolveTemplatePath_AbsolutePath_151(t *testing.T) {
	absPath := "/absolute/path/to/template.tmpl"
	if runtime.GOOS == "windows" {
		absPath = "C:\\absolute\\path\\to\\template.tmpl" // Adjust for windows if needed, though IsAbs handles both '/' and '\'
	}
	if !filepath.IsAbs(absPath) {
		t.Skip("Cannot determine an absolute path for testing environment.")
	}

	resolvedPath, err := ResolveTemplatePath(absPath)
	require.NoError(t, err)
	assert.Equal(t, absPath, resolvedPath)
}

// Test generated using Keploy
func TestResolveTemplatePath_RelativePath_161(t *testing.T) {
	relPath := "templates/../email/verify.tmpl"
	curwd, err := os.Getwd()
	require.NoError(t, err, "Failed to get current working directory")

	expectedPath := filepath.Join(curwd, "email/verify.tmpl") // Simplified by Clean
	expectedPath = filepath.Clean(expectedPath)               // Ensure OS-specific cleaning

	resolvedPath, err := ResolveTemplatePath(relPath)
	require.NoError(t, err)
	assert.Equal(t, expectedPath, resolvedPath)
}

// Test generated using Keploy
func TestReadTemplateFile_Success_171(t *testing.T) {
	tempDir := t.TempDir()
	lang := "en"
	templateFileName := "test_template.tmpl"
	templateContent := "Hello {{.Name}}"

	// Path template generates path like: /tmp/somedir/en/test_template.tmpl
	pathTempl, err := template.New("path").Parse(filepath.Join(tempDir, "{{.Language}}", templateFileName))
	require.NoError(t, err)

	// Create the directory and the actual template file
	langDir := filepath.Join(tempDir, lang)
	err = os.MkdirAll(langDir, 0755)
	require.NoError(t, err)
	templateFilePath := filepath.Join(langDir, templateFileName)
	err = os.WriteFile(templateFilePath, []byte(templateContent), 0644)
	require.NoError(t, err)

	// Execute the function under test
	parsedTempl, resolvedPath, err := ReadTemplateFile(pathTempl, lang)

	// Assertions
	require.NoError(t, err)
	assert.Equal(t, templateFilePath, resolvedPath)
	require.NotNil(t, parsedTempl)

	// Verify the parsed template works
	var buf bytes.Buffer
	err = parsedTempl.Execute(&buf, map[string]string{"Name": "Tester"})
	require.NoError(t, err)
	assert.Equal(t, "Hello Tester", buf.String())
}
