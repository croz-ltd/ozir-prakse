from flask import Flask, render_template
import os

app = Flask(__name__)


@app.route('/')
def index():
    link = os.environ.get('BACKEND_URL')

    if not link:
        link = 'http://localhost:8000'

    return render_template('index.html', backend_url=link)


@app.route('/healthz')
def healthz():
    return 'ok'


@app.route('/readyz')
def readyz():
    return 'ok'
