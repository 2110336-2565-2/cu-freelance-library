# About Utils Function

Utils lang function is the function that use to check or manipulate the language 

# Getting Start

## Initialize

```go
utilLang := NewUtilLang(languageList)
```

#### Parameters
| name         | description                                                           | example        |
|--------------|-----------------------------------------------------------------------|----------------|
| languageList | the list of languages that utils can be detected in `lingua.Language` | lingua.English |

## Functions

### CheckLang
the function that check language type

```go
lang := utilLang.CheckLang(text)
```

#### Parameters
| name | description    | example     |
|------|----------------|-------------|
| text | the input text | hello world |

#### Return

| name | description | example         |
|------|-------------|-----------------|
| lang | result lang | lingua.English  |
