# Temperature Calculator - CLI Apps

This is a command-line interface (CLI) application written in Golang for converting temperatures between different units. It includes conversion functions for Celcius, Fahrenheit, Kelvin, and Reamur. The app provides two main functionalities: converting the current temperature or performing a custom conversion based on user input. [Demo](https://drive.google.com/file/d/1cJWLd5KgNefEBJm3k3TgBSyLBQKZuHyT/view?usp=sharing)

## Feature

1. **Convert current temperature**
Convert the temperature currently (celcius) set in the application to either Fahrenheit, Kelvin, or Reamur.
2. **Custom convert temperature**
Input a temperature manually and convert it between Celcius, Fahrenheit, Kelvin, and Reamur.

## How to run

### Prerequisites

- Go 1.22.1 or higher
- Ensure that you have Golang installed on your system. If not, you can download it from https://golang.org/dl/.

### Installation

1. Clone this repository to your local machine:
    
    ```powershell
    git clone <https://github.com/Safiramdhn/project-app-cli-golang-safira.git>
    ```
    
2. Navigate to project directory
    
    ```powershell
    cd project-app-cli-golang-safira
    ```
    
3. Run CLI Program
    
    ```powershell
    go run main.go
    ```
    

## Project Structure

- `main.go`: Contains the main logic of the application, including displaying the menu, capturing user input, and invoking the temperature conversion functions.
- `converter`: Package that handles temperature conversion logic.
- `suhu`: Contains the definitions for temperature structs (e.g., Celcius, Fahrenheit, etc.).

## Usage

Upon running the application, you will be presented with a menu offering two options:

1. Convert the current temperature.
2. Perform a custom conversion by inputting a specific temperature.

Simply follow the prompts in the terminal to use the application.
