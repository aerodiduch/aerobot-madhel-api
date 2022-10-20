# Madhel API
Este servicio nace de la necesidad de tener un **flujo de datos estables** de la **API del MADHEL de ANAC**. Debido a que el anterior mencionado servicio sufre de considerables fluctuaciones y caidas por periodos largos de tiempo, decidí construir una alternativa que no rompe todo una vez que se cae.

Está construido en dos partes, en la primera es un cliente también en **Go** que realiza peticiones sobre los **700 aeropuertos** y actualiza una **base de datos Redis**, la cual es consumida por esta API. La frecuencia de actualización es **cada 3 dias**, más que suficiente ya los aerodromos **no presentan grandes cambios.**

## Uso de la API
----
Actualmente la API es de uso **privado** y la información es solo accesible mediante una **API KEY**. Si estás interesado en acceder a este servicio, podes contactarme para que pueda propocionarte una **API KEY** para tu proyecto -> **aerodiduch@gmail.com**

> Este servicio está *desplegado a producción* y puede ser accedido a través del siguiente endpoint mediante el método **GET**:
```
https://madhel.aerobot.com.ar/json/{CODE}
```
Siendo ```{CODE}``` el código FAA del aeródromo en cuestión, por ejemplo ```ATE```.

La respuesta de la API es la misma que la que otorga ANAC, esto es asi para evitar modificaciones en el código y maximizar tanto la eficiencia como la escalabilidad. 

Ejemplo de respuesta de la API para ```ATE```. 
> Algunos campos del siguiente ejemplo han sido censurados para no exponer **datos de personas**.
```json
{
    "data": {
        "vor_verification_point": "",
        "telephone": []
        "sei_category": null,
        "fato": [],
        "rwy_declared_distances": [],
        "ats": [],
        "human_readable_localization": "340703S 0590456W - 6 KM al SW de la ciudad de ZÁRATE (Pcia. de BUENOS AIRES) - ELEV 26 M 85 FT.",
        "fuel": "",
        "fax": "",
        "tlof": [],
        "atz": "",
        "thr": [
            "27 MOV 150 M - RWY AVBL 1200 M por OBST."
        ],
        "twy": [],
        "local": "ATE",
        "helpers_system": {
            "visual": "",
            "radio": []
        },
        "parking_lots": "",
        "human_readable_identifier": "ZÁRATE - (ATE) - DRCE - PÚBLICO NO CONTROLADO",
        "apn": [],
        "rwy": [
            "01/19 1200x40 M - Tierra.",
            "09/27 1350x40 M - Tierra."
        ],
        "service_schedule": "",
        "norms": {
            "general": {
                "content": "Las OPS VFR deberán ajustarse a lo establecido en el ANEXO BRAVO (Ver AIP VOL. I - ENR 1.1-3).",
                "related_documents": [
                    {
                        "uri": "....",
                        "description": "NORMAS GENERALES PARA OPERACIONES EN AERÓDROMOS UBICADOS DEBAJO DE ÁREAS DE CONTROL TERMINAL (ENR 1.1.5)",
                        "title": "ANEXO BRAVO"
                    }
                ]
            },
            "particular": {
                "content": "ACFT con motor deberán realizar el circuito de tránsito con virajes por izquierda a 1000 M del AD.\r\nCTN, actividad habitual de planeadores remolcados por avión sobre la RWY hasta 2000 FT. Frecuencia de comunicación informativa sobre actividad de planeadores 123,00 MHz.",
                "related_documents": []
            }
        }
    },
    "pdf_share": "/madhel/api/v2/airports/ATE/pdf?format=json",
    "the_geom": {
        "type": "Feature",
        "geometry": {
            "type": "Point",
            "coordinates": [
                -59.08266,
                -34.11829
            ]
        },
        "properties": {
            "gg_point_coordinates": [
                -34.11829,
                -59.08266
            ],
            "name": "ZÁRATE - (ATE) - DRCE - PÚBLICO NO CONTROLADO"
        }
    },
    "updated_at": "2022-07-22T17:57:55Z",
    "type": "AD",
    "human_readable_identifier": "ZÁRATE - (ATE) - DRCE - PÚBLICO NO CONTROLADO",
    "metadata": {
        "localization": {
            "distance_reference": "6",
            "state": "BUENOS AIRES",
            "region": "RACE",
            "fir": "SAEF",
            "direction_reference": "SO",
            "coordinates": {
                "lat": -34.11829,
                "lng": -59.08266
            },
            "elevation": 26,
            "city_reference": "Zárate"
        },
        "reference_sketch": "",
        "identifiers": {
            "iata": null,
            "icao": null,
            "local": "ATE"
        },
        "control": "NON-CONTROLLED",
        "condition": "PUBLICO",
        "usage": null,
        "status": "OK",
        "ansp": null,
        "sna": false,
        "traffic": "NTL",
        "is_granted": false
    },
    "notam": []
}
```

# ¿Por qué usar este servicio?
Tal como mencionaba el principio, estos datos **no son en tiempo real** si no que tiene un delay de 2-3 dias con respecto a los proporcionados de **ANAC**. Pero visto y considerado que los aerodromos no se mudan repentinamente de provincia y los cambios que pueden haber muchas veces son minimos, es una gran alternativa para no estar pendientes de las caidas del servicio original.

Por otro lado, este servicio **tiene un certificado verificado de seguridad** (SSL), mientras que el servicio proporcionado por ANAC **presenta fallas continuas con respecto a esto** y muchas veces las peticiones fallan porque no se puede verificar el certificado en cuestión. 

Ventajas con respecto al servicio de **ANAC**:
--
- Tiempo de respuesta significativamente menor **30 ms** vs **75-90 ms**.
- No hay problemas con el **certificado SSL**.
- **Alta disponibilidad del servicio.**

Desventajas
--
- La información no es en tiempo real.
