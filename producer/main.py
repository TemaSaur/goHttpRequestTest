from fastapi import FastAPI

app = FastAPI()


@app.get("/")
def data():
    return [
        {
            "key1": 0,
            "key2": "string",
            "values": [0, 1, 2],
        },
        {
            "key1": 0,
            "key2": "string2",
            "values": [],
        }
    ]
