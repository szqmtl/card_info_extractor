# Card Information Extractor

This is a small utility for practising the gemini image recognization to extract the text information from a card image, such as an ID card. 
However, for a product processing ID cards, there'd be some security concerns with this way. Therefore, for ID card management, please consider other solutions, 
such as Google Cloud Document AI(DocAI), AWS Intelligent Document Processing(IDP), etc.  

## Configuration

In order to configure GEMINI API, set the environment variable.

```aiignore
GEMINI_API_KEY=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

## Output

If gemini successfully extracts the card info, the application prints out something like:

```aiignore
card : &{Id:123-456-789 Name:JANE DOE Issuer:Library, Quebec Expiration:2016/08/02}
```

otherwise, it shows like:

```aiignore
2025/12/07 09:07:05 XXXXXX.png is not a valid Identity card
```