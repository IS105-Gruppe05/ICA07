# ICA07
 
Navnliste: Abdikani Gureye, Brede Knutsen Meli, Eirik Aanestad Fintland, Jan Kevin Henriksen, Mats Skjærvik, Mikael Kimerud, Morten Johan Mygland, Nils Fredrik Iselvmo Bjerk.

- I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor vil det ikke alltid være at alle har pushet opp noe til github.

Midlertidig innlevering, vi er enda ikke ferdige med alle oppgavene i ICA07.

1 a og b) 

![Bilde1](https://raw.githubusercontent.com/IS105-Gruppe05/ICA07/master/Bilder/Bilde1.png)

c)
I:


![Bilde2](https://github.com/IS105-Gruppe05/ICA07/blob/master/Bilder/Bilde2.png?raw=true)

![Bilde3](https://github.com/IS105-Gruppe05/ICA07/blob/master/Bilder/Bilde3.png?raw=true)
 
 









1)	Ca 20% for selve transporten (Spør janis om hjelp)

![Bilde4](https://github.com/IS105-Gruppe05/ICA07/blob/master/Bilder/Bilde4.png?raw=true)
 
2) Mange ulike tjenester begrenser de største UDP pakkene til rundt 576 bytes ( som for eksempel dns). Det er mulig å gå over ca 500 bytes men da risikerer man packet loss og andre problemer har større sannsynlighet for å oppstå. 512 bytes er en fin størrelse, som blir brukt flere steder og dermed takler ulike rutere denne pakkestørrelsen. For IPv6 kan pakkene være av større størrelse på rundt 1472-1500. Det er viktig å trekke fra IPV4/IPv6 header, og UDP header for begge IP typene.

ii: Over NIC.
1)
2)
3)	
















Oppgave 2 (TCP)
a) 

![Bilde5](https://github.com/IS105-Gruppe05/ICA07/blob/master/Bilder/Bilde5.png?raw=true)

b) 

i) De største forskjellene mellom UDP og TCP er at UDP ikke trenger tilkobling, det vil si at når data eller meldinger sendes, vet man ikke om den kommer frem eller blir tapt på veien mellom sender og mottaker. TCP er tilkoblings-orientert, som vil si at når en melding blir sendt så vil meldingen komme frem så lenge tilkoblingen ikke blir brutt. Hvis tilkoblingen skulle vise seg å bli brutt så vil serveren fremdeles spørre etter den tapte delen.

Når det gjelder rekkefølge av meldinger, så vil man med TCP ikke måtte bekymre seg over at data kan forekomme i feil rekkefølge. med UDP så vet man ikke hvilken melding som ble sendt ut først, hvis man for eksempel sender ut to meldinger.

Om et program bruker TCP eller UDP er opp til utviklerne, og kommer an på hva programmet trenger. Ønsker du feil-korrigering og en sterk tilkobling så er TCP det bedre valget. UDP brukes der hastighet er viktig, og nettverket/ operativsystemet ikke trenger å gjøre mye for å oversette data som kommer fra pakkene.


ii) En TCP pakke kan være på rundt 1500 bytes. Blir den større enn det, kan det by på flere problemer som tap av data, pakke forkortelse osv. Det er som regel lurt å være under den øverste grensen på 1500, siden man må regne med data for Ip/TCP header ( 20 bytes hver ca)  osv som kan gjøre at du kommer over toppen. 1400 bytes før man legger til bytes for IP og TCP header, gjør at man kommer under grensen der problemer lettere kan oppstå. Hvis man er over 1500 bytes kan en løsning være å dele opp i to pakker i stedet. 

NB: Nå skrev jeg pakker som jo eksisterer for TCP, men at man ikke kan se de. TCP håndterer data i form av segmenter.

iii) Fragmentering oppstår når lagringsplass ikke brukes effektivt, og reduserer da både kapasitet og opptreden. Hvordan fragmentering oppstår, avhenger av hvordan datamaskinen du bruker lagrer filer. Hvis en fil som behøver å lagres er for stor til å passe de ledige plassene i driven, som for eksempel en stor filmfil på mange GB. Når dette skjer, så må den delen av operativsystemet som kontroller hvordan filer lagres ta en avgjørelse om hvor filen skal lagres, og noen ganger er den beste avgjørelsen og dele filen opp i mange mindre deler. Man kan løse fragmenteringsproblemer ved å sette til side plass om det skulle trengs. (Pre-Allocation).

iv)
TCP: Sending og mottak av email, SMS service, musikkstreaming.

UDP:  - Voice over IP. Eksempel: Du bruker Skype/Viber og har en videokonferanse med kollegaer. Du ender opp med å miste pakker, men kan fortsatt høre vedkommende selv om det stopper/hakker litt i samtalen. Det samme eksemplet gjelder også for video spill online og streaming av filmer.


Oppgave 3:
a)
b)
