# 🛠️ Build a JSON Parser Program from Scratch

**go-json-parser** is a CLI tool that reads a JSON file or standard input and determines whether it is syntactically valid.

It focuses on:

- 🔤 Lexical analysis (splitting input into tokens)
- 🧱 Syntax parsing (verifying structural rules)
- 🛠️ Building foundational compiler components in Go

## 🧱 Core Concepts

### 🔍 Lexer (Lexical Analyzer)

The **lexer** (or lexical analyzer) reads raw JSON input and **splits it into tokens**.  
Think of it like a scanner moving left-to-right over the characters in the input string.


Example:

```go
input := `{}`
tokens := []Token{
    {Type: TokenCurlyOpen, Literal: "{"},
    {Type: TokenCurlyClose, Literal: "}"},
    {Type: TokenEOF, Literal: ""},
}
```

### 🧩 Token

A **token** is the smallest meaningful element in the input. 

Examples:

```
- `{` → TokenCurlyOpen
- `}` → TokenCurlyClose
- `true`, `false`, `"string"`, `123` → More token types.
```

Each token in this project is represented by:

- `Type`: An enum describing what the token is (e.g. `TokenCurlyOpen`)
- `Literal`: The actual character(s) read from the input

---

### 🧠 Parser

The **parser** reads the sequence of tokens from the lexer and checks if they match the **expected grammar** (syntax rules). of the JSON subset.

---



## 🚀 How to Run

#### Run from the command line

```bash
./go-json-parser myfile.json
```

- `myfile.json` should contain your JSON data (e.g., `{}`)
- Exit code will be 0 for valid and 1 for invalid

## 🧪 Tests

You can run tests for both the lexer and parser:

```bash
go test ./...
ok  	github.com/HrithikSawant/go-json-parser/lexer	0.002s
ok  	github.com/HrithikSawant/go-json-parser/parser	0.002s
```

##  ✅ Step 1: Recognize Simple JSON Object `{}`

Add functionality to validate the **simplest JSON object**.


```bash
./go-json-parser myfile.json
DEBUG: State = Start                | Token = {          | Literal = {
DEBUG: State = ExpectKeyOrEnd       | Token = }          | Literal = }
DEBUG: State = Done                 | Token = EOF        | Literal = 
Valid JSON structure

```

### ✅ Step 2: Parse JSON Object with String Key-Value Pairs
Add functionality to parse JSON objects with string keys and string values.

```bash
./go-json-parser myfile2.json

DEBUG: State = Start                | Token = {          | Literal = {
DEBUG: State = ExpectKeyOrEnd       | Token = STRING     | Literal = key
DEBUG: State = ExpectColon          | Token = :          | Literal = :
DEBUG: State = ExpectValue          | Token = STRING     | Literal = value
DEBUG: State = ExpectCommaOrEnd     | Token = }          | Literal = }
DEBUG: State = Done                 | Token = EOF        | Literal = 
Valid JSON structure

```

## ✅ Step 3: Support for Primitive Values (`true`, `false`, `null`, numbers)

Extend support for primitive value types in JSON: booleans (`true`, `false`), `null`, and numeric values (integers and floats).

```bash
./go-json-parser myfile3.json

DEBUG: State = Start                | Token = {          | Literal = {
DEBUG: State = ExpectKeyOrEnd       | Token = STRING     | Literal = key1
DEBUG: State = ExpectColon          | Token = :          | Literal = :
DEBUG: State = ExpectValue          | Token = BOOL       | Literal = true
DEBUG: State = ExpectCommaOrEnd     | Token = ,          | Literal = ,
DEBUG: State = ExpectKeyOrEnd       | Token = STRING     | Literal = key2
DEBUG: State = ExpectColon          | Token = :          | Literal = :
DEBUG: State = ExpectValue          | Token = BOOL       | Literal = false
DEBUG: State = ExpectCommaOrEnd     | Token = ,          | Literal = ,
DEBUG: State = ExpectKeyOrEnd       | Token = STRING     | Literal = key3
DEBUG: State = ExpectColon          | Token = :          | Literal = :
DEBUG: State = ExpectValue          | Token = NULL       | Literal = null
DEBUG: State = ExpectCommaOrEnd     | Token = ,          | Literal = ,
DEBUG: State = ExpectKeyOrEnd       | Token = STRING     | Literal = key4
DEBUG: State = ExpectColon          | Token = :          | Literal = :
DEBUG: State = ExpectValue          | Token = STRING     | Literal = value
DEBUG: State = ExpectCommaOrEnd     | Token = ,          | Literal = ,
DEBUG: State = ExpectKeyOrEnd       | Token = STRING     | Literal = key5
DEBUG: State = ExpectColon          | Token = :          | Literal = :
DEBUG: State = ExpectValue          | Token = NUMBER     | Literal = 101
DEBUG: State = ExpectCommaOrEnd     | Token = }          | Literal = }
DEBUG: State = Done                 | Token = EOF        | Literal = 
Valid JSON structure
```

