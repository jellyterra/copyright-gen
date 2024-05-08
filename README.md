# copyright-gen
Automatically update copyright headers for source files.

It is simple and designed for sharing between different editors and plugins.

### Setup

```shell
$ go install github.com/jellyterra/copyright-gen
```

### Usage

```
Usage: copyright-gen profile source-file...
```

```shell
$ copyright-gen cxx.json main.cpp lib.cpp
```

## Template

The program uses Go template engine.

Predefined values:

- ```$.name``` System profile name
- ```$.year``` This month
- ```$.month``` This year
- ```$.day``` Today

For example:
```gotemplate
Copyright {{ $.year }} {{ $.name }}
Use of this source code form is governed under the MIT license.
```

## Profile

*Profile* is the generator configuration encoded in JSON.

```json
{
  "template": "Copyright {{ $.year }} {{ $.name }}\nUse of this source code form is governed under the MIT license."
}
```

> [!NOTE]
> Preset profiles are available in repo [copyright-gen-profiles](https://github.com/jellyterra/copyright-gen-profiles) .

**Java, C, C++** etc.
```json
{
  "prefix": "/*",
  "suffix": " */",
  "line_prefix": "    ",
  "template": "..."
}
```
```java
/*
    Copyright 2024 Jelly Terra
    Use of this source code form is governed under the MIT license.
 */

public class Main {}
```

**Go, Rust** etc.
```json
{
  "line_prefix": "// ",
  "template": "..."
}
```
```go
// Copyright 2024 Jelly Terra
// Use of this source code form is governed under the MIT license.

package main
```

**Bash, Python** etc.
```json
{
  "line_prefix": "# ",
  "template": "..."
}
```
```shell
# Copyright 2024 Jelly Terra
# Use of this source code form is governed under the MIT license.

echo ""
```

**HTML, XML** etc.
```json
{
  "line_prefix": "<!-- ",
  "line_suffix": " -->",
  "template": "..."
}
```
```html
<!-- Copyright 2024 Jelly Terra -->
<!-- Use of this source code form is governed under the MIT license. -->

<html lang="en"></html>
```
