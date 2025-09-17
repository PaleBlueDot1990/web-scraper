# Top 10 Action Items for Web Scraper Project

This document lists the **10 highest-priority improvements** for the web-scraper project, with details on **what** to do and **how** to do it.

---

## 1. Add a complete `README.md`
**Why:** Usability & onboarding.  
**Where:** `README.md` (currently placeholder).

**How:**
- Add project description, prerequisites, build & run steps (`go build`, `go run`).
- Document CLI flags (see item 2).
- Show sample outputs (JSON/CSV).
- Include how to run tests (`go test ./...`).
- Add contribution and license section.

**What this means**
The README file is the front door of our project. It’s usually the first thing someone sees when they open the repository. Right now, it’s empty and does not tell anyone how to use the scraper.

**Why it matters**
- Without documentation, new developers or external users won’t know how to run the project.  
- A strong README reduces onboarding time and support questions.  
- It makes the project look professional and trustworthy.  

**What success looks like**
- A newcomer can clone the repo, follow the README instructions, and run the scraper in less than 5 minutes.  
- The README explains what the scraper does, how to install it, how to run it, and what kind of output to expect.  
- Includes troubleshooting and testing instructions.  


---

## 2. Add explicit CLI flags & config support
**Why:** Configurable & automatable.  
**Where:** `main.go`, `configure.go`.

**How:**
- Use `flag` or `pflag` to expose:
  - `--start-url`, `--max-depth`, `--concurrency`, `--per-host`, `--output`, `--output-format`, `--ignore-robots`, `--timeout`, `--resume`, `--store`, `--verbose`.
- Flags override env vars → env vars override config file.
- Print `--help` with usage instructions.

**What this means**
Right now, the scraper only takes raw command-line arguments in a rigid way. We want to improve this so users can run the tool with clear, descriptive flags (like `--start-url https://example.com`).

**Why it matters**
- Makes the tool more user-friendly.  
- Allows automation in scripts and CI pipelines.  
- Ensures repeatability: people can rerun the same command and get the same results.  

**What success looks like**
- The scraper supports modern CLI flags for configuration.  
- Users can set important options like starting URL, crawl depth, output format, etc.  
- Running the scraper with `--help` shows all available options with descriptions.  


---

## 3. Implement `robots.txt` support
**Why:** Ethical & safe crawling.  
**Where:** New `robots.go`, integrate into `crawl_page.go`.

**How:**
- Fetch and parse `robots.txt` per host.
- Support `User-agent: *`, `Disallow`, and `Crawl-delay`.
- Default: honor robots.  
- Add flag `--ignore-robots` to bypass.

**What this means**
Most websites publish a `robots.txt` file that tells crawlers which parts of the site they’re allowed to visit. Respecting this file makes our scraper a "polite citizen" of the internet.

**Why it matters**
- Avoids legal and ethical issues by following site owners’ rules.  
- Prevents the scraper from overloading websites or accessing forbidden areas.  
- Makes the project usable in professional environments where compliance is mandatory.  

**What success looks like**
- By default, the scraper checks `robots.txt` and avoids disallowed pages.  
- The scraper respects crawl delays (if specified).  
- There is an option to override this behavior (`--ignore-robots`) for testing purposes.  


---

## 4. Add robust HTTP client (timeouts, retries, backoff)
**Why:** Reliability & network safety.  
**Where:** `get_html.go`.

**How:**
- Create single `http.Client` with tuned `Transport`.
- Implement retry wrapper with exponential backoff + jitter.
- Handle `429 Too Many Requests` with `Retry-After`.
- Make retry count configurable.

**What this means**
When the scraper requests pages from the internet, those requests sometimes fail (network hiccups, server errors, rate limits). We need to handle these situations gracefully.

**Why it matters**
- Prevents the scraper from crashing on simple network errors.  
- Avoids spamming a website with repeated requests when it’s overloaded.  
- Provides predictable behavior and better reliability.  

**What success looks like**
- The scraper automatically retries failed requests with increasing wait times.  
- It respects server instructions like “retry after 30 seconds.”  
- Runs complete successfully even if some pages are temporarily unavailable.  


---

## 5. Add per-host rate limiting & worker pool
**Why:** Polite crawling & predictable concurrency.  
**Where:** `crawl_page.go`, `get_html.go`.

**How:**
- Use `rate.Limiter` (`golang.org/x/time/rate`) per host.
- Enforce `--per-host` concurrency.
- Central URL queue with N worker goroutines (`--concurrency`).
- Before fetch, wait on host limiter.

**What this means**
Currently, the scraper can launch many requests at once. Without control, this can overwhelm a single website. Rate limiting ensures we don’t hit one site too hard, and a worker pool controls overall concurrency.

**Why it matters**
- Prevents our scraper from being blocked or blacklisted by websites.  
- Ensures consistent performance without spikes in resource usage.  
- Lets us balance speed (faster scraping) with politeness.  

**What success looks like**
- Users can set the maximum number of total concurrent requests.  
- Users can also limit requests per individual website.  
- The scraper runs efficiently while keeping websites happy.  


