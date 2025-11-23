import json
import requests
url = "http://127.0.0.1:7860/api/v1/run/7497172d-9fa2-4ea8-b440-c73093886402"  # The complete API endpoint URL for this flow

# Request payload configuration
def execute(input_value):
    payload = {
        "input_value": input_value, 
        "output_type": "text",  # Specifies the expected output format
        "input_type": "text"  # Specifies the input format
    }

# Request headers
    headers = {
        "Content-Type": "application/json"
    }

    try:
        # Send API request
        response = requests.request("POST", url, json=payload, headers=headers)
        response.raise_for_status()  # Raise exception for bad status codes

        # Print response
        resp = json.loads(response.text)
        # print(resp["outputs"][0]["outputs"][0]["messages"]["message"])
        print(resp.get("outputs", [])[0].get("outputs", [])[0].get("messages", [])[0].get("message"))

    except requests.exceptions.RequestException as e:
        print(f"Error making API request: {e}")
    except ValueError as e:
        print(f"Error parsing response: {e}")
        
# execute("Um colega, ao meu lado, martelou o dedo")
# execute("estou com gripe, devo me preocupar?")
# execute("meu oculos de EPI está quebrado")
execute("a serralheira está com o disco de corte solto e com alguns parafusos frouxos")
