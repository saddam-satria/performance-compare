
from fastapi import APIRouter
from components.post.repository import getPosts
from fastapi.responses import JSONResponse
from http import HTTPStatus

postRoute = APIRouter(prefix="/api/v1")


@postRoute.get("/posts")
def getAPIPosts():
    posts = getPosts()
    

    content = {
        "total" : len(posts),
        "message": "welcome to fastapi - by Saddam", 
        "statusCode" : HTTPStatus.OK, 
        "data": {
            "posts" : posts
        }
    }

    return JSONResponse(content=content)