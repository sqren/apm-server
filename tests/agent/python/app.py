# -*- coding: utf-8 -*-

import os

from flask import Flask, render_template, request, redirect, url_for
from elasticapm.contrib.flask import ElasticAPM

app = Flask(__name__)
app.debug = True
app.config['APM'] = {
    'SERVERS': [os.environ.get('apm-server', 'http://localhost:8080')],
    'DEBUG': True,
}

apm = ElasticAPM(
    app,
    organization_id='123',
    app_id='123',
    secret_token='123',
)


@app.route('/')
def index():
    return 'OK'


@app.route('/error')
def error():
    raise ValueError(request.args.get('error', ''))


if __name__ == '__main__':
    port = int(os.environ.get('PORT', 5000))
    app.run(host=os.environ.get('HOST', 'localhost'), port=port)
