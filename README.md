# 📝 NixMkDoc: Simple and Fast Nix Documentation

NixMkDoc is a lightweight tool written in **Golang** designed to make documenting your **Nix** files effortless. It scans your `.nix` files, finds options defined with `lib.mkOption`, and automatically generates a **Markdown** file with all the details you need.

Originally created to help with the documentation for the **Neve Project**, NixMkDoc is perfect for anyone who wants an easy way to document their Nix projects without hassle.

## ✨ Why Use NixMkDoc?

- **Easy to Use**: No complicated setup. Just point it to a file or directory, and it does the work for you.
- **Automatic Documentation**: Extracts key details such as:
  - Full option name
  - Type and default value
  - Description
  - Examples (if available)
- **Markdown Output**: Ready to use in GitHub, README files, or other documentation platforms.

## 🚀 How Does It Work?

### Installation

#### Using Nix

If you have Nix installed, you can run NixMkDoc directly from the repository:

```bash
nix run github:Neve-Project/nixmkdoc -- --help
```

#### Without Nix (Manual Build)

1. Install **Go**.
1. Clone the repository:
   ```bash
   git clone https://github.com/Neve-Project/nixmkdoc.git
   cd nixmkdoc
   ```
1. Build the executable:
   ```bash
   go build .
   ```
1. Run the program:
   ```bash
   ./nixmkdoc --help
   ```

### How to Use

NixMkDoc works through simple commands. You can:

- **Scan a directory** with `--dir` or `-d`:
  ```bash
  ./nixmkdoc --dir /path/to/nix/project
  ```
- **Document a single file** with `--file` or `-f`:
  ```bash
  ./nixmkdoc --file /path/to/file.nix
  ```
- Use `--help` or `-h` to see all available options:
  ```bash
  ./nixmkdoc --help
  ```

Once executed, NixMkDoc creates a file called `options.md` in the current directory, containing all the extracted documentation.

______________________________________________________________________

## 🤝 Contribute!

Do you have an idea to improve NixMkDoc? Want to add a new feature? Contributing is easy:

1. Fork the repository.
1. Create a branch for your changes:
   ```bash
   git checkout -b feature-your-feature-name
   ```
1. Submit a pull request.

Every suggestion or contribution is welcome!

## 📜 License

NixMkDoc is distributed under the **GPL v2** license. For more details, see the `LICENSE` file.

## 💬 Questions or Issues?

Open an **issue** on the official repository:\
[**Neve-Project/nixmkdoc**](https://github.com/Neve-Project/nixmkdoc)

Thank you for choosing NixMkDoc! 🎉 Making Nix documentation as simple as it should be.
