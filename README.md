# Work in Progress in studying Golang properly

Chang Tan<br>
changtan@listerunlimited.com<br>
Two-time National Cyber League Winner (Individuals & Team Fall 2021 and Spring 2022)

<a href="https://www.udemy.com/course/building-modern-web-applications-with-go">Udemy Course which this code is derived from</a>

I am just doing this to study more advanced methods of Golang programming, specifically web servers and web applications since alot of features of Golang is not currently present despite being well documented in... uh... shittier languages like Java.

I'm planning to use this as a template codebase to build custom libraries and Golang packages for...

1. Universal reverse proxies to protect web apps concealed behind a virtual private cloud with a built-in DNS resolver, input validation, support for common API methods such as REST, SOAP, JSON-RPC, and GraphQL
2. Cross-compilable mobile apps
3. Agricultural robotics libraries and semi-conscious "artificial intelligence"
4. Counteracting reverse-engineering attempts by competitors for production-level deployments

# Cross compilation instructions (Linux)

GOOS=linux GOARCH=amd64 go build -o webapp.elf cmd/web/*.go

# Cross compilation instructions (Windows)
GOOS=windows GOARCH=amd64 go build -o webapp.exe cmd/web/*.go

# Cross compilation instructions (MacOS Desktop/MacBooks)
GOOS=darwin GOARCH=amd64 go build -o webapp.macho cmd/web/*.go
