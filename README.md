Configuration files are missing and need to be added manually.

```mkdir conf```

```touch app.ini```

The configuration format is as follows:

RUN_MODULE = debug

[app]
PAGE_SIZE = 10
JWT_SECRET = 312321$3213123213

[server]
HTTP_PORT = 8000
READ_TIMEOUT = 60
WRITE_TIMEOUT = 60

[database]
TYPE = mysql
USER = admin
PASSWORD = admin
#127.0.0.1:3306
HOST = 127.0.0.1:3306
NAME = mysql
TABLE_PREFIX = blog_
