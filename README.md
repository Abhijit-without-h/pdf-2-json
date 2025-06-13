
---

# 📄 PDF to JSON Converter (Go)

A simple command-line tool written in Go that extracts plain text from a PDF file and exports it as structured JSON.

## 🚀 Features

* Converts PDF content into plain text
* Outputs the result in JSON format
* Lightweight and easy to use
* Built with [`github.com/ledongthuc/pdf`](https://github.com/ledongthuc/pdf)

## 📦 Installation

1. **Clone the repo:**

```bash
git clone https://github.com/your-username/pdf2json-go.git
cd pdf2json-go
```

2. **Install dependencies:**

```bash
go mod tidy
```

3. **Build the project:**

```bash
go build -o pdf2json
```

## 🛠 Usage

```bash
./pdf2json -input input.pdf -output output.json
```

### Example

```bash
./pdf2json -input sample.pdf -output sample.json
```

## 🧾 Output Format

```json
{
  "content": "Extracted text from the PDF goes here..."
}
```

## ⚠️ Notes

* This tool reads all pages of the PDF and concatenates the text into a single string.
* It only extracts **plain text**, not images, tables, or formatting.

## 📚 Dependencies

* [ledongthuc/pdf](https://github.com/ledongthuc/pdf)

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---
