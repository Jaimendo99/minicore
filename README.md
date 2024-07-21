<p align="center">
  <a href="core.jaimendo.online" rel="noopener">
</p>

<h3 align="center">MINICORE</h3>


---

<p align="center"> A Web Development Project for UDLA's Web Development Course
    <br> 
    You can see the project live at <a href="https://minicore.jaimendo.online">minicore.jaimendo.online</a>
</p>

## 📝 Table of Contents


- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)


## 🏁 Getting Started <a name = "getting_started"></a>

This project uses the [GOTTH Stack](https://github.com/arejula27/goth-stack), which is:
- [Go Lang](https://golang.org/) as the backend
- [Echo](https://echo.labstack.com/) as the web framework
- [Templ](https://templ.guide/) as the template engine
- [TailwindCSS](https://tailwindcss.com/) for staying with tailwindUI
- [HTMX](https://htmx.org/) for the frontend magic

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

You will need to have 
[Go Lang](https://golang.org/) installed on your machine, follow the instructions on the website.
#### [Air:](https://github.com/cosmtrek/air) hot reload for Go

Air must be in path, you can install it by running the following command:
```bash
go install github.com/cosmtrek/air@latest
```
#### [Templ:](https://templ.guide/) template engine for Go
Templ must also be in PATH, you can install it by running the following command:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```
#### [Echo](https://echo.labstack.com/) web framework for Go
Echo is a go module, you can install it by running the following commands:
```bash
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
```


### Installing

The best way to self host this app is by using the dockerfile provided in the project, you can build the image by running the following command:

Clone the project
```bash
git clone Jaimendo99/talibox
```
Build the image
```bash
docker build -t talibox .
```
Run the image
```bash
docker run -p 8080:8080 talibox
```
