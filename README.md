# Eureka Home

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

The Eureka project aims to organize and publish academic offers.
It is divided into three modules that work together

- 🔎 Eureka Scraper, which is in charge of extracting, organizing and storing possible academic offers. For this, it
  performs a web scraping and consumption of APIs from metadata from the websites of universities, government
  institutions and other pages that have academic offers.
- 🔔 Eureka Notification, which is in charge of notifying the project leaders of our organization, when an academic
  offer to which they can apply is published, for this it has metadata of the active projects.
- 🏠 Eureka Home, which is in charge of making available the list, filter and organization service of the academic
  offers in a Rest API that will be consumed by our [web page](https://turingbox.co/), so that from of filtersa and a
  description of the project, be able to list and organize the calls to which you could apply. This service is
  independent and asynchronous from Eureka Scraper and Eureka Notification, since they are triggered by a scheduled
  event.

## Context Map

The overall project context map

![ContextMap](https://raw.githubusercontent.com/Turing-Core-Team/StaticEurekaFiles/main/images/design_eureka_2.0-Context%20Map.drawio.png)

## EndPoints

### Opportunities

- **GET** [/eureka/v1.0/opportunities/filters/:who/:type/:area/:extra"]()

#### request

```json
{
  "who": "personas",
  "type": "carreras",
  "area": "ingenieria",
  "extra": ""
}
```

#### response

```json
[
  {
    "tags": "ingenieria-sistemas-uniandes",
    "link": "https://mecanica.uniandes.edu.co/es/pregrado",
    "title": "Pregrado Ingeniería Mecánica",
    "requirements": "El proceso de admisión al programa de pregrado ...",
    "awards": "Título: Ingeniero Mecánico de la Universidad de los Andes",
    "description": "Nuestro principal objetivo es generar ...",
    "publication_date": "11 de mayo de 2021",
    "update_date": "11 se septiembre de 2022",
    "due_date": ""
  }
]
```

## Filters

Filters are the way offers are searched through query parameters

- The first filter (Who) refers to who is applying to the academic offer
#### Valid values
```
  "who": "personas"
  "who": "proyectos"
```

- The second filter (Type) refers to the type of application and is dependent on the first filter (Who).
Multiple types can be received by separating them by "-"

#### Valid values for person
```
  "type": "carreras"
  "type": "tecnologos"
  "type": "cursos"
  "type": "diplomados"
  "type": "carreras-tecnologos-cursos"
```

#### Valid values for projects
```
  "type": "concursos"
  "type": "eventos"
  "type": "divulgacion"
  "type": "financiacion"
  "type": "ferias"
  "type": "concursos-ferias"
```

- The third filter (Area), refers to the area of knowledge of the academic offer

#### Valid values
```
  "area": "ingenieria"
  "area": "programacion"
```

- The fourth filter (Extra), is optional and provides additional information for the order of the opportunities