from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def root():
    return [{"plan": "basico", "price": 0}, {"plan": "intermedio", "price": 19}, {"plan": "premium", "price": 39}]