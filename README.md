# goldmark-SuperSubscript

This Go module contains two extensions for the Markdown parser [Goldmark](https://github.com/yuin/goldmark) introducing super- and subscripts.

## Syntax

Similar to [markdown-it](https://github.com/markdown-it/markdown-it) the new markup characters are:

- The circumflex `^` for superscript

- The tilde `~` for subscript

We need to place one before and one after the text segment like `H~2~O` or `x^2^`. 

**No whitespace** between the two surrounding markup characters is allowed. In case we explicitly want to use space we need to place a non-breaking whitespace---either directly or as HTML entity `&nbsp;`. 
