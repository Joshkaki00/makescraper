# makescraper

A Go web scraper built with Colly and goquery. It crawls the quotes.toscrape.com
practice site and extracts each quote's text, author, and tags.

## Install

```bash
git clone git@github.com:Joshkaki00/makescraper.git
cd makescraper
go mod download
```

## Usage

```bash
go run scrape.go
```

Expected output:

```
Visiting https://quotes.toscrape.com/

Scraped 10 quotes:

"The world as we have created it is a process of our thinking. It cannot be changed without changing our thinking."
  - Albert Einstein [change deep-thoughts thinking world]

...

Wrote 10 quotes to output.json
```

The scraper prints each quote to stdout, then writes the full result set as
JSON to `output.json` in the project root.

## Configuration

The target site, rate limits, and timeout are set directly in `scrape.go`:

- `colly.AllowedDomains("quotes.toscrape.com")` restricts the crawl to one host
- `c.Limit(&colly.LimitRule{...})` caps concurrency and adds a delay between requests
- `c.SetRequestTimeout(30 * time.Second)` bounds how long a single request can hang
- `c.IgnoreRobotsTxt = false` respects the site's robots.txt

## License

No license file yet. Add one before treating this as reusable outside the
original assignment.
