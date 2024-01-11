from config.database import session
from sqlalchemy import text

def getPosts():
    try:
        query = text('SELECT "public"."Post"."postId", "public"."Post"."title", "public"."Post"."body", "public"."Post"."createdAt", "public"."Post"."updatedAt", "public"."Author"."name" AS author_name FROM "public"."Post" INNER JOIN "public"."Author" ON "public"."Author"."authorId" = "public"."Post"."author_id"')
        result = session.execute(query).fetchall()
        session.commit()
    except Exception as e:
        session.rollback()

    result = [{
    "postId": post.postId,
    "title": post.title, "body": post.body, "createdAt": post.createdAt.strftime("%Y-%m-%d %H:%M:%S"), 
    "updatedAt": post.updatedAt.strftime("%Y-%m-%d %H:%M:%S"),
    "author_name": post.author_name
    } for post in result]

    return result
    