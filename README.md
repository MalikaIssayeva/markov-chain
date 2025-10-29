
---

# 🧠 Markov Chain Text Generator

A **Go** program that generates realistic-sounding random text using a **Markov Chain** algorithm.
This project focuses on **algorithms**, **I/O handling**, **file processing**, and **software design principles** — with special attention to **data structure design** and **performance**.

---

## 📚 Learning Objectives

* Understand and implement the **Markov Chain** algorithm
* Practice **input/output** and **file handling** in Go
* Design efficient **data structures** for large-scale text processing
* Apply **software design principles** (clarity, modularity, data-first thinking)

---

## 🧩 Abstract

This project implements a **text generator** that learns from a source text and produces new text that mimics its style. The algorithm analyses sequences of words (prefixes) and records which words typically follow them (suffixes). When generating text, it randomly selects suffixes based on these learned probabilities.

This method simulates how predictive text systems work (like those on your phone keyboard). The algorithm uses a **Markov chain**, where the next state (word) depends only on the current state (prefix of previous words).

---

## ⚙️ Algorithm Overview

The **Markov Chain algorithm** works as follows:

1. Read all words from the input text.
2. Create a mapping between **prefixes** (sequences of N words) and the list of **possible suffixes** that follow them.
3. Start from an initial prefix and repeatedly:

   * Randomly pick one of its suffixes,
   * Print it,
   * Slide the prefix window forward by one word.

### Example

Given the text:

> Bad programmers worry about code. Good programmers worry about data-structures and their relationships.

| Prefix                | Possible Suffixes      |
| --------------------- | ---------------------- |
| Bad programmers       | worry                  |
| programmers worry     | about, about           |
| worry about           | code., data-structures |
| about code.           | Good                   |
| Good programmers      | worry                  |
| about data-structures | and                    |
| data-structures and   | their                  |
| and their             | relationships.         |

Generated sequence:

```
Bad programmers worry about data-structures and their relationships.
```

---

## 🧠 Design Decisions

### Data Structure

A **map** is used to efficiently store and retrieve prefix–suffix relationships:

```go
map[string][]string
```

* **Key:** A joined string of `prefixLength` words (the prefix)
* **Value:** A list of all suffix words that followed this prefix in the input

This allows:

* O(1) lookups during text generation
* Efficient appending of new suffixes
* Flexibility for different prefix lengths

---

## 🏗️ Program Architecture

The program runs in two main stages:

1. **Build Stage** – Reads input, tokenizes it into words, constructs prefix→suffix mappings.
2. **Generate Stage** – Randomly generates new text based on those mappings.

---

## 🧰 Installation and Compilation

```bash
$ go build -o markovchain .
```

This produces an executable named **`markovchain`** in the current directory.

---

## ▶️ Usage

```bash
markovchain [-w <N>] [-p <S>] [-l <N>]
markovchain --help
```

### Options

| Option   | Description                                | Constraints     |
| -------- | ------------------------------------------ | --------------- |
| `--help` | Show usage information                     | —               |
| `-w N`   | Maximum number of words to generate        | `0 < N ≤ 10000` |
| `-p S`   | Starting prefix (must exist in input text) | String of words |
| `-l N`   | Prefix length                              | `1 ≤ N ≤ 5`     |

---

## 📥 Input and Output

* **Input:** Entire text read from **stdin**
* **Output:** Generated text printed to **stdout**

### Examples

#### 1️⃣ Default run

```bash
$ cat the_great_gatsby.txt | ./markovchain | cat -e
Chapter 1 In my younger and more stable, become for a job. He hadn't eat anything for a long, silent time. ...
```

* Default prefix length: `2`
* Default maximum words: `100`
* Default starting prefix: first two words of input

#### 2️⃣ Limit number of words

```bash
$ cat the_great_gatsby.txt | ./markovchain -w 10
Chapter 1 In my younger and more stable, become for
```

#### 3️⃣ Custom starting prefix

```bash
$ cat the_great_gatsby.txt | ./markovchain -w 10 -p "to play"
to play for you in that vast obscurity beyond the
```

#### 4️⃣ Custom prefix length

```bash
$ cat the_great_gatsby.txt | ./markovchain -w 10 -p "to something funny" -l 3
to something funny the last two days," remarked Wilson. "That's
```

#### 5️⃣ Help

```bash
$ ./markovchain --help
Markov Chain text generator.

Usage:
  markovchain [-w <N>] [-p <S>] [-l <N>]
  markovchain --help

Options:
  --help  Show this screen.
  -w N    Number of maximum words
  -p S    Starting prefix
  -l N    Prefix length
```

#### 6️⃣ Error handling

```bash
$ ./markovchain
Error: no input text
```

---

## 🚦 Constraints & Validation

| Case                  | Constraint     | Action                           |
| --------------------- | -------------- | -------------------------------- |
| No input text         | —              | Print `"Error: no input text"`   |
| Negative word count   | `< 0`          | Print error                      |
| Word count > 10,000   | —              | Print error                      |
| Prefix not in text    | —              | Print error                      |
| Invalid prefix length | `< 1` or `> 5` | Print error                      |
| Runtime panic         | —              | Disqualifies program (grade = 0) |

---

## 🧾 Implementation Guidelines

* Must compile successfully using:

  ```bash
  go build -o markovchain .
  ```
* Must **not** panic (handle all errors gracefully)
* Must use **only built-in packages**
* Must follow **gofumpt** formatting standards
  (zero tolerance — non-compliant code = 0 grade)

---

## 🧠 Key Takeaways

> "Bad programmers worry about code. Good programmers worry about data structures and their relationships."
> — *Linus Torvalds*

This project reinforces that **good software design** begins with **data modeling**, not just code writing. By thinking carefully about how data flows and interacts, the code naturally becomes more efficient, elegant, and maintainable.

---

## 📄 License

This project is provided for educational purposes.
All literary text used for training (e.g., *The Great Gatsby*) remains under its original copyright.

---