## ✅ Step 4: Support for Nested Objects and Arrays

Extend support to allow **nested JSON objects** and **arrays** as values.

```bash
./go-json-parser myfile4.json

DEBUG: State = Start                | Token = {          | Literal = {
DEBUG: State = ExpectKeyOrEnd       | Token = STRING     | Literal = key
DEBUG: State = ExpectColon          | Token = :          | Literal = :
DEBUG: State = ExpectValue          | Token = STRING     | Literal = value
DEBUG: State = ExpectCommaOrEnd     | Token = ,          | Literal = ,
DEBUG: State = ExpectKeyOrEnd       | Token = STRING     | Literal = key-n
DEBUG: State = ExpectColon          | Token = :          | Literal = :
DEBUG: State = ExpectValue          | Token = NUMBER     | Literal = 101
DEBUG: State = ExpectCommaOrEnd     | Token = ,          | Literal = ,
DEBUG: State = ExpectKeyOrEnd       | Token = STRING     | Literal = key-o
DEBUG: State = ExpectColon          | Token = :          | Literal = :
DEBUG: State = ExpectValue          | Token = {          | Literal = {
DEBUG: State = ExpectKeyOrEnd       | Token = }          | Literal = }
DEBUG: State = ExpectCommaOrEnd     | Token = ,          | Literal = ,
DEBUG: State = ExpectKeyOrEnd       | Token = STRING     | Literal = key-l
DEBUG: State = ExpectColon          | Token = :          | Literal = :
DEBUG: State = ExpectValue          | Token = [          | Literal = [
DEBUG: State = ArrayValueOrEnd      | Token = ]          | Literal = ]
DEBUG: State = ExpectCommaOrEnd     | Token = }          | Literal = }
DEBUG: State = Done                 | Token = EOF        | Literal = 
Valid JSON structure
```

## ✅ Step 5: Add Your Own Edge Case Tests and Run Against JSON Standard Suite

In this step, you should build confidence in your parser by:

- Writing your **own test cases** (e.g. deeply nested objects, invalid characters, unterminated strings, extra commas, etc.)
- Running the parser against the official [JSONChecker Test Suite](http://www.json.org/JSON_checker/test.zip)

---
## 🔗 References

- 📘 [Lexical Analysis (Lexer)](https://en.wikipedia.org/wiki/Lexical_analysis) – Wikipedia article explaining how input is split into tokens.
- 📗 [Parsing (Parser)](https://en.wikipedia.org/wiki/Parsing) – Wikipedia article covers Parsing, syntax analysis, or syntactic analysis.
-  📘 [RFC 8259](https://datatracker.ietf.org/doc/html/rfc8259) – The JavaScript Object Notation (JSON) Data Interchange Format
- 🧰 [Graphical JSON Syntax Guide](https://www.json.org/json-en.html) – Visual representation of JSON grammar and structure.
- 🐉 [*The Dragon Book* – Compilers: Principles, Techniques, and Tools](https://en.wikipedia.org/wiki/Compilers:_Principles,_Techniques,_and_Tools) – Classic textbook on compiler construction.

---

## 📘  Note
Originally sparked by the **"Build Your Own JSON Parser"** challenge from [codingchallenges.fyi](https://codingchallenges.fyi) by John Crickett



## 🪪 License

This project is licensed under the MIT License — see [LICENSE](./LICENSE) for details.
