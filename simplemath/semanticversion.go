package simplemath

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

// function that takes a type as a parameter
func StringFunction(sv SemanticVersion) string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

// method on a type
func (sv SemanticVersion) StringMethod() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

// method with value receiver - returns new obj
func (sv SemanticVersion) IncrementMajor() SemanticVersion {
	sv.major += 1
	return sv
}

// method with pointer receiver - mutates passed in obj
func (sv *SemanticVersion) IncrementMinor() {
	sv.minor += 1
}

func (sv SemanticVersion) IncrementPatch() {
	sv.patch += 1
}
