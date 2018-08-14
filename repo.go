package main

// Store and query version information from repository
type Repository interface {
	SetAsLatest(*FormatterAll) (*FormatterAll, error)
}

// Store files in folder versioned via named files.
type DirectoryRepo struct {
	source string
}

// Create new DirectoryRepo instance
func NewDirectoryRepo(dir string) *DirectoryRepo {
	return &DirectoryRepo{source: dir}
}

// Set the provided version as the latest in repo
func (self *DirectoryRepo) SetAsLatest(*FormatterAll) (*FormatterAll, error) {
	return new(FormatterAll), nil
}
