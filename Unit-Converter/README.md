# Unit Converter

A simple web application that converts between different units of measurement including length, weight, and temperature.

## Features

- Convert between different units of measurement:
  - **Length**: meters, kilometers, centimeters, millimeters, feet, inches, miles, yards
  - **Weight**: grams, milligrams, kilograms, ounces, pounds
  - **Temperature**: Celsius, Fahrenheit, Kelvin
- User-friendly interface with dropdown selectors
- Real-time conversion results
- Client-server architecture with REST API

## Technologies Used

- Frontend: HTML, CSS, JavaScript
- Backend: Node.js with native HTTP module
- No external dependencies or databases

## Installation

1. Clone the repository:  
    git clone [your-repository-url]
2. Navigate to the project directory:  
    cd unit-converter
3. Start the server:  
    node server.js
4. Open your browser and visit:  
    http://localhost:3000

## Project Structure

    unit-converter/
    ├── client/
    │   ├── client.html      # Main HTML file
    │   └── style.css        # CSS styles
    └── server.js            # Node.js server with conversion logic

## Usage

1. Select the conversion type (Length, Weight, or Temperature) from the dropdown  
2. Enter the value you want to convert  
3. Select the unit you're converting from  
4. Select the unit you're converting to  
5. Click the "Convert" button  
6. View the converted value displayed below the button  

## API Endpoint

The application uses a single API endpoint for conversions:

    POST /api/units

Request body:

    {
      "type": "length|weight|temperature",
      "value": number,
      "from": "unit-name",
      "to": "unit-name"
    }

Response:

    {
      "success": boolean,
      "convertedValue": number,
      "originalData": object
    }

## Conversion Logic

The application handles three types of conversions:

### Length and Weight
- Uses multiplicative conversion factors based on standard units  
- Converts through an intermediate standard unit (meters for length, grams for weight)  

### Temperature
- Uses functional conversions with Celsius as the intermediate unit  
- Implements the standard temperature conversion formulas  

## Error Handling
- Validates input to ensure positive numeric values  
- The server returns appropriate HTTP status codes for errors  
- Basic error messages are displayed to the user  
