package dockerhub

func validateCount(count int) error {
	if count < 1 || count > 100 {
		return ErrInvalidCount
	}
	return nil
}

func validateRepository(repository string) error {
	if repository == "" {
		return ErrEmptyRepository
	}
	if !repositoryRegex.MatchString(repository) {
		return ErrInvalidRepository
	}
	return nil
}
