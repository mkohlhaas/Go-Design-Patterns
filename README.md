# Go Design Patterns
This is the code repository for [Go Design Patterns](https://www.packtpub.com/application-development/go-design-patterns?utm_source=github&utm_medium=repository&utm_campaign=9781786466204), published by [Packt](https://www.packtpub.com/?utm_source=github). It contains all the supporting project files necessary to work through the book from start to finish.
## About the Book
This book will take you through the history of concurrency, how Go utilizes it, how Go differs from other languages, and the features and structures of Go's concurrency core. Each step of the way, the book will present real, usable examples with detailed descriptions of the methodologies used. By the end, you will feel comfortable designing a safe, data-consistent, high-performance concurrent application in Go.


## Instructions and Navigation
All of the code is organized into folders. Each folder starts with a number followed by the application name. For example, ch02.

The code will look like the following:
```
package main

func main() {
  ten := 10
  if ten == 20 {
    println("This shouldn't be printed as 10 isn't equal to 20")
  } else {
    println("Ten is not equals to 20")
  }
}
```

Most of the chapters in this book are written following a simple TDD approach, here the requirements are written first, followed by some unit tests and finally the code that satisfies those requirements. We will use only tools that comes with the standard library of Go as a way to better understand the language and its possibilities. This idea is key to follow the book and understanding the way that Go solves problems, especially in distributed systems and concurrent applications.
