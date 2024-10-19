
![Logo](https://github.com/DiegoDev2/Fleet/blob/main/Fleet(5).png?raw=true)


# Fleet
**Fleet** is a package manager built with Go, designed to install, manage, and configure different tools in a simple and efficient way. It allows you to pull and install tools directly from repositories.

### Other Readme
![Español](https://github.com/DiegoDev2/Fleet/blob/main/READMEes.md)
![Русский](https://github.com/DiegoDev2/Fleet/blob/main/READMEru.md)
### Features

- Lightweight and efficient tool installation.
- Simple commands for installing, updating, and removing tools.
- Customizable formulas for different tools.

### Installation

1. Clone the repository:

```bash
   git clone https://github.com/DiegoDev2/Fleet
   cd fleet
   go build -o fleet
   sudo mv fleet /usr/local/bin
```
Usage

To install a tool, simply use:

```bash
fleet install <tool-name>
```
For example, to install nmap:

```bash

fleet install nmap
```
### Adding New Tools

- Open `lib/tools.go` and add the new tool to the list of available tools.

- Create a formula for the tool in the `formulas/` directory.

- Use the install command to add it to your system.

### Contributing

We welcome contributions to improve Fleet. Feel free to submit a pull request or open an issue if you find a bug or have a suggestion.
### License

Fleet is licensed under the Apache 2.0 License. See the LICENSE file for more details. 
