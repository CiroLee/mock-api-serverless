# mock-api-server

serverless, written in go and deployed on vercel

## APIs

| Path          | Method | Parameters | Description                                          |
| ------------- | ------ | ---------- | ---------------------------------------------------- |
| /api/hello    | GET    |            | test api, return "hello world"                       |
| /api/consumer | GET    | num        | return consumer data, accept `num` to special amount |
