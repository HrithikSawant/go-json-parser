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
