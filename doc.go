/*
Fragmenta provides a set of independent libraries for building golang web applications, and a command line tool for generating, developing and running apps. 

The following subcommands are available when using the command line fragmenta tool:

    ------
      fragmenta version -> display version
      fragmenta help -> display help
      fragmenta new [app|cms|blog|go gettable URL] path/to/app -> creates a new app from the repository at URL at the path supplied
      fragmenta -> builds and runs a fragmenta app
      fragmenta server -> builds and runs a fragmenta app
      fragmenta test  -> run tests
      fragmenta migrate -> runs new sql migrations in db/migrate
      fragmenta backup [development|production|test] -> backup the database to db/backup
      fragmenta restore [development|production|test] -> backup the database from latest file in db/backup
      fragmenta deploy [development|production|test] -> build and deploy using bin/deploy
      fragmenta generate resource [name] [fieldname]:[fieldtype]* -> creates resource CRUD actions and views
      fragmenta generate migration [name] -> creates a new named sql migration in db/migrate
    ------


These let you create a new website, run the website, run tests on all packages, migrate, backup and restore the database, deploy the app to production, and generate resources (model, actions, views, migration) and migrations. 


Fragmenta also provides a suite of libraries for web development, which can be used independently:

        assets - an asset pipeline with minification and concatenation with asset fingerprinting
        auth - utilities for authentication and authorisation
        fragmenta - a command line tool for generating and developing websites
        model - a base model class for optional inclusion in models
        file - a package for handling file uploads
        validate - a package for handling field validation
        query - a query builder and result wrapper for mysql, psql and optionally sqlite
        router - a router which allows pattern matching, routes, redirects, filters and provides a handler interface
        server - a simple server based on http.listenandserve
        view - a library for rendering view templates using html/template
        helpers - helpers for form fields, currencies etc


The fragmenta tool requires the following two files to exist in a project directory:

server.go
secrets/fragmenta.json (app config)

otherwise the structure of your app is up to you. The default structure given in examples is to have a package per REST resource, so for example pages has a package under src, which contains the following folders/files

    src/pages
        actions -> 
        assets -> js, css, and images for pages
        views -> golang html/template files ending in .got
        pages.go -> the pages model
        pages_test.go -> tests for the pages model

These are suggested patterns and can be overridden by editing the files in lib/template in your project, which define exactly what is generated by fragmenta generate. 
*/
package main