# mdserve
A web server providing RESTful API for markdown contents

## Usage

```bash
$ mdserve <path/to/dir>
```

This command starts a web server. It provides following endpoints.

* **`/path/to/dir/`**: return JSON response including list of contents.
* **`/path/to/dir/contents.json`**: return JSON response including a content.
