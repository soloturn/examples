import os
import requests

# Set the API endpoint and your API key
api_endpoint = "https://api.openai.com/v1/images/generations"
api_key = os.environ["OPENAI_API_KEY"]

# Set the image data and metadata
data = {
    "prompt": "jumping cat flash",
    "n": 1,
    "size": "1024x1024"
}

# Set the headers
headers = {
    "Authorization": f"Bearer {api_key}"
}

# Send a POST request to the API to create the image
response = requests.post(api_endpoint, json=data, headers=headers)

# Print the response
print(response.json())
