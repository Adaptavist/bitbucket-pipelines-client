from uuid import uuid4

from flask import Flask, render_template

app = Flask(__name__)


@app.route("/")
def home():
    return "stable"


@app.route('/2.0/repositories/<workspace>/<repo>/refs/tags')
def tags(workspace, repo):
    return render_template("tags.json", workspace=workspace, repo=repo)


@app.route('/2.0/repositories/<w>/<r>/pipelines/<i>')
def get_pipeline(w, r, i):
    return render_template("pipeline.json", workspace=w, repo=r, uuid=i)


@app.route('/2.0/repositories/<w>/<r>/pipelines/<i>/steps/')
def get_steps(w, r, i):
    return render_template("steps.json", workspace=w, repo=r, uuid='{' + uuid4().__str__() + '}')


@app.route('/2.0/repositories/<w>/<r>/pipelines/<pi>/steps/<si>/log')
def get_step_log(w, r, pi, si):
    return f'''
SOME LOG OUTPUT
    '''


@app.route('/2.0/repositories/<w>/<r>/pipelines/', methods=['POST'])
def post_pipeline(w, r):
    return get_pipeline(w, r, '{' + uuid4().__str__() + '}')
