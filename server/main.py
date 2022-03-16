# main.py
import random
from fastapi import FastAPI

app = FastAPI()

# xkcd
# 1-2591 

@app.get("/item/")
def item():
    try:
        return random.randrange(1, 2591)
    except:
        return "error"