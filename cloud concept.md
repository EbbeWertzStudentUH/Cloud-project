# Cloud project

## Idee

Mijn idee is een microSaaS (SaaS voor een specifieke niche van doelgroep):
Een productivity & management platform, maar voor kleine projecten zoals groepswerken in schoolopdrachten. Dit zou ik eventueel nog specifieker kunnen maken tot enkel software projecten.

## Inpsiratie

De inspiratie komt van het populair SaaS concept van managemnet & productivity platforms, zoals Monday of SmartSheet. Maar deze tools bieden enkel een meerwaarde bij grote teams en lange termijn projecten. Voor kleinere projecten is de setup en het onderhoud van agendas en schemas meer tijdconsumerend dan hoe tijdbesparend de tool is. Kortom: Gantt-chart gebaseerde planning, Kanban boards, requirements, en SCRUM principes zijn overkill bij zeer kleine groepen.

## Concreet

Het doel van mijn Micro-SaaS is dus management en planning voordelen te inspireren uit deze tools, en ze toe te passen in een veel kleinere footprint.

Ik realiseer dit in drie delen:
(Dit is een lijst van een hoop mogelijke ideeen. Het kan zijn dat er ideen bij komen of er geschrapt worden, maar dit is alvast een mogelijke versie)
### Planning
 - Gantt-chars worden vervangen door een agenda
    - Je kan projecten ingeven met titel, beschrijving, ... en natuurlijk een deadline
    - Teamleden kunnen tijden ingeven wanneer ze tijd hebben voor fysieke of online meetings
    - Eventueel kan er aan meer dan 1 projecten tegenlijk gewerkt worden
    - Sub-deadlines kunnen gemaakt worden om een project aan te pakken als verschillende iteraties met elk een eigen doel
 - Kanban boards worden vervangen door een simpele tabel
    - Elke project iteratie wordt in een lijst weergeven
    - Bij elke iteratie hoort een checklist met taken met naam beschrijving en een schatting van aantal uren
    - Teamleden kunnen openstaande taken kiezen als hun verantwoordelijkheid
    - Er is dan zichtbaar wie aan welke taak bezig is
    - Elk lid kan ook aanduiden dat een taak voltooid is.
    - Elke taak heeft een urenteller die elk lid kan verhogen om bij te houden hoe lang er aan dat deel gewerkt is.
    - Een lid kan ook problemen toevoegen op een taak. Er is dan een duidelijke indicator zichtbaar zodat andere leden kunnen helpen.
### Collaboration
 - Projecten en eventueel iteraties hebben een collaboration omgeving
   - Hier kunnen notities (eventueel markdown) gemaakt worden
   - Men kan een sessie starten om met meerdere leden te schrijven. Notities worden dan realtime ge-update
### Statistics
 - Van alle taken kan de timespan laten zien worden in een Gantt-chart. De Gantt is hier enkel ter visualisatie
 - Men kan zien per lid hoeveel taken er gedaan zijn of hoeveel uren er gewerkt zijn.
 - Eventueel kan op basis hierop machine learning toegepast worden om openstaande taken aan te raden voor specifieke personen.



## services

### Algemeen

- authenticatie (eventueel extern via OAuth)

### Planning
 - Agenda database (service met sub-services):
    - projecten (beschrijving, requirements, deadlines) -> [simpel: GraphQL/REST]
    - taakverdeling planner -> [functies: gRPC]
    - agenda querying serice -> [complexe objecten: SOAP]
        - agenda/kalender formaat
        - gantt chart of ander diagram formaat
        - eventuele REST wrapper om om te vormen naar een html view
### collaboration
 - LIVE (markdown) notities/samenvatting (service met sub-services):
     - Markdown to html parser (eventueel extern). [Puur tekst: REST]
     - Session service voor realtime berwerkingen. [Realtime: Websocket]

### Statistics
 - Taken en uren teller (kan sub-service van Planning service zijn)
 - Eventueel als het specifiek voor software-projecten wordt: github API voor commit en code frequency [support REST/GraphQL -> dus GraphQL]
 - Python machine learning model 

