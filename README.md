# Coursera - Golang for Beginners - Travel Fare Calculator

This repository tracks my progress on the main project for the Coursera course, "Golang for Beginners: Create a travel fare calculator." The goal is to build a command-line application that calculates travel fares based on user input for origin, destination, and cabin class.

## 🎯 Project Goal

The application is designed to be an interactive console program that performs the following tasks:

1. Prompts the user to enter an origin city, destination city, and cabin class.
2. Validates the user's input against a predefined set of data.
3. Calculates the distance and the final travel fare based on the inputs.
4. Displays the formatted results to the user.

## ✅ Project Status: Complete

This repository contains the final, completed version of the application. All core features have been implemented:

* The main program flow and structure[cite: 11].
* User input handling for origin, destination, and cabin class via the console.
* Custom data structures (`structs`) to hold city and travel class data.
* Validation logic to ensure the origin, destination, and cabin class codes entered by the user are valid.
* Implementation of distance and fare calculation using the `math` package.
* Display of formatted results to the user using `fmt.Printf`.

## 💡 Key Concepts Learned & Applied

This project serves as a practical application of fundamental Golang concepts covered in the course:

* **Packages & Program Flow**:
  * Using the `main` package as the program entry point and the `main` function which is run automatically.
  * Employing the `fmt` package for console input/output operations.
  * Utilizing the `math` package to perform complex mathematical calculations.

* **Data Structures**:
  * Defining custom data types using **structs** to group variables of different types.
  * Using **arrays** to hold fixed-size collections of elements.

* **Functions & Control Flow**:
  * Creating functions that can return multiple values.
  * Iterating through data structures using the `range` keyword.

* **Variables & Pointers**:
  * Using the `:=` operator for concise variable declaration and initialization.
  * Using the `&` operator to get the memory address of a variable (pointer).

* **Error Handling**:
  * Following Go's convention of using an `error` type to convey error conditions, where `nil` means no error occurred.
  * Using the underscore (`_`) to explicitly discard unused function return values.

* **Build & Formatting**:
  * Using `fmt.Printf` with templates to output strings with formatted variable values.
  * Using the `go build` command to compile the application into a single executable file for production.

## 🎓 Course Context

This project is the primary assignment for the Coursera course linked below. It is designed to reinforce beginner-level Go programming concepts in a practical, hands-on manner.

* **Course Link:** [Golang for Beginners: Create a travel fare calculator](https://www.coursera.org/projects/golang-beginners-data-types-functions-packages)
