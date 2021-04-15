# vaccine-progress

Dette er min lille Twitter "bot" der hver dag vil tweete hvor mange procent af Danmarks befolkning der er blevet vaccineret. Du kan finde den på [@VaccineDk](https://twitter.com/VaccineDk).

## Data
Jeg ville allerhelst have fået mit data direkte fra Statens Serum Institut, men åbenbart så er SSI ikke helt med på noderne endnu, så den eneste måde jeg kunne få data fra dem, var ved at downloade en zip hvor data så er i csv format. Siden kan ses [her](https://covid19.ssi.dk/overvagningsdata/download-fil-med-vaccinationsdata).

Det kunne vi godt arbejde med, HVIS jeg så kunne downloade filen på en smart måde, men igen så har SSI været så smart at smide et lille random id på alle deres filer f.eks.:
- [covid19-vaccinationsdata-15042021-1hi7](https://files.ssi.dk/covid19/vaccinationsdata/zipfil/covid19-vaccinationsdata-15042021-1hi7)
- [covid19-vaccinationsdata-14042021-23uu](https://files.ssi.dk/covid19/vaccinationsdata/zipfil/covid19-vaccinationsdata-14042021-23uu)
- [covid19-vaccinationsdata-13042021-2tal](https://files.ssi.dk/covid19/vaccinationsdata/zipfil/covid19-vaccinationsdata-13042021-2tal)

Så det kunne ikke rigtig bruges, da jeg ikke kommer til at manuelt downloade det data hver dag. Jeg skrev en e-mail til SSI for at høre om det var muligt at få det data via et API eller i JSON format uden at skulle downloade en zip fil.

**Hvad gør man så?**

Efter en hurtig Google søgning, så er der masser af danske nyhedssider der har en dedikeret side til at vise dagens vaccinationstal. Jeg tog den første der virkede nogenlunde troværdig, det blev [TV2](https://nyheder.tv2.dk/samfund/hvor-mange-er-vaccineret-i-danmark-nyeste-vaccinetal-overblik). Efter kort tid i Chrome Developer Tool fandt jeg frem til at de har deres data i en fin lille JSON fil - https://static.editorialdev.tv2a.dk/assets/2020/covid19/vaccinesDashboard.json

```json
{
  "Danmark": {
    "vaccinestarted": "1.007.166",
    "vaccinestartedper100": "17,2",
    "vaccinedone": "479.185",
    "vaccinedoneper100": "8,2",
    "dosestotal": "1.486.351",
    "dosestotalper100": "25,4",
    "dosesdaily": "22.724",
    "updated": "15. april, 14.00"
  },
  "World": {
    "vaccinestarted": "468.036.108",
    "vaccinestartedper100": "6",
    "fullyvaccinated": "184.654.423",
    "fullyvaccinatedper100": "2,37",
    "updated": "15. april, 18.18"
  },
  "disclaimer": "Flokimmunitet kræver, at 60-70 procent er immune, har SSI oplyst i december 2020."
}
```

## Koden
Så skulle der bare skrive lidt kode så vi kunne lave vores Twitter bot - endelig !

Kort fortalt så gør koden følgende:

1. Læs JSON filen
2. Find det data der er relevant for os
3. Lav en "progress bar" i tekst format så vi kan tweete det.
4. Tweet det !
