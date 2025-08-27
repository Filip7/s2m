# s2m

> [!WARNING]
> Work in progress, currently it can read from stdin and files, also it can save the output to a file
> BUT, I have not tested it on larger sql files, so there might be performance issues
> Will work on them

Single to multi line SQL

Goal of this project is to turn this

```sql
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
```

To this

```sql
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy'),
                                                             ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
```

Examples from this stack overflow [question](https://stackoverflow.com/a/28168010)

## Usage

```sh
./s2m "SELECT * from here; insert into mail(id) values (1);insert into mail(id) values (2);"
```

Or

```sh
./s2m "SELECT * from mail; \
INSERT into mail(id) values (1); \
INSERT into mail(id) values (2)"
```

Tool supports reading files and printing output to standard out.  
Use it like this:

```sh
./s2m -f export.sql
```

To save the output to file, pass `-o` command line flag with the name of the file.

```sh
./s2m -f export.sql -o output.sql
```

Also works with standard in

```sh
./s2m -o output.sql "SELECT * from mail; \
INSERT into mail(id) values (1); \
INSERT into mail(id) values (2)"
```

## Do I need this?

Actually no, if you use Intellij/DataGrip or any other Jetbrains IDE, the IDE can do that for you.  
If you are using pg_dump to generate an sql dump, and want to speed up the queries by turning them into multi line inserts, then just define rows per insert flag `--rows-per-insert=100` <- insert number that works for your use case.

Read more here [https://www.postgresql.org/docs/current/app-pgdump.html](https://www.postgresql.org/docs/current/app-pgdump.html)
