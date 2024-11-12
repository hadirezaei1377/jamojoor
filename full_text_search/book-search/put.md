 PUT /books
  {
    "mappings": {
      "properties": {
        "title": { "type": "text" },
        "author": { "type": "text" },
        "genre": { "type": "keyword" },
        "publish_date": { "type": "date" },
        "description": { "type": "text" }
      }
    }
  }