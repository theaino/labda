# LaBDA: <ins>La</ins>mbda <ins>B</ins>inded <ins>D</ins>efinition <ins>A</ins>pplicator

# Syntax

## Function definition (abstraction)

```
$arg.
```

Example:

```
$x.x
```

## Function application

```
function arg
```

Example:

```
($x.x) y
```

## Variables

With the above definitions, variables already work: just define an abstraction and apply it directly:

```
($varname.
/* Some expression */
) value
```

Since this doesn't really look good, there is an alternative:

## Wrapper

```
$varname=value.
/* Some expression */
```

This evaluates to exactly the same as the other variable example.

# Example

```
$true = $a.$b.a.
$false = $a.$b.b.

$or = $x.$y.x true (y true false).

$V = $VALID.VALID.       /* A valid result */
$IV = $INVALID.INVALID.  /* An invalid result */

or true false V IV
```

# Plans

- stdlib including [church encoding](https://en.wikipedia.org/wiki/Church_encoding)

# Ideas

- Compile to [de Bruijn Indexing](https://en.wikipedia.org/wiki/De_Bruijn_index)
- Optimize stdlib-church-integers using real integers
- Use deterministic behaviour for optimization (abstraction hash?)
