## How to run

1. Setup .env

```
GEMINI_API_KEY=your_api_key
```

2. Spin weaviate database

```bash
   docker-compose up -d
```

3. Install dependencies

```bash
   go mod download
```

4. Run with make

```bash
   make dev
```

## Endpoints

### POST /document

add a document to the knowledge

```json
{
  "sessionId": "id-to-group-documents",
  "documents": [
    {
      "text": "Paris is the capital and largest city of France. With an official estimated population of 2,102,650 residents in January 2023 in an area of more than 105 km2 (41 sq mi), Paris is the fourth-largest city in the European Union and the 30th most densely populated city in the world in 2022.Since the 17th century, Paris has been one of the world's major centres of finance, diplomacy, commerce, culture, fashion, and gastronomy. For its leading role in the arts and sciences, as well as its early and extensive system of street lighting, in the 19th century, it became known as the City of Light."
    }
  ]
}
```

### POST /query

ask question about documents

```json
{
  "sessionId": "id-to-group-documents",
  "query": "what is the capital of france?"
}
```

```json
{
  "data": "The capital of France is Paris.",
  "message": "Query processed successfully"
}
```
