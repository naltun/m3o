curl "https://api.m3o.com/v1/sentiment/Analyze" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $M3O_API_TOKEN" \
-d '{
  "text": "this is amazing"
}'