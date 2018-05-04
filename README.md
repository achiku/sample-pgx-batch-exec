# sample-pgx-batch-exec


```
\copy bulk_insert (val, num, created_at) from './data.csv' with csv
\copy bulk_insert (val, num, created_at) from './data.psv' (delimiter '|', format csv, header false);
```
