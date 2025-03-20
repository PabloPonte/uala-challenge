import csv
import pymongo
import os
import random
from datetime import datetime
from dotenv import load_dotenv

# Load environment variables
load_dotenv('../.env')

# MongoDB connection info from .env file
MONGO_HOST = os.getenv("MONGO_HOST")
MONGO_PORT = os.getenv("MONGO_PORT")
MONGO_USER = os.getenv("MONGO_USER")
MONGO_PASSWORD = os.getenv("MONGO_PASSWORD")
DB_NAME = os.getenv("MONGO_DATABASE")

TWEETS_COLLECTION_NAME = "tweets"
FOLLOWS_COLLECTION_NAME = "follows"
CSV_FILE_PATH = "initial_data.csv" 

# URI build
MONGO_URI = f"mongodb://{MONGO_USER}:{MONGO_PASSWORD}@{MONGO_HOST}:{MONGO_PORT}/"

# Database connection
client = pymongo.MongoClient(MONGO_URI)
db = client[DB_NAME]
collection = db[TWEETS_COLLECTION_NAME]

# Drop the actual collection in order to recreate and reallocate the data and indexes
collection.drop()
print(f"Droping the collection: {TWEETS_COLLECTION_NAME}.")

users = set()

# Read CSV and insert into MongoDB
with open(CSV_FILE_PATH, newline='', encoding='utf-8') as csvfile:
    reader = csv.DictReader(csvfile)
    posts = []
    for row in reader:
        row["userId"] = int(row["userId"])  # Convert userId to integer
        row["creation_datetime"] = datetime.strptime(row["creation_datetime"], "%Y-%m-%d %H:%M:%S")  # Convert to datetime
        posts.append(row)
        users.add(row["userId"])

    if posts:
        collection.insert_many(posts)
        print(f"Inserted {len(posts)} posts into MongoDB.")
    else:
        print("No posts to insert.")

# index the recently created collection
collection.create_index([("userId", pymongo.ASCENDING)], name="userId_index")


# followers generation
collection = db[FOLLOWS_COLLECTION_NAME]

# Drop the actual collection in order to recreate and reallocate the data and indexes
collection.drop()
print(f"Droping the collection: {TWEETS_COLLECTION_NAME}.")

followers = []

for user in users:
    num_following = random.randint(1, 15)
    following = random.sample([uid for uid in users if uid != user], num_following)    
    followers.append({"userId": user, "followedUsers": following})

collection.insert_many(followers)

# index the recently created collection
collection.create_index([("userId", pymongo.ASCENDING)], name="userId_index")

print(f"Inserted {len(followers)} users with follow relationships into MongoDB.")

# Close the MongoDB connection
client.close()

print("Data loading completed.")