# Passkey-CLI 🔑

A simple, local password manager built with Go and Cobra. This project serves as a learning exercise for building command-line interface (CLI) applications in Go.

-----

### ⚠️ Disclaimer: Learning Project

This is one of my first projects in Go, created for learning purposes. It was a great way for me to understand CLI development with Cobra, file I/O, and basic cryptography in Go.

**Please DO NOT use this tool to store real, sensitive passwords.** The project might contain bugs, security vulnerabilities, or practices that are not ideal for a production-grade security application or anything close to that.

***

### A Note on Vault Location

For development and testing convenience, the `vault.json` file is currently created in the same directory where you run the command.

If you'd like to experiment, you can easily change this to a more conventional location, such as your user's home or configuration directory. To do so, simply modify the `GetVaultPath()` function within the `utils` package to return a static path like `~/.config/passkey-cli/vault.json`.

This change would prevent a new vault from being created in every directory you use the tool from, making it feel more like a system-wide application. Just remember, this is still a learning project, and this modification does not make it secure for real-world use.

-----

### Features

  - **Secure Vault Initialization**: Creates a local vault protected by a master passkey.
  - **CRUD Operations**: Full support for adding, listing, updating, and deleting services.
  - **Automatic Password Generation**: Automatically generates a strong, random password for each new service.
  - **Clipboard Integration**: Easily copy any service's password directly to the clipboard.
  - **Secure Passkey Hashing**: Uses **Argon2id** to securely hash the master passkey.

-----

### Installation

You need to have **Go** installed on your system to run this application.

1.  Clone the repository to your local machine:

    ```bash
    git clone https://github.com/diegoAndradeD/passkey-cli.git
    cd passkey-cli
    ```

2.  Install the binary:

    ```bash
    go install .
    ```

    This will compile the project and place the `passkey-cli` executable in your Go bin directory (usually `$GOPATH/bin` or `$HOME/go/bin`). Make sure this directory is in your system's `PATH`.

-----

### Usage

All commands require a master **passkey** to access the vault. This ensures that your stored credentials can only be accessed by you.

#### 1\. First-Time Setup

Before you can use the password manager, you must initialize the vault. This command creates the `vault.json` file in your current directory.

```bash
passkey-cli setup --passkey "your-strong-master-password"
```

-----

#### 2\. Add a New Service

Adds a new service to the vault and automatically generates a password for it.

**Usage:**

```bash
passkey-cli add --name <service-name> --passkey <your-master-passkey>
```

**Example:**

```bash
passkey-cli add --name "github" --passkey "your-strong-master-password"
```

-----

#### 3\. List Services

Lists all stored services or shows the details for a specific one.

**Usage:**

```bash
# List all services
passkey-cli list --passkey <your-master-passkey>

# Show details for a specific service
passkey-cli list --service <service-name> --passkey <your-master-passkey>
```

**Examples:**

```bash
passkey-cli list --passkey "your-strong-master-password"

passkey-cli list --service "github" --passkey "your-strong-master-password"
```

-----

#### 4\. Copy a Password

Copies a service's password directly to your system's clipboard.

**Usage:**

```bash
passkey-cli copy --name <service-name> --passkey <your-master-passkey>
```

**Example:**

```bash
passkey-cli copy --name "github" --passkey "your-strong-master-password"
```

-----

#### 5\. Update a Service

Updates a service's name and can optionally regenerate its password.

**Usage:**

```bash
passkey-cli update --old <current-name> --new <new-name> --passkey <your-master-passkey>
```

**Examples:**

```bash
# Update the name only
passkey-cli update --old "github" --new "github-work" --passkey "your-strong-master-password"

# Update the name and regenerate the password
passkey-cli update --old "github" --new "github-work" --regen --passkey "your-strong-master-password"
```

-----

#### 6\. Delete a Service

Removes a service from the vault permanently.

**Usage:**

```bash
passkey-cli delete --name <service-name> --passkey <your-master-passkey>
```

**Example:**

```bash
passkey-cli delete --name "github" --passkey "your-strong-master-password"
```

-----

### How It Works

  - The application stores all data in a single `vault.json` file in the directory where you run the command.
  - The contents of the vault are stored in plain text, but the vault itself is protected by a hashed master passkey.
  - The master passkey you provide is hashed using **Argon2id**, a modern and secure key derivation function.
  - Access to the vault is only granted if the passkey you provide matches the stored hash. This prevents anyone from reading the vault file without knowing the master passkey.

-----

### License

This project is licensed under the MIT License.
