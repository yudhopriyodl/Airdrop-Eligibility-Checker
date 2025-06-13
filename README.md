# Airdrop Eligibility Checker 🚀

📝 **Project Description**
This Go project automates the process of checking eligible users for a cryptocurrency airdrop based on predefined criteria. It reads user data from a JSON file, applies eligibility rules, and outputs a list of users who qualify for the airdrop.

✨ **Features**

- **Automated Eligibility Check**: Quickly determine eligible users based on custom rules.
- **Configurable User Data**: Easily integrate user data via a `users.json` file.
- **Clear Output**: Generates a clean list of eligible users.
- **Go-based**: High performance and concurrency for efficient processing.

⚙️ **Prerequisites**
Before you begin, ensure you have the following installed:

- **Go**: Version 1.16 or higher. You can download it from [golang.org](https://golang.org/dl/).

📥 **Installation**

1. **Clone the repository**:

    ```bash
    git clone https://github.com/yudhopriyodl/Airdrop-Eligibility-Checker.git
    cd Airdrop-Eligibility-Checker
    ```

2. **Download dependencies**:

    ```bash
    go mod tidy
    ```

🚀 **Usage**
To run the airdrop eligibility checker, execute the following command in the project root directory:

```bash
go run main.go
```

🛠 **Configuration**
The project relies on a `users.json` file for user data. Ensure this file is present in the root directory of the project.

**`users.json` format example:**

```json
[
  {
    "id": "user123",
    "wallet_address": "0xAbC123...",
    "balance": 1000,
    "transactions": 5
  },
  {
    "id": "user456",
    "wallet_address": "0xDeF456...",
    "balance": 50,
    "transactions": 2
  }
]
```

(Note: The actual fields and criteria for eligibility will depend on the `main.go` implementation.)

📤 **Output**
Upon successful execution, the program will print a list of eligible users to the console. The output format will depend on the `main.go` implementation, but typically it will list user IDs or wallet addresses.

Example output:

```
Eligible Users:
- user123 (0xAbC123...)
- user789 (0xGhI789...)
```

📄 **License**
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. (Note: Create a `LICENSE` file if you intend to use the MIT License.)

🧑‍💻 **Contributing**
Contributions are welcome! Please feel free to submit a Pull Request or open an issue if you find a bug or have a feature request.
