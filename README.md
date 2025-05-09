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

**More features coming soon!**

## 📦 Installation

Clone the repo and build with Go:

```bash
git clone https://github.com/zampitek/filecheck.git
cd filecheck
make build 		// build for all platforms
```
or, if you want to build for a specific system
```bash
make [system]   // linux, mac, windows
```

The executable will be in the bin/ directory.


## 🛠 Usage

```bash
filecheck scan [directory] [flags]
```

### Example

```bash 
filecheck scan ~/Downloads --checks=age
```

### Flags

`--checks=[checks]`: Specify what checks to perform

## 📁 Report Example

```
AGE GROUP SUMMARY
+--------+------------+------------+
| GROUP  | FILE COUNT | TOTAL SIZE |
+--------+------------+------------+
| LOW    |     353278 | 37.96 GB   |
| MEDIUM |       3638 | 1.12 GB    |
| HIGH   |        850 | 0.35 GB    |
+--------+------------+------------+
```

## 🧪 Development

Format code:
```bash
make fmt
```

Run build:
```bash
make run
```

Clean output:
```bash
make clean
```

## 🧾 License
This project is licensed under the MIT License. See [LICENSE](./LICENSE) for details.

## 🙋‍♂️ Contributing
Contributions, suggestions, and issues are welcome! Feel free to fork the repo or open a PR.