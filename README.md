
# sanity-groq-stress-test

This utility is designed to execute GROQ, a query language designed by the Sanity team.

It reads a text file as an input argument. Each line in the provided file is read and concurrently executed. An output csv is generated containing the timed query results. The `-n` flag can be provided to specify the number of samples to gather.

An .env file containing Sanity project configuration is looked for in the directory.

A samples folder is provided with an example env and query file containing the samples to test.

```txt
SANITY_STUDIO_PROJECT_ID=foo
SANITY_STUDIO_API_DATASET=development
```

## Usage

The command can be run via a terminal.

```bash
groq-test -out example-test.csv ./samples/get-by-id.txt
```

```bash
Usage of groq-test:
  -e example.env
        optionally specify an env file with example.env (default ".env")
  -out output.csv
        optionally specify a file to output results output.csv (default "output.csv")
```

## Build

```bash
go build -o ./bin/santiy-test ./src/main.go
```

## Install

Run the following command while in the project directory. This generates a line to include to your shell path.
Alternatively it can be run in place provided the path is correct `./bin/groq-test`

```bash
printf 'export PATH="%s:$PATH"' "$PWD/bin"
```
