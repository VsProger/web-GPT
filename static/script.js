document.getElementById('chat-form').addEventListener('submit', function(event) {
    event.preventDefault();
    var userInput = document.getElementById('user-input').value;
    var responseArea = document.getElementById('chat-response');

    responseArea.innerHTML = 'Loading...'; // Display loading text while waiting for response

    // Make an AJAX request to your server
    fetch('http://localhost:8080/ask', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({prompt: "Hello, World!"})
    })
    .then(response => response.json())  // Make sure to parse the response as JSON
    .then(data => {
        console.log(data);  // Log the data to see what's received
        if (data.response) {
            console.log("Response from server:", data.response);
        } else {
            console.log("No response field in data:", data);
        }
    })
    .catch(error => console.error('Error:', error));
    
});


