package dockerhub

import "testing"

func TestValidateRepositoryTableDriven(t *testing.T) {
	tests := []struct {
		name        string
		repository  string
		expectedErr error
	}{
		{
			name:        "empty",
			repository:  "",
			expectedErr: ErrEmptyRepository,
		},
		{
			name:        "trailing-slash",
			repository:  "invalid/",
			expectedErr: ErrInvalidRepository,
		},
		{
			name:        "leading-slash",
			repository:  "/invalid",
			expectedErr: ErrInvalidRepository,
		},
		{
			name:        "valid-repository",
			repository:  "zaptross/godohuver",
			expectedErr: nil,
		},
		{
			name:        "valid-official-underscore",
			repository:  "_/golang",
			expectedErr: nil,
		},
		{
			name:        "valid-official",
			repository:  "golang",
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateRepository(test.repository)
			if err != test.expectedErr {
				t.Errorf("expected %v, got %v", test.expectedErr, err)
			}
		})
	}
}

func TestValidateCountTableDriven(t *testing.T) {
	tests := []struct {
		name        string
		count       int
		expectedErr error
	}{
		{
			name:        "zero",
			count:       0,
			expectedErr: ErrInvalidCount,
		},
		{
			name:        "negative",
			count:       -1,
			expectedErr: ErrInvalidCount,
		},
		{
			name:        "too-large",
			count:       101,
			expectedErr: ErrInvalidCount,
		},
		{
			name:        "valid",
			count:       1,
			expectedErr: nil,
		},
		{
			name:        "valid",
			count:       100,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateCount(test.count)
			if err != test.expectedErr {
				t.Errorf("expected %v, got %v", test.expectedErr, err)
			}
		})
	}
}
