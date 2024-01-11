from config.constant import DB_URL
from sqlalchemy import create_engine
from sqlalchemy.orm import Session

engine = create_engine(DB_URL, pool_size=15, max_overflow=10)
session=Session(engine)
