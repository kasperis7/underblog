Fork of [underblog](https://github.com/freetonik/underblog)

# Underblog

An extremely simple, fast static blog generator.

## How it works

You only need 4 things:

1. `index.html` template for blog's index page.
2. `post.html` template for single post.
3. `css/styles.css` for CSS styles.
3. `markdown` folder.

There is no front-matter. **Slug** is derived from the filename and **date** is the last modification date of the file. **Title** is derived from the first line of markdown file. Make sure the first line starts with `#`.

**Step 1:** create the following folder structure:

```
.
├── css
│   └── styles.css
├── markdown
│   └── Slug_1.md
│   └── Slug_2.md
│   └── Slug_3.md
├── index.html
├── post.html
```

(See [/example](example))

**Step 2:** run `underblog`.

**Step 3:** Your site is generated in `public`. Posts are ordered by last modification date.

## Live preview

Run underblog in watch mode:

```
underblog -watch
```

Then go to http://localhost:8080/. Changing markdown files will automatically re-generate the site.

## Features

- NO front matter
- NO themes
- NO JavaScript
- NO tags, categories, taxonomy
- NO template lookup logic
- NO plugins
- NO dependencies
---

## Building

In the root repository do:

`make build`
