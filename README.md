# makescraper

A Go web scraper built with [Colly](https://github.com/gocolly/colly). It crawls
[quotes.toscrape.com](https://quotes.toscrape.com/) and extracts each quote's
text, author, and tags.

## Install

```bash
git clone git@github.com:Joshkaki00/makescraper.git
cd makescraper
go mod download
```

Requires Go 1.27+.

## Usage

```bash
go run scrape.go
```

The scraper:

1. Visits the homepage and collects every `.quote` block
2. Prints each quote to stdout
3. Prints the full result set as indented JSON
4. Writes that JSON to `output.json` in the project root

Expected output (abbreviated):

```
Visiting https://quotes.toscrape.com/

Scraped 10 quotes:

"The world as we have created it is a process of our thinking. It cannot be changed without changing our thinking."
  - Albert Einstein [change deep-thoughts thinking world]

...

JSON output:
[
  {
    "text": "...",
    "author": "Albert Einstein",
    "tags": ["change", "deep-thoughts", "thinking", "world"]
  },
  ...
]

Wrote 10 quotes to output.json
```

## Configuration

Settings live in `scrape.go`:

- `colly.AllowedDomains("quotes.toscrape.com")` — crawl one host only
- `colly.UserAgent("makescraper/1.0 ...")` — identifies the client
- `c.Limit(&colly.LimitRule{...})` — parallelism 2, ~1s random delay
- `c.SetRequestTimeout(30 * time.Second)` — per-request timeout
- `c.IgnoreRobotsTxt = false` — respects robots.txt

## License

MIT — see [LICENSE](LICENSE).
