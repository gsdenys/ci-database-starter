# Database Starter Composite Action 

This is just a composite action to help start the persistence databases during the unit tests and benchmark. This starter usage are shown below.

```yaml
steps:
    - id: database
      uses: posteris/ci-database-starter@main
```

## PostgreSQL

* __Host:__     localhost
* __Port:__     5432
* __User:__     postgres
* __Password:__ postgres
* __Database:__ postgres


## MySQL

* __Host:__     localhost
* __Port:__     3306
* __User:__     root
* __Password:__ mysql
* __Database:__ mysqls