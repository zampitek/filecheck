# FileCheck

**FileCheck** is a command-line utility for analyzing the state of a given directory. It generates structured reports on file age, size, and other criteria, helping users clean up, audit, or automate maintenance tasks.

## 🚀 Features

- Scan directories and classify files by:
  - Last access date (age-based severity levels)
  - File size (low, medium, high thresholds)
- Generate reports summarizing:
  - File count per category
  - Top N oldest or heaviest files
- CLI interface with configurable flags
- YAML-based rule system for automation (WIP)
- Fast execution via Go concurrency

## 📦 Installation

Clone the repo and build with Go:

```bash
git clone https://github.com/zampitek/filecheck.git
cd filecheck
make build
