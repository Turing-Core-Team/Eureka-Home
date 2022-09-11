# Eureka Home


The Eureka project aims to organize and publish academic offers.
It is divided into three modules that work together

- 🔎 Eureka Scraper, which is in charge of extracting, organizing and storing possible academic offers. For this, it performs a web scraping and consumption of APIs from metadata from the websites of universities, government institutions and other pages that have academic offers.
- 🔔 Eureka Notification, which is in charge of notifying the project leaders of our organization, when an academic offer to which they can apply is published, for this it has metadata of the active projects.
- 🏠 Eureka Home, which is in charge of making available the list, filter and organization service of the academic offers in a Rest API that will be consumed by our [web page](https://turingbox.co/), so that from of filtersa and a description of the project, be able to list and organize the calls to which you could apply. This service is independent and asynchronous from Eureka Scraper and Eureka Notification, since they are triggered by a scheduled event.

## Context Map

The overall project context map

![ContextMap](https://raw.githubusercontent.com/Turing-Core-Team/StaticEurekaFiles/main/images/design_eureka_2.0-Context%20Map.drawio.png)

## Local EndPoints

### Opportunities

- **GET** [/eureka/v1.0/opportunities/filters/who/type/area/extra"]()

#### request

```json

{
  
}

```

## Filters

Filters are the way offers are searched through query parameters