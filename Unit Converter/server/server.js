const http = require("http");
const fs = require("fs");
const path = require("path");

var convertedData = {};

// Server
const server = http.createServer((req, res) => {
    if (req.method === "GET" && req.url === "/style.css") {
        const cssPath = path.join(__dirname, '../client/style.css');
        fs.readFile(cssPath, (err, data) => {
            if (err) {
                res.writeHead(404);
                res.end("CSS not found");
            } else {
                res.writeHead(200, { 'Content-Type': 'text/css' });
                res.end(data);
            }
        });
    }

    else if (req.method === "GET" && req.url === "/") {
            const htmlPath = path.join(__dirname, '../client/client.html');
            fs.readFile(htmlPath, (err, data) => {
            if (err) {
                res.writeHead(404);
                res.end("CSS not found");
            } else {
                res.writeHead(200, { 'Content-Type': 'text/html' });
                res.end(data);
            }
        })
    }

    else if (req.method === "POST" && req.url === "/api/units") {
        let body = '';
        req.on("data", (chunk) => body += chunk.toString());
        req.on("end", () => {
            try {
                const data = JSON.parse(body);
                console.log("Received:", data);
                convertedData = convert(data.type, data.value, data.from, data.to);
                console.log("Converted Data:", convertedData);
                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({ 
                    success: true, 
                    convertedValue: convertedData,
                    originalData: data
                }));
            } catch (error) {
                res.writeHead(400, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({ success: false, error: "Invalid JSON" }));
            }
        });
    }
    
    else {
        res.writeHead(404, { 'Content-Type': 'text/plain' });
        res.end('Page not found');
    }
});

server.listen(3000, () => console.log("Server running on http://localhost:3000"));

// Conversion Logic
const unitsData = {
    length: {
        meters: 1,          // 1 meter = 1 meter (base)
        kilometers: 0.001,  // 1 meter = 0.001 km
        centimeters: 100,   // 1 meter = 100 cm
        millimeters: 1000,  // 1 meter = 1000 mm
        feet: 3.28084,      // 1 meter ≈ 3.28084 ft
        inches: 39.3701,    // 1 meter ≈ 39.3701 in
        miles: 0.000621371, // 1 meter ≈ 0.000621371 mi
        yards: 1.09361      // 1 meter ≈ 1.09361 yd
    },
    weight: {
        grams : 1,
        milligrams: 1000,            
        kilograms: 0.001,   
        ounces: 0.03527396,  
        pounds: 0.00220462,      
    },
    temperature: {
        Celsius: {
            toFahrenheit: (c) => (c * 9/5) + 32,
            toKelvin: (c) => c + 273.15
        },
        Fahrenheit: {
            toCelsius: (f) => (f - 32) * 5/9,
            toKelvin: (f) => (f - 32) * 5/9 + 273.15
        },
        Kelvin: {
            toCelsius: (k) => k - 273.15,
            toFahrenheit: (k) => (k - 273.15) * 9/5 + 32
        }
    }
};

// General Case Conversion
function convert(unitCategory, value, fromUnit, toUnit) {
    if (unitCategory === 'temperature') {
        return convertTemperature(value, fromUnit, toUnit);
    }

    var conversionFrom = unitsData[unitCategory][fromUnit];
    var conversionTo = unitsData[unitCategory][toUnit];

    const originalValue = value / conversionFrom;
    return originalValue * conversionTo;
}

function convertTemperature(value, fromUnit, toUnit){
    var convertedVal = value;
    // Turn to Celsius First
    if (fromUnit === "Fahrenheit"){
        convertedVal = unitsData.temperature.Fahrenheit.toCelsius(convertedVal);
    }
    else if (fromUnit === "Kelvin"){
        convertedVal = unitsData.temperature.Kelvin.toCelsius(convertedVal);
    }

    // Then to target value
    if (toUnit === "Fahrenheit"){
        convertedVal = unitsData.temperature.Celsius.toFahrenheit(convertedVal);
    }
    else if (toUnit === "Kelvin"){
        convertedVal = unitsData.temperature.Celsius.toKelvin(convertedVal);
    }

    return convertedVal;
}

