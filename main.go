package main

// Goals:

// Provide file or folder as source
// Provide folder as destination (can be same folder as source)

// Using defined format, query latest entry in destination
// Compare file (or recursively all files) from source with that of latest in destination
// If file matches latest already, no changes have been made. Exit happily.

// Using defined format, query highest version number in destination
// If source is also a version, assume we're rolling back and strip version number out as per format
// Copy over source to destination with name that matches next version index of format
// Remove latest, and link new version to latest

func main() {

}
