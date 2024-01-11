
from components.post.controller import postRoute

def routing(app):
    app.include_router(postRoute)