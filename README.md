# CDP History
Nice History renderer

How to use it : 
- Launcher `./cdp-history --file /srv/www/cdp-history/history.json --host localhost:8888`
- We need to put the folder of templates next to the binary launch (cpd-history)

## Requirement
You need a json on this format : 
```
[
  {
    "date": "1550665031",
    "cdp": {
      "ID": 13021,
      "DaiDebt": 21287.95,
      "EthCol": 248.88,
      "Price": 145.18,
      "Ratio": 1.61,
      "DaiNet": 13745.99,
      "EthNet": 94.68
    }
  },
  {
    "date": "1550665184",
    "cdp": {
      "ID": 13021,
      "DaiDebt": 21287.95,
      "EthCol": 248.88,
      "Price": 145.18,
      "Ratio": 1.61,
      "DaiNet": 13745.99,
      "EthNet": 94.68
    }
  }
]
```

### Libraries
 - chart.js : https://www.chartjs.org/docs/latest/
