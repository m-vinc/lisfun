# Lisfun

The main goal of this repository is to build a reliable starter project to bootstrap a new one using a really simple Go WEB stack: [Echo](https://echo.labstack.com/)/[Ent](https://entgo.io/)/[Templ](https://templ.guide/)/[HTMX](https://htmx.org/) and everything we need to build a fancy UI like plain [Tailwindcss](https://tailwindcss.com/) or something like [daisyUI](https://daisyui.com/).

I'll use that repository to explorate and learn htmx so consider every other branch than master like a sandbox.

## Requirements

First, you need installed on your dev environment :
* Make
* [Air](https://github.com/air-verse/air?tab=readme-ov-file#installation)
* [golangci-lint](https://golangci-lint.run/welcome/install/)
* [Templ](https://templ.guide/quick-start/installation)

## Developments

Use the `make watch` command to launch air which build and reload your app on every code changes, not require at first but very confortable. Because I don't need to manage configuration for now I add the require arguments to the Lisfun binary via the `.air.toml` configuration file.

```yaml
# /.air.toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = ["run", "--port", ":6116", "--log_level", "debug"]
  bin = "build/bin/lisfun_dev_darwin_arm64"
```

## What I'm building ?

This name is not random, maybe I've something in mind ...
