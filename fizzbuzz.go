<!DOCTYPE html>
<html>
<head>
    <title>FizzBuzz Mini Site</title>
</head>
<body>
    <h1>FizzBuzz Mini Site</h1>
    <label for="number">Enter a number:</label>
    <input type="number" id="number" min="1" step="1">
    <button onclick="calculateFizzBuzz()">Calculate</button>
    <p id="result"></p>

    <script>
        function calculateFizzBuzz() {
            const number = parseInt(document.getElementById("number").value);
            let result = "";

            for (let i = 1; i <= number; i++) {
                if (i % 3 === 0 && i % 5 === 0) {
                    result += "FizzBuzz, ";
                } else if (i % 3 === 0) {
                    result += "Fizz, ";
                } else if (i % 5 === 0) {
                    result += "Buzz, ";
                } else {
                    result += i + ", ";
                }
            }

            result = result.slice(0, -2); // Remove the last comma and space
            document.getElementById("result").textContent = result;
        }
    </script>
</body>
</html>
