# Exercise 1A: Find Stuff

## Goals

- Familiarize yourself with the golang.org website and resources

## Setup

- Visit `golang.org`

## Directions

Answer the following questions

1. Read about `for loops` in the _Effective Go_ document

- What kind of loop doesn’t exist in Go? do..while
Common for loop methods for-statement, do-while statement, labeled statement, break statement, continue statement, for-in statement, for-of statement/ in common programming language for-loops separate into three method, for, while, and do while.

2. Read about the `fmt` _package_

- What does `fmt.Println()` return? number of bytes written and any write error encountered. When we talk about errors and kind of what some of these functions return, its kind thinking of console.log, we would expect it to return a string.

3. Find a _blog post_ about the recent release of Go 1.13

- What are some of the new features?
Go 1.13 supports a more uniform and modernized set of number literal prefixes, checksum database (A checksum database is a stype of database that stores checksums for files or data. A checksum is a value that is computed from a data set, such as a file, and is used to verify the integrity of the data. It is a way to detect errors or changes in the data by comparing the checksum value before and after the data transfer or storage)
Checksum databases are commonly used in a various scenarios, such as:
  - Data Integrity Verification: To ensure that data remains unchanged during transmission or storage, especially in criticl systems or sensitive data transfers
  - Anti-Malware and Security: In antivirus software, checksum databases are used to detect known malware by comparing the checksums of files against a database of known malicious checksums.
  - Data Deduplications: In storage systems, checksums can be used to identify duplicate data blocks and save storage space.
  - Software Updates: Software vendors use checksums to verify the integrity of software updates before installation
