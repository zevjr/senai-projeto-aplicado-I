import json
import requests
# url = "http://localhost:7860/api/v1/run/c807140c-a3f7-4e41-b196-e5977b8174f9"  # The complete API endpoint URL for this flow

url = "http://localhost:5000/ia/call"

# Request payload configuration
def execute(input_value):
    # payload = {
    #     "input_value": input_value, 
    #     "output_type": "text",  # Specifies the expected output format
    #     "input_type": "text"  # Specifies the input format
    # }
    payload = {"message": input_value}

# Request headers
    headers = {
        "Content-Type": "application/json",
        # "x-api-key": "sk-QaXGi4ts4UMQBozow31TL-3QoYYbv1Ph4waI9krGXaU"
    }

    try:
        # Send API request
        response = requests.request("POST", url, json=payload, headers=headers)
        response.raise_for_status()  # Raise exception for bad status codes
        # Print response
        # breakpoint()
        # resp = json.loads(response.text)
        # print(resp["outputs"][0]["outputs"][0]["messages"]["message"])
        # print(resp.get("outputs", [])[0].get("outputs", [])[0].get("messages", [])[0].get("message"))
        print(response.json())

    except requests.exceptions.RequestException as e:
        print(f"Error making API request: {e}")
    except ValueError as e:
        print(f"Error parsing response: {e}")
        
# execute("Um colega, ao meu lado, martelou o dedo")
# execute("estou com gripe, devo me preocupar?")
# execute("meu oculos de EPI está quebrado")
execute("a serralheira está com o disco de corte solto e com alguns parafusos frouxos")
