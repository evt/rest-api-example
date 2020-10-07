![Github CI/CD](https://img.shields.io/github/workflow/status/evt/rest-api-example/Go)
![Repository Top Language](https://img.shields.io/github/languages/top/evt/rest-api-example)
![Scrutinizer Code Quality](https://img.shields.io/scrutinizer/quality/g/evt/rest-api-example/main)
![Codacy Grade](https://img.shields.io/codacy/grade/c9467ed47e064b1981e53862d0286d65)
![Github Repository Size](https://img.shields.io/github/repo-size/evt/rest-api-example)
![Github Open Issues](https://img.shields.io/github/issues/evt/rest-api-example)
![Lines of code](https://img.shields.io/tokei/lines/github/evt/rest-api-example)
[![License](https://img.shields.io/badge/license-MIT-green)](./LICENSE.txt)

<img align="right" width="50%" src="./images/big-gopher.jpg">

# REST API Server example
Ladies and gentlemen, once upon a time I asked a candidate for Golang Junior Developer position to create a simple Golang REST API server doing something useful (anything useful).

I won't show his code here :) But as a result I realized I don't have my answer to this question :) So, here it is!

This is Golang REST API server example including the following features:
* based on high performance, extensible, minimalist Go web framework - Echo - <https://echo.labstack.com> 
* made with Clean Architecture in mind (controller -> service -> repository)
* has services that work with both PostgreSQL database (user CRUD) and Google Cloud Storage (file upload/download)
* includes service & controller go tests based on database mocks auto-generated with go:generate and mockery (<https://github.com/vektra/mockery>)
* Postman tests included aswell
* config based on envconfig (<https://github.com/kelseyhightower/envconfig>)
* made relatively fast :) but with :heart: :)
* more details in my Youtube video that I'm going to make tomorrow :)
