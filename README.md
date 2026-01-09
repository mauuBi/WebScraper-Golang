Go TechCrunch Scraper
A high-performance web scraper built in Go that extracts articles from TechCrunch and stores them in a structured database. This project was developed from scratch to master the Go ecosystem, focusing on the standard library and DOM manipulation.

Features
Targeted Extraction: Specifically designed to parse TechCrunch article data (titles, links, authors, and timestamps).

Native HTTP Client: Uses net/http with custom headers to simulate authentic browser traffic.

DOM Navigation: Implements goquery to traverse the HTML structure and isolate post elements.

Database Integration: Automatically maps extracted data to a database schema for long-term storage.

JSON Serialization: Full support for converting Go structs into JSON format.

Project Architecture
The application is built around a modular pipeline to ensure clean code and easy maintenance:

Request Layer: Handles the connection to the target URL and manages the response lifecycle.

Parsing Layer: Uses CSS selectors to identify post containers and extract relevant metadata into Go structs.

Persistence Layer: Connects to the database and handles the insertion of structured data, preventing raw HTML clutter.

Technical Stack
Language: Go

HTTP Client: net/http

HTML Parser: goquery

Data Format: JSON / SQL (Database)

Installation
Bash

git clone https://github.com/yourusername/go-web-scraper.git
cd go-web-scraper
go mod tidy
Usage
Configure your database credentials in the environment variables or the config file, then run:

Bash

go run main.go
Key Learning Milestones
Resource Management: Efficiently closing response bodies using defer to prevent memory leaks.

Data Transformation: Successfully converting unstructured HTML into a clean, queryable database format.

Header Customization: Learning how to bypass basic bot detection using custom User-Agent strings.

Type Safety: Leveraging Go's strong typing system to ensure data integrity before database insertion.
