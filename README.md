# Card Information Extractor

This utility demonstrates a straightforward approach to using the Gemini API's image recognition capabilities. Its primary function is to extract text information (via OCR and document understanding) from an image, such as a sample ID card.

## Security Warning for Production Use
This utility is for demonstration and practice only.

Using a general-purpose Large Language Model (LLM) like Gemini for processing sensitive Identity Documents (IDs) in a production environment introduces significant security and compliance risks.

For enterprise-level, secure, and compliant ID card management and data extraction, please consider specialized document processing solutions:

* Google Cloud Document AI (DocAI)

* AWS Intelligent Document Processing (IDP)

* Other dedicated identity verification services

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