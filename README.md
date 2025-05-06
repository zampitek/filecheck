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
make build
```

The executable will be in the bin/ directory.


## 🛠 Usage

```bash
filecheck scan [flags]
```

### Example

```bash 
filecheck scan ~/Downloads --extended
```

### Flags

`--extended, -e`: Generate an extended report (includes Top 5 lists)

## 📁 Report Example

```

	[LOW AGE] 	357522
	[MEDIUM AGE] 	43
	[HIGH AGE] 	817

	[LOW SIZE] 	307536
	[MEDIUM SIZE] 	66
	[HIGH SIZE] 	14


+---+------------------------------+-------------+
| # | PATH                         | AGE (days)  |
+---+------------------------------+-------------+
| 1 | /home/user/notes.txt         | 589         |
| 2 | /home/user/work.odt          | 374         |
| 3 | ...                          | ...         |
+---+------------------------------+-------------+

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