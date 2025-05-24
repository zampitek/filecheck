// Package checks provides utilities for analyzing and categorizing files
// based on various characteristics such as age, size, and potentially others.
//
// Each function in this package typically accepts a list of FileInfo objects
// and returns categorized results. For example, CheckAge separates files into
// three groups based on how recently they were last accessed.
//
// This package is designed to work in conjunction with internal.Scan and
// report generation logic found in the report package.
package checks
