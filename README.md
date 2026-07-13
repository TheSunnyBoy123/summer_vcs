# SOL VCS

A lightweight version control system built in Go, inspired by Git.

## Commands

### `init`

> **sol init** [-h|--help]

Creates a `.sol` repository in the current directory. Sets up the object store and reference directories.

### `add`

> **sol add** [files...]

Computes SHA-1 hashes of the given files, stores them as blob objects, builds a tree object, and records the tree hash in the staging area. If no files are given, all files in the working directory are added.

### `commit`

> **sol commit** -m <message>

Creates a new commit from the staged tree (or walks the working directory if nothing is staged). Requires `~/.solconfig` with `SOL_AUTHOR_NAME` and `SOL_AUTHOR_EMAIL`.

### `log`

> **sol log**

Displays the commit history starting from HEAD, following parent links.

### `cat-file`

> **sol cat-file** [-t|-p|-s] <object_sha>

Prints object information:
- `-t`  object type
- `-p`  pretty-print object content
- `-s`  object size

### `hash-file`

> **sol hash-file** [-w] <file>

Computes the SHA-1 hash of a file. With `-w`, writes the object to the database.

### `diff`

> **sol diff** <object1_sha> <object2_sha>

Shows the difference between two stored objects.

### `dissolve`

> **sol dissolve**

Removes the `.sol` repository directory and `.solignore` file (if present) from the current directory.

## Configuration

Create `~/.solconfig` with the following format:

```
SOL_AUTHOR_NAME=Your Name
SOL_AUTHOR_EMAIL=your@email.com
```

## Storage

Objects are stored in `.sol/objects/<first-2-chars>/<remaining-hash>` as zlib-compressed content formatted as `type size\x00<content>`.
