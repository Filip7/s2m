# s2m

> [!WARNING]
> Work in progress, currently this does not work

Single to multi line SQL

Goal of this project is to turn this

```sql
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy');
INSERT INTO films ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
```

to this

```sql
INSERT INTO films (code, title, did, date_prod, kind) VALUES ('B6717', 'Tampopo', 110, '1985-02-10', 'Comedy'),
                                                             ('HG120', 'The Dinner Game', 140, DEFAULT, 'Comedy');
```

Examples from this stack overflow [question](https://stackoverflow.com/a/28168010)

## Do I need this?

Actually no, if you use Intellij/DataGrip or any other Jetbrains IDE, the IDE can do that for you.  
If you are using pg_dump to generate an sql dump, and want to speed up the queries by turning them into multi line insets, then just define rows per insert flag `--rows-per-insert=100` <- insert number that works for your use case.

Read more here [https://www.postgresql.org/docs/current/app-pgdump.html](https://www.postgresql.org/docs/current/app-pgdump.html)
