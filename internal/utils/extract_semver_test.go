package utils

import "testing"

func TestExtractSemverTableDriven(t *testing.T) {
	tests := []struct {
		name        string
		version     string
		expectedSem string
		expectedErr error
	}{
		{
			name:        "empty",
			version:     "",
			expectedSem: "",
			expectedErr: ErrNoSemver,
		},
		{
			name:        "no-semver",
			version:     "latest",
			expectedSem: "",
			expectedErr: ErrNoSemver,
		},
		{
			name:        "valid-semver",
			version:     "v1.2.3",
			expectedSem: "1.2.3",
			expectedErr: nil,
		},
		{
			name:        "valid-semver-no-v",
			version:     "1.2.3",
			expectedSem: "1.2.3",
			expectedErr: nil,
		},
		{
			name:        "valid-semver-with-dashes",
			version:     "v1.2.3-rc.1",
			expectedSem: "1.2.3",
			expectedErr: nil,
		},
		{
			name:        "valid-semver-with-dashes-no-v",
			version:     "1.2.3-rc.1",
			expectedSem: "1.2.3",
			expectedErr: nil,
		},
		{
			name:        "valid-semver-with-dashes-no-v",
			version:     "1.2.3-rc.1",
			expectedSem: "1.2.3",
			expectedErr: nil,
		},
		{
			name:        "valid-semver-with-dashes-no-v",
			version:     "1.2.3-rc.1",
			expectedSem: "1.2.3",
			expectedErr: nil,
		},
		{
			name:        "valid-semver-with-dashes-no-v",
			version:     "1.2.3-rc.1",
			expectedSem: "1.2.3",
			expectedErr: nil,
		},
		{
			name:        "valid-semver-with-dashes-no-v",
			version:     "1.2.3-rc.1",
			expectedSem: "1.2.3",
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			semver, err := ExtractSemver(test.version)
			if err != test.expectedErr {
				t.Errorf("expected %v, got %v", test.expectedErr, err)
			}
			if semver != test.expectedSem {
				t.Errorf("expected %v, got %v", test.expectedSem, semver)
			}
		})
	}
}
