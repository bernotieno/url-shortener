function shortenUrl() {
    const longUrl = document.getElementById('longUrl').value;

    if (!longUrl.trim()) {
        document.getElementById('result').innerHTML = '<p class="error">Please enter a valid URL.</p>';
        return;
    }

    fetch('/api/shorten', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ url: longUrl })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('result').innerHTML = 
            `<p>Short URL: <a href="${data.short_url}" target="_blank">${data.short_url}</a></p>`;
    })
    .catch(error => {
        document.getElementById('result').innerHTML = 
            `<p class="error">Error: ${error.message}</p>`;
    });
}
