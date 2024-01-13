# godirect

godriect is my personal URL redirect service. This has the purpose of providig a
reliable way for me to share links on social media. When I'm writing a post for
my employer and share this post via godirect, I want the ability to exchange the
URL to point to my own blog in case my employer decides to remove my post or the
whole blog page (yes, I'm looking at you fino).

## Why not just an nginx?

I don't know. Writing Go apps is just fun. This is a really minimalistic Go
http server with a configurable list of redirects.

## Getting started

This section provides a quick guide to get you up and running with the
application. 

### Prerequisites

- Ensure you have Go installed on your system. You can download and install Go
from [the official site](https://go.dev/dl/).

### Building the Application

To build the application, navigate to the directory containing the source code
and run the following command:

```bash
go build -o godirect cmd/api/main.go
```

This command compiles the source code into an executable named godirect.

### Running the Application

After building the application, you can run it by providing the path to your
YAML configuration file as follows:

```bash
./godirect -redirects /path/to/your/redirects.yaml
```

Replace **/path/to/your/redirects.yaml** with the actual path to your YAML
configuration file.

### Example YAML Configuration

Here's an example of the YAML configuration file for redirects:

```yaml
Redirects:
  - Path: foo-bar-123
    Target: 'https://something.de'
  - Path: important-topic
    Target: 'https://something-else.de'
```

Copy this content into a file (e.g., **redirects.yaml**) and use the path to
this file when running the application as shown above.

