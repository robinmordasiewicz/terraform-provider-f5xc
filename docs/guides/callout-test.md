---
page_title: "Callout Rendering Test"
subcategory: "Testing"
description: |-
  Test page for various callout and bullet list rendering combinations.
---

# Callout Rendering Test

This page tests various markdown syntax combinations for Note callouts with bullet lists.

---

## Test 1: No blank line, no indent

-> **Note:** Only one of the following may be set:
- `attr1` - (Optional) Description one.
- `attr2` - (Optional) Description two.
- `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 2: Blank line, no indent

-> **Note:** Only one of the following may be set:

- `attr1` - (Optional) Description one.
- `attr2` - (Optional) Description two.
- `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 3: Blank line, 2-space indent

-> **Note:** Only one of the following may be set:

  - `attr1` - (Optional) Description one.
  - `attr2` - (Optional) Description two.
  - `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 4: Blank line, 3-space indent

-> **Note:** Only one of the following may be set:

   - `attr1` - (Optional) Description one.
   - `attr2` - (Optional) Description two.
   - `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 5: Blank line, 4-space indent

-> **Note:** Only one of the following may be set:

    - `attr1` - (Optional) Description one.
    - `attr2` - (Optional) Description two.
    - `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 6: Blank line, 5-space indent

-> **Note:** Only one of the following may be set:

     - `attr1` - (Optional) Description one.
     - `attr2` - (Optional) Description two.
     - `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 7: No blank line, 4-space indent

-> **Note:** Only one of the following may be set:
    - `attr1` - (Optional) Description one.
    - `attr2` - (Optional) Description two.
    - `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 8: Asterisk bullets, no indent

-> **Note:** Only one of the following may be set:

* `attr1` - (Optional) Description one.
* `attr2` - (Optional) Description two.
* `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 9: Plus bullets, no indent

-> **Note:** Only one of the following may be set:

+ `attr1` - (Optional) Description one.
+ `attr2` - (Optional) Description two.
+ `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 10: HTML unordered list

-> **Note:** Only one of the following may be set:

<ul>
<li><code>attr1</code> - (Optional) Description one.</li>
<li><code>attr2</code> - (Optional) Description two.</li>
<li><code>attr3</code> - (Optional) Description three.</li>
</ul>

Regular paragraph after.

---

## Test 11: Inline list in callout

-> **Note:** Only one of the following may be set: `attr1`, `attr2`, or `attr3`.

Regular paragraph after.

---

## Test 12: Numbered list, no indent

-> **Note:** Only one of the following may be set:

1. `attr1` - (Optional) Description one.
2. `attr2` - (Optional) Description two.
3. `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 13: Tab indent (single tab)

-> **Note:** Only one of the following may be set:

	- `attr1` - (Optional) Description one.
	- `attr2` - (Optional) Description two.
	- `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 14: Blockquote style

-> **Note:** Only one of the following may be set:

> - `attr1` - (Optional) Description one.
> - `attr2` - (Optional) Description two.
> - `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 15: Warning callout with list

~> **Warning:** Only one of the following may be set:

- `attr1` - (Optional) Description one.
- `attr2` - (Optional) Description two.
- `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 16: Danger callout with list

!> **Danger:** Only one of the following may be set:

- `attr1` - (Optional) Description one.
- `attr2` - (Optional) Description two.
- `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 17: Multi-line callout attempt (soft break)

-> **Note:** Only one of the following may be set:\
- `attr1` - (Optional) Description one.\
- `attr2` - (Optional) Description two.

Regular paragraph after.

---

## Test 18: Two blank lines before list

-> **Note:** Only one of the following may be set:


- `attr1` - (Optional) Description one.
- `attr2` - (Optional) Description two.
- `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 19: Colon at end, immediate list

-> **Note:** Only one of the following may be set:
- `attr1` - (Optional) Description one.
- `attr2` - (Optional) Description two.

Regular paragraph after.

---

## Test 20: Definition list style

-> **Note:** Only one of the following may be set:

`attr1`
: (Optional) Description one.

`attr2`
: (Optional) Description two.

Regular paragraph after.

---

## Test 21: Separate note + list (no association)

-> **Note:** Only one of the following may be set.

Below are the options:

- `attr1` - (Optional) Description one.
- `attr2` - (Optional) Description two.
- `attr3` - (Optional) Description three.

Regular paragraph after.

---

## Test 22: Bold bullets no backticks

-> **Note:** Only one of the following may be set:

- **attr1** - (Optional) Description one.
- **attr2** - (Optional) Description two.
- **attr3** - (Optional) Description three.

Regular paragraph after.

---

## Test 23: Code block after callout (for comparison)

-> **Note:** This shows what a code block looks like:

```
attr1 - Description one
attr2 - Description two
attr3 - Description three
```

Regular paragraph after.

---

## Test 24: Single line list items

-> **Note:** Only one of the following may be set:

- `attr1` - Description one.

- `attr2` - Description two.

- `attr3` - Description three.

Regular paragraph after.

---

## Test 25: Tight list with descriptions on new lines

-> **Note:** Only one of the following may be set:

- `attr1`
  Description for attr1.
- `attr2`
  Description for attr2.
- `attr3`
  Description for attr3.

Regular paragraph after.

---

## Test 26: List with nested content (2-space continuation)

-> **Note:** Only one of the following may be set:

- `attr1` - (Optional) Description one.
  Additional details about attr1.
- `attr2` - (Optional) Description two.
  Additional details about attr2.

Regular paragraph after.

---

## Test 27: Mixed - callout then separate heading then list

-> **Note:** Only one of the following may be set.

### Options

- `attr1` - (Optional) Description one.
- `attr2` - (Optional) Description two.

Regular paragraph after.

---

## Test 28: Table instead of list

-> **Note:** Only one of the following may be set:

| Attribute | Type | Description |
|-----------|------|-------------|
| `attr1` | Optional | Description one. |
| `attr2` | Optional | Description two. |
| `attr3` | Optional | Description three. |

Regular paragraph after.

---

## Test 29: All inline in callout (wrapped)

-> **Note:** Only one of the following may be set: `attr1` (Optional) - Description one; `attr2` (Optional) - Description two; `attr3` (Optional) - Description three.

Regular paragraph after.

---

## Test 30: Paragraph bullets (no dash prefix)

-> **Note:** Only one of the following may be set:

`attr1` - (Optional) Description one.

`attr2` - (Optional) Description two.

`attr3` - (Optional) Description three.

Regular paragraph after.

---

## Summary

| Test | Syntax | Expected |
|------|--------|----------|
| 1 | No blank, no indent | List attached to callout? |
| 2 | Blank line, no indent | **Clean bullet list** |
| 3 | Blank line, 2-space | Indented list? |
| 4 | Blank line, 3-space | Indented list? |
| 5 | Blank line, 4-space | **CODE BLOCK** |
| 6 | Blank line, 5-space | Code block? |
| 7 | No blank, 4-space | Code block? |
| 8-9 | Asterisk/Plus bullets | Same as dash? |
| 10 | HTML ul/li | Raw HTML? |
| 11 | Inline in callout | Single paragraph |
| 12 | Numbered list | Ordered list? |
| 13 | Tab indent | Code block? |
| 14 | Blockquote | Quoted list? |
| 15-16 | Warning/Danger | Different callout colors |
| 17 | Soft break | Continued line? |
| 18 | Two blank lines | Separated? |
| 19-30 | Various | See results |
