from flask import Flask
from flask import request

import json

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def index():
    x = request.get_json()
    print(json.dumps(x, indent=2))
    return 'OK'

if __name__ == '__main__':
    app.run(debug=False)
