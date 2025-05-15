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
make build 		# build for all platforms
```
or, if you want to build for a specific system
```bash
make [system]   # linux, mac, windows
```

The executable will be in the bin/ directory.


## 🛠 Usage

```bash
filecheck scan [directory] [flags]
```

### Example

```bash 
filecheck scan /home/user --checks=size --size-top=3
```

### Flags

`--checks=[checks]`: Specify what checks to perform

## 📁 Report Example

```
==================================================
		FILE ANALYSIS REPORT
==================================================

###################
# BY FILE SIZE    #
###################

--- SIZE GROUP SUMMARY ---
  LOW (modified in last 90 days):          331804 files | 13.96 GB
  MEDIUM (modified 90-180 days ago):           50 files | 11.50 GB
  HIGH (modified pver 180 days ago):            9 files | 16.07 GB
--------------------------------------------------

[ LOW ] - Files under 100 MB
  Top 3:
    1. /home/user/sdk/android/ndk/26.3.11579264/toolchains/llvm/prebuilt/linux-x86_64/lib/libLTO.so.17                  90 days ago |  90.22 MB
    2. /home/user/flutter/project/build/app/outputs/flutter-apk/app-debug.apk                                           90 days ago |  81.13 MB
    3. /home/user/flutter/project/build/app/outputs/apk/debug/app-debug.apk                                             90 days ago |  81.13 MB


[ MEDIUM ] - Files between 100 MB and 1 GB
  Top 3:
    1. /home/user/folder/whatsapp-video_1                                                                               72 days ago | 926.23 MB
    2. /home/user/.android/avd/Medium_Phone.avd/sdcard.img                                                              90 days ago | 800.00 MB
    3. /home/user/folder/whatsapp-video_2                                                                               72 days ago | 789.59 MB


[ HIGH ] - Files over 1 GB
  Top 3:
    1. /home/user/Downloads/Win11.iso                                                                                    51 days ago |   5.34 GB
    2. /home/user/sdk/android/system-images/android-35/google_apis_playstore/x86_64/system.img                           90 days ago |   2.01 GB
    3. /home/user/folder/whatsapp-video_3                                                                                72 days ago |   1.69 GB

```

## 🧪 Development

Format code:
```bash
make fmt
```

Clean output:
```bash
make clean
```

## 🧾 License
This project is licensed under the MIT License. See [LICENSE](./LICENSE) for details.

## 🙋‍♂️ Contributing
Contributions, suggestions, and issues are welcome! Feel free to fork the repo or open a PR.