---

## 6. Strengthen URL normalization & deduplication
**Why:** Correctness & avoiding duplicates.  
**Where:** `normalize_url.go`, `get_urls_from_html.go`.

**How:**
- Lowercase scheme/host, strip default ports, drop fragments.
- Optionally strip tracking params (`utm_*`, `fbclid`).
- Sort query params for canonical form.
- Respect `<base>` tag in HTML.
- Expand tests in `normalize_url_test.go`.

**What this means**
Websites often have multiple URLs that point to the same page (with different cases, tracking parameters, or fragments). Without proper normalization, the scraper may crawl the same content multiple times.

**Why it matters**
- Avoids wasting time and resources on duplicates.  
- Produces cleaner, more accurate reports.  
- Makes analysis easier since every page has a single “canonical” URL.  

**What success looks like**
- The scraper treats equivalent URLs as the same.  
- Tracking parameters like `utm_source` are removed (configurable).  
- Reports show unique pages only, not duplicates.  


---

## 7. Add persistent resume with SQLite
**Why:** Resume long crawls & auditability.  
**Where:** New `storage_sqlite.go`, integrate into `crawl_page.go`.

**How:**
- Use `database/sql` with SQLite.
- Tables: `pages` (url, normalized_url, status, last_crawled, content_hash), `edges` (src, dst, anchor).
- On `--resume`, skip already-seen URLs.
- On finish, upsert crawled pages + edges.

**What this means**
Right now, if the scraper is stopped halfway, it loses all progress. By saving crawl data in a small local database (SQLite), we can pause and resume runs.

**Why it matters**
- Saves time and bandwidth on large crawls.  
- Makes the scraper usable for long-running jobs (hours or days).  
- Enables more advanced reporting and history tracking.  

**What success looks like**
- The scraper can be run with a `--resume` option to pick up where it left off.  
- Crawl results (pages, links) are stored in a local database file.  
- Users can query this database independently for analysis.  


---

## 8. Add multi-format outputs (JSON, CSV, NDJSON)
**Why:** Downstream analysis & flexibility.  
**Where:** `print_report.go`, new `output_*.go`.

**How:**
- Add `--output-format=json|csv|ndjson`.
- Implement adapters:
  - JSON: all results in one file.
  - NDJSON: one record per line (streaming).
  - CSV: edge list (`src,target,anchor,rel,status`).
- Update README with examples.

**What this means**
Currently, the scraper prints results directly to the console. We want to let users export data in formats commonly used for analysis.

**Why it matters**
- JSON is widely used for APIs and integrations.  
- CSV is popular for spreadsheets and data analysis tools.  
- NDJSON (newline-delimited JSON) works well for very large datasets.  

**What success looks like**
- Users can choose the output format with a simple flag.  
- The scraper writes results to a file instead of just the screen.  
- Data can be loaded into Excel, Python, or visualization tools easily.  


---

## 9. Improve test coverage & add integration tests
**Why:** Prevent regressions & validate crawling logic.  
**Where:** `normalize_url_test.go`, `get_urls_from_html_test.go`, new `integration_test.go`.

**How:**
- Unit tests: more cases (ports, fragments, broken HTML, `<base>`).
- Integration tests: use `httptest.Server` to serve fake pages.
- Assert crawler finds correct links, respects robots, retries on errors.

**What this means**
The scraper already has a few tests, but they only cover small pieces of logic. We need more tests that cover real-world scenarios, including full crawling flows.

**Why it matters**
- Prevents future changes from breaking existing functionality.  
- Builds confidence that the scraper works correctly on real websites.  
- Makes the project more maintainable for multiple contributors.  

**What success looks like**
- Unit tests cover edge cases (URLs, HTML quirks).  
- Integration tests simulate crawling a fake website and verify correct results.  
- Test coverage is reported in CI so the team can track improvements.  


---

## 10. Add GitHub Actions CI workflow
**Why:** Automatic validation on PRs.  
**Where:** `.github/workflows/ci.yml`.

**How:**
- Run on `push` and `pull_request`.
- Steps:
  - `go test ./...`
  - `gofmt -l .` (fail if output)
  - `go vet ./...`
  - (optional) `golangci-lint run ./...`
- Add CI badge to README.

**What this means**
Every time someone pushes code or opens a pull request, GitHub can automatically run tests and checks. This is called Continuous Integration (CI).

**Why it matters**
- Ensures broken code never lands in the main branch.  
- Saves developer time by catching problems early.  
- Shows external users that the project is well-maintained.  

**What success looks like**
- A GitHub Actions pipeline runs tests, formatting checks, and linting.  
- The results are visible on each pull request.  
- A badge in the README shows build status (green = good).  


---

# Suggested Implementation Order
1. README (item 1)  
2. CLI flags (item 2)  
3. Robust HTTP client (item 4)  
4. Robots support + per-host limiter (items 3 + 5)  
5. URL normalization (item 6)  
6. SQLite resume (item 7)  
7. Output adapters (item 8)  
8. Tests (item 9)  
9. CI (item 10)  

---
