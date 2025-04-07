package auth

import (
	"testing"

	"errors"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestLevelString_ValidLevels_123(t *testing.T) {
	tests := []struct {
		level    Level
		expected string
	}{
		{LevelNone, ""},
		{LevelAnon, "anon"},
		{LevelAuth, "auth"},
		{LevelRoot, "root"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.level.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test generated using Keploy
func TestParseAuthLevel_ValidAndInvalidInputs_456(t *testing.T) {
	tests := []struct {
		input    string
		expected Level
	}{
		{"anon", LevelAnon},
		{"ANON", LevelAnon},
		{"auth", LevelAuth},
		{"AUTH", LevelAuth},
		{"root", LevelRoot},
		{"ROOT", LevelRoot},
		{"unknown", LevelNone},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ParseAuthLevel(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test generated using Keploy
func TestLevelMarshalText_ValidAndInvalidLevels_789(t *testing.T) {
	tests := []struct {
		level    Level
		expected []byte
		err      error
	}{
		{LevelNone, []byte(""), nil},
		{LevelAnon, []byte("anon"), nil},
		{LevelAuth, []byte("auth"), nil},
		{LevelRoot, []byte("root"), nil},
		{Level(999), nil, errors.New("auth.Level: invalid level value")},
	}

	for _, tt := range tests {
		t.Run(string(tt.expected), func(t *testing.T) {
			result, err := tt.level.MarshalText()
			if tt.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// Test generated using Keploy
func TestLevelUnmarshalText_ValidAndInvalidInputs_321(t *testing.T) {
	tests := []struct {
		input    []byte
		expected Level
		err      error
	}{
		{[]byte(""), LevelNone, nil},
		{[]byte("anon"), LevelAnon, nil},
		{[]byte("ANON"), LevelAnon, nil},
		{[]byte("auth"), LevelAuth, nil},
		{[]byte("AUTH"), LevelAuth, nil},
		{[]byte("root"), LevelRoot, nil},
		{[]byte("ROOT"), LevelRoot, nil},
		{[]byte("unknown"), LevelNone, errors.New("auth.Level: unrecognized")},
	}

	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			var level Level
			err := level.UnmarshalText(tt.input)
			if tt.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, level)
			}
		})
	}
}

// Test generated using Keploy
func TestLevelMarshalJSON_ValidAndInvalidLevels_654(t *testing.T) {
	tests := []struct {
		level    Level
		expected []byte
		err      error
	}{
		{LevelNone, []byte(`""`), nil},
		{LevelAnon, []byte(`"anon"`), nil},
		{LevelAuth, []byte(`"auth"`), nil},
		{LevelRoot, []byte(`"root"`), nil},
		{Level(999), nil, errors.New("auth.Level: invalid level value")},
	}

	for _, tt := range tests {
		t.Run(string(tt.expected), func(t *testing.T) {
			result, err := tt.level.MarshalJSON()
			if tt.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// Test generated using Keploy
func TestLevelUnmarshalJSON_ValidAndInvalidInputs_987(t *testing.T) {
	tests := []struct {
		input    []byte
		expected Level
		err      error
	}{
		{[]byte(`""`), LevelNone, nil},
		{[]byte(`"anon"`), LevelAnon, nil},
		{[]byte(`"auth"`), LevelAuth, nil},
		{[]byte(`"root"`), LevelRoot, nil},
		{[]byte(`"unknown"`), LevelNone, errors.New("auth.Level: unrecognized")},
		{[]byte(`invalid`), LevelNone, errors.New("syntax error")},
	}

	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			var level Level
			err := level.UnmarshalJSON(tt.input)
			if tt.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, level)
			}
		})
	}
}

// Test generated using Keploy
func TestFeatureUnmarshalText_ValidAndInvalidInputs_222(t *testing.T) {
	tests := []struct {
		input    []byte
		expected Feature
		err      error
	}{
		{[]byte("V"), FeatureValidated, nil},
		{[]byte("L"), FeatureNoLogin, nil},
		{[]byte("VL"), FeatureValidated | FeatureNoLogin, nil},
		{[]byte("invalid"), Feature(0), errors.New("Feature: invalid character 'i'")},
	}

	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			var feature Feature
			err := feature.UnmarshalText(tt.input)
			if tt.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, feature)
			}
		})
	}
}

// Test generated using Keploy
func TestLevelString_InvalidLevel_001(t *testing.T) {
	invalidLevel := Level(999) // An invalid level
	expected := "unkn"
	result := invalidLevel.String()
	assert.Equal(t, expected, result)
}

