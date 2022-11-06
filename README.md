# service-move-to-school
Primus opiskelijarekisterin yhteydessä hyödynnettävä mikropalvelu, jonka avulla hakija-koululla oleva opiskelija siirretään, poistetaan tai arkistoidaan.

Palvelu on alunperin tarkoitettu ajettavaksi kontissa esim. Kubernetes klusterissa ajastettuna. Mikropalvelu on osa laajempaa sähköisen paikanvastaanoton kokonaisuutta.

Dockerfile, jonka sisälle voi upottaa primusqueryn ja käännetyn binäärin.

https://raw.githubusercontent.com/pasiol/dockerfile-buster-slim-pq/main/Dockerfile

![kaavio](images/sähköinen_paikanvastaanotto.png)

## Käyttö binäärinä

Kääntäminen

    make compile

---
    HOST=palvelimen_osoite PORT=NNNN ./bin/service-move-to-school
## Primus-tuontimääritykset

main.go

	accountToRemove         = ""
	accountToMoveConfig     = ""
	accountToArchieveConfig = ""

## Suodattimet hakija- ja opiskelijarekisteriin

query.go

- täydennä filterit riveille

    pq.Search = ""


## Salaisuudet

config/secrets.go