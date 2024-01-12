from fastapi import FastAPI, Request
from fastapi.middleware.cors import CORSMiddleware
from config.database import engine
from config.routes import routing
from sqlalchemy.exc import SQLAlchemyError
import uvicorn
from psutil import Process, virtual_memory
import os

app = FastAPI(debug=True,docs_url="/api/documentations")
app.add_middleware(CORSMiddleware,  
                        allow_origins=["/"], 
                        allow_credentials=True,
                        allow_methods=["*"],
                        allow_headers=["*"]
                        )


@app.middleware("http")
async def  getUsageResource(request: Request, call_next):

    pid = os.getpid()
    process = Process(pid=pid)
    print(f"CPU usage {process.cpu_percent()} - Memory usage {round(process.memory_info().rss / (1024 * 1024))} MB")

    response = await call_next(request)

    return response

def main():
   
    try:
        engine.connect()
    except SQLAlchemyError as error:
        raise error
    
    routing(app)



    


@app.on_event("startup")
def startUp():
    main()
    