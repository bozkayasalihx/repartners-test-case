# Packaging Calculator


## Getting Started

```sh
git clone https://github.com/bozkayasalihx/repartners-test-case.git
```

## API Usage 

```sh 
curl -X POST http://159.65.241.100:3000/calculate 
   -H "Content-Type: application/json"
   -d '{ "pack_sizes": [250,500,1000,2000,5000],"ordered_packs": [251]'  
```

## Response 
```sh
{
  "251": "500 x 1"
}

```


## UI Usage 
* UI can accesable from `http://159.65.241.100:3000` this url
![screen shot](./screenshot.png?raw=true "UI screeshot")


## NOTE
UI written in static HTML; therefore, please refresh the page after each subsequent to remove previously inserted DOM Element