// Test generated using Keploy
func TestFeatureMarshalText_ValidFeatures_111(t *testing.T) {
	tests := []struct {
		name     string
		feature  Feature
		expected []byte
	}{
		{"Validated", FeatureValidated, []byte("V")},
		{"NoLogin", FeatureNoLogin, []byte("L")},
		{"ValidatedAndNoLogin", FeatureValidated | FeatureNoLogin, []byte("VL")},
		{"None", Feature(0), []byte("")},
		{"ArbitraryValidBits", Feature(FeatureValidated | FeatureNoLogin | 4 | 8), []byte("VL")}, // Other bits ignored
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.feature.MarshalText()
			assert.NoError(t, err) // Feature.MarshalText should not return error
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test generated using Keploy
func TestFeatureUnmarshalText_ExpandedCases_223(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected Feature
		err      error
	}{
		{"Empty", []byte(""), Feature(0), nil},
		{"LowercaseV", []byte("v"), FeatureValidated, nil},
		{"LowercaseL", []byte("l"), FeatureNoLogin, nil},
		{"LowercaseVL", []byte("vl"), FeatureValidated | FeatureNoLogin, nil},
		{"MixedCaseVL", []byte("Vl"), FeatureValidated | FeatureNoLogin, nil},
		{"NumericValidated", []byte("1"), FeatureValidated, nil},             // FeatureValidated = 1
		{"NumericNoLogin", []byte("2"), FeatureNoLogin, nil},                 // FeatureNoLogin = 2
		{"NumericBoth", []byte("3"), FeatureValidated | FeatureNoLogin, nil}, // 1 | 2 = 3
		{"NumericInvalidString", []byte("1a"), Feature(0), errors.New("strconv.Atoi: parsing \"1a\": invalid syntax")},
		{"InvalidCharMiddle", []byte("VX"), FeatureValidated, errors.New("Feature: invalid character 'X'")},
		{"InvalidCharStart", []byte("XV"), Feature(0), errors.New("Feature: invalid character 'X'")},
		{"ValidExisting", []byte("V"), FeatureValidated, nil},                                            // From original test
		{"InvalidExisting", []byte("invalid"), Feature(0), errors.New("Feature: invalid character 'i'")}, // From original test
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var feature Feature
			err := feature.UnmarshalText(tt.input)
			if tt.err != nil {
				assert.Error(t, err)
				// Use Contains because strconv error messages might differ slightly across Go versions/OS
				assert.Contains(t, err.Error(), tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
			// Assert expected feature value even if there was an error (it should be reset or partially set)
			assert.Equal(t, tt.expected, feature)
		})
	}
}

// Test generated using Keploy
func TestFeatureString_ValidFeatures_555(t *testing.T) {
	tests := []struct {
		name     string
		feature  Feature
		expected string
	}{
		{"Validated", FeatureValidated, "V"},
		{"NoLogin", FeatureNoLogin, "L"},
		{"ValidatedAndNoLogin", FeatureValidated | FeatureNoLogin, "VL"},
		{"None", Feature(0), ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.feature.String()
			assert.Equal(t, tt.expected, result) // String() should not return error for Feature
		})
	}
}

// Test generated using Keploy
func TestFeatureMarshalJSON_ValidFeatures_666(t *testing.T) {
	tests := []struct {
		name     string
		feature  Feature
		expected []byte
	}{
		{"Validated", FeatureValidated, []byte(`"V"`)},
		{"NoLogin", FeatureNoLogin, []byte(`"L"`)},
		{"ValidatedAndNoLogin", FeatureValidated | FeatureNoLogin, []byte(`"VL"`)},
		{"None", Feature(0), []byte(`""`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.feature.MarshalJSON()
			assert.NoError(t, err) // MarshalJSON should not return error for Feature
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test generated using Keploy
func TestFeatureUnmarshalJSON_QuotedString_333(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected Feature
		err      error
	}{
		{"QuotedV", []byte(`"V"`), FeatureValidated, nil},
		{"QuotedL", []byte(`"L"`), FeatureNoLogin, nil},
		{"QuotedVL", []byte(`"VL"`), FeatureValidated | FeatureNoLogin, nil},
		{"QuotedEmpty", []byte(`""`), Feature(0), nil},
		{"QuotedInvalidChar", []byte(`"VX"`), FeatureValidated, errors.New("Feature: invalid character 'X'")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var feature Feature
			err := feature.UnmarshalJSON(tt.input)
			if tt.err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, feature)
		})
	}
}
