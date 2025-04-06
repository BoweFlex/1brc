class CityRecord {
    constructor(cityName, temp, min, mean, max) {
        this.cityName = cityName;
        this.temp = temp;
        this.min = min;
        this.mean = mean;
        this.max = max;

    }
}

const fs = require('fs');

let data = fs.readFileSync('inputs/measurements_10k.txt', 'utf8')

const myArray = data.split("\n")




