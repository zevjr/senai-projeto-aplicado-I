from flask import Flask, request, jsonify
import json
import requests

app = Flask(__name__)

LANGFLOW_API_URL = "http://langflow:7860/api/v1/run/7497172d-9fa2-4ea8-b440-c73093886402"

def execute(input_value):
    payload = {
        "input_value": input_value,
        "output_type": "text",
        "input_type": "text"
    }

    headers = {
        "Content-Type": "application/json"
    }

    try:
        response = requests.post(LANGFLOW_API_URL, json=payload, headers=headers)
        response.raise_for_status()
        resp = response.json()
        message = resp.get("outputs", [])[0].get("outputs", [])[0].get("messages", [])[0].get("message")
        return message
    except requests.exceptions.RequestException as e:
        return f"Error making API request: {e}"
    except (ValueError, IndexError, KeyError) as e:
        return f"Error parsing response: {e}"

@app.route("/ia/call", methods=["POST"])
def call_ia():
    data = request.get_json()
    input_text = data.get("message")

    if not input_text:
        return jsonify({"error": "Missing 'message' in request body"}), 400

    result = execute(input_text)
    return jsonify({"response": result})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
