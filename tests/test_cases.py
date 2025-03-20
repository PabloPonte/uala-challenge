import random
import requests

print("Runing all the end-to-end tests")

API_URL = "http://localhost:5000"

writer_user_1 = random.randint(30000, 39999)
writer_user_2 = random.randint(30000, 39999)

reader_user_1 = random.randint(30000, 39999)
reader_user_2 = random.randint(30000, 39999)
reader_user_3 = random.randint(30000, 39999)

message_1 = "Thinking about upgrading my car, any suggestions?"
message_2 = "Just finished my first marathon, I'm so happy!"
message_3 = "I'm so tired of this weather, I need a vacation"
message_4 = "I'm so happy to be part of this community, you guys are amazing!"
message_too_long = "This is a very long message, I'm just testing the limit of 280 characters, I hope this is enough to reach the limit, if not, I will keep writing until I reach it, I'm almost there, just a few more words, and that's it! Still not reached the upper limit...I mean, sky is the limit but I only need 280 letters to hit the limit in this case"

#### Behaviour Table ####
# writer_user_1: will tweet message_1 and message_too_long
# writer_user_2: will tweet message_2, message_3, message_4
# reader_user_1: will follow writer_user_1
# reader_user_2: will follow writer_user_1 and writer_user_2
# reader_user_3: will not follow any writer

testsExecuted = 0
testsFailed = 0


# Tweet creation tests
print("\nRunning tweet creation tests\n")

# Test case 1: New Tweet with less than 280 characters
# Expected Result: The tweet is created successfully
print("Running test case 1: New Tweet with less than 280 characters")

request_body = {    
    "userId": writer_user_1,
    "content": message_1
}

response = requests.post(f"{API_URL}/tweet", json=request_body)

testPass = True

if response.status_code == 201:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False
    
_id = response.json().get("data",{}).get("_id")
userId = response.json().get("data",{}).get("userId")
content = response.json().get("data",{}).get("content")
creationDate = response.json().get("data",{}).get("creationDate")

if _id is None or userId is None or content is None or creationDate is None:
    print("Response Body not OK: ", response.json())
    testPass = False
else:   
    print("Response Body OK")

if testPass:
    print("Test case 1 passed successfully")
else:
    print("Test case 1 failed")
    testsFailed += 1

testsExecuted += 1
    
print("\n#######################################################\n")

# Test case 2: New Tweet with more than 280 characters
# Expected Result: The tweet is not created and an error message is returned
print("Running test case 2: New Tweet with more than 280 characters")

request_body = {    
    "userId": writer_user_1,
    "content": message_too_long
}

response = requests.post(f"{API_URL}/tweet", json=request_body)

testPass = True

if response.status_code == 400:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False
    
_id = response.json().get("data",{}).get("_id")
userId = response.json().get("data",{}).get("userId")
content = response.json().get("data",{}).get("content")
creationDate = response.json().get("data",{}).get("creationDate")
error = response.json().get("error")

if _id is not None or userId is not None or content is not None or creationDate is not None or error is None:
    print("Response Body not OK: ", response.json())
    testPass = False
else:   
    print("Response Body OK")

if testPass:
    print("Test case 2 passed successfully")
else:
    print("Test case 2 failed")
    testsFailed += 1

testsExecuted += 1

print("\n#######################################################\n")

# Test case 3: New Tweet with empty content
# Expected Result: The tweet is not created and an error message is returned
print("Running test case 3: New Tweet with empty content")

request_body = {    
    "userId": writer_user_1
}

response = requests.post(f"{API_URL}/tweet", json=request_body)

testPass = True

if response.status_code == 400:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False
    
_id = response.json().get("data",{}).get("_id")
userId = response.json().get("data",{}).get("userId")
content = response.json().get("data",{}).get("content")
creationDate = response.json().get("data",{}).get("creationDate")
error = response.json().get("error")

if _id is not None or userId is not None or content is not None or creationDate is not None or error is None:
    print("Response Body not OK: ", response.json())
    testPass = False
else:   
    print("Response Body OK")

if testPass:
    print("Test case 3 passed successfully")
else:
    print("Test case 3 failed")
    testsFailed += 1

testsExecuted += 1

print("\n#######################################################\n")

# Test case 4: New Tweet with empty userId
# Expected Result: The tweet is not created and an error message is returned
print("Running test case 4: New Tweet with empty userId")

request_body = {
    "content": message_1
}

response = requests.post(f"{API_URL}/tweet", json=request_body)

testPass = True

if response.status_code == 400:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False
    
_id = response.json().get("data",{}).get("_id")
userId = response.json().get("data",{}).get("userId")
content = response.json().get("data",{}).get("content")
creationDate = response.json().get("data",{}).get("creationDate")
error = response.json().get("error")

if _id is not None or userId is not None or content is not None or creationDate is not None or error is None:
    print("Response Body not OK: ", response.json())
    testPass = False
else:   
    print("Response Body OK")

if testPass:
    print("Test case 4 passed successfully")
else:
    print("Test case 4 failed")
    testsFailed += 1

testsExecuted += 1
print("\n#######################################################\n")


# Follow creation tests
print("Running follow creation tests\n")

# Test case 5: New Follow with userId and followedUser
# Expected Result: The followed user is added to the list of followed users
print("Running test case 5: New Follow with userId and followedUser")

request_body = {
    "userId" : reader_user_1,
	"followedUser" : writer_user_1
}

response = requests.post(f"{API_URL}/follow", json=request_body)

testPass = True

if response.status_code == 201:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False
    
if response.json().get("message") != "Follow created successfully":
    print("Response Body not OK: ", response.json())
    testPass = False
else:   
    print("Response Body OK")

if testPass:
    print("Test case 5 passed successfully")
else:
    print("Test case 5 failed")
    testsFailed += 1

testsExecuted += 1

print("\n#######################################################\n")


# Test case 6: New Follow with empty userId
# Expected Result: The follow is not created and an error message is returned

print("Running test case 6: New Follow with empty userId")

request_body = {
	"followedUser" : writer_user_1
}

response = requests.post(f"{API_URL}/follow", json=request_body)

testPass = True

if response.status_code == 400:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False
    
if response.json().get("error") is None:
    print("Response Body not OK: ", response.json())
    testPass = False
else:   
    print("Response Body OK")

if testPass:
    print("Test case 6 passed successfully")
else:
    print("Test case 6 failed")
    testsFailed += 1

testsExecuted += 1

print("\n#######################################################\n")

# Test case 7: New Follow with empty followedUser
# Expected Result: The follow is not created and an error message is returned
print("Running test case 7: New Follow with empty followedUser")

request_body = {
    "userId" : reader_user_1
}

response = requests.post(f"{API_URL}/follow", json=request_body)

testPass = True

if response.status_code == 400:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False
    
if response.json().get("error") is None:
    print("Response Body not OK: ", response.json())
    testPass = False
else:   
    print("Response Body OK")

if testPass:
    print("Test case 7 passed successfully")
else:
    print("Test case 7 failed")
    testsFailed += 1

testsExecuted += 1

print("\n#######################################################\n")

# Test case 8: New Follow with userId and followedUser already followed
# Expected Result: The followed user is not added to the list of followed users, no error is returned
print("Running test case 8: New Follow with userId and followedUser already followed")

request_body = {
    "userId" : reader_user_1,
	"followedUser" : writer_user_1
}

response = requests.post(f"{API_URL}/follow", json=request_body)

testPass = True

if response.status_code == 201:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False
    
if response.json().get("message") != "Follow created successfully":
    print("Response Body not OK: ", response.json())
    testPass = False
else:   
    print("Response Body OK")

if testPass:
    print("Test case 8 passed successfully")
else:
    print("Test case 8 failed")
    testsFailed += 1

testsExecuted += 1

print("\n#######################################################\n")


# Test case 9: New Follow with an user following himself
# Expected Result: The follow is not created and an error message is returned
print("Running test case 9: New Follow with an user following himself")

request_body = {
    "userId" : reader_user_1,
    "followedUser" : reader_user_1
}

response = requests.post(f"{API_URL}/follow", json=request_body)

testPass = True

if response.status_code == 400:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False
    
if response.json().get("error") is None:
    print("Response Body not OK: ", response.json())
    testPass = False
else:   
    print("Response Body OK")

if testPass:
    print("Test case 9 passed successfully")
else:
    print("Test case 9 failed")
    testsFailed += 1

testsExecuted += 1

print("\n#######################################################\n")


print("Creating the tweets and relations for the timeline tests")

# writer_user_1: will tweet message_1 and message_too_long, done in test case 1 and 2

# writer_user_2: will tweet message_2, message_3, message_4
response = requests.post(f"{API_URL}/tweet", json={ "userId": writer_user_2, "content": message_2 })
response = requests.post(f"{API_URL}/tweet", json={ "userId": writer_user_2, "content": message_3 })
response = requests.post(f"{API_URL}/tweet", json={ "userId": writer_user_2, "content": message_4 })

# reader_user_1: will follow writer_user_1, done in test case 5

# reader_user_2: will follow writer_user_1 and writer_user_2

response = requests.post(f"{API_URL}/follow", json={ "userId" : reader_user_2, "followedUser" : writer_user_1 })
response = requests.post(f"{API_URL}/follow", json={ "userId" : reader_user_2, "followedUser" : writer_user_2 })

# reader_user_3: will not follow any writer, nothing to do here


# Timeline tests
print("Running timeline tests\n")

# Test case 10: Get Timeline with userId with followed users
# Expected Result: The timeline is returned with the tweets of the followed users ordered by creation datetime
print("Running test case 10: Get Timeline with userId with followed users")

response = requests.get(f"{API_URL}/tweet/{reader_user_2}")

testPass = True

if response.status_code == 200:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False

if len(list(response.json().get("payload"))) != 4:
    print("Response Size not ok: ", len(list(response.json().get("payload"))))
    testPass = False
else:
    print("Response Size OK")

if testPass:
    print("Test case 10 passed successfully")
else:
    print("Test case 10 failed")
    testsFailed += 1

testsExecuted += 1

print("\n#######################################################\n")

# Test case 11: Get Timeline with userId with no followed users
# Expected Result: An error message is returned
print("Running test case 11: Get Timeline with userId with no followed users")

response = requests.get(f"{API_URL}/tweet/{reader_user_3}")

testPass = True

if response.status_code == 404:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False

if response.json().get("error") is None:
    print("Response Body not OK: ", response.json())
    testPass = False
else:   
    print("Response Body OK")

if testPass:
    print("Test case 11 passed successfully")
else:
    print("Test case 11 failed")
    testsFailed += 1

testsExecuted += 1

print("\n#######################################################\n")

# Test case 12: Get Timeline with empty userId
# This endpoint receives the userId as a param, so an empty value is not allowed
# Expected Result: An error message is returned (404 page not found)
print("Running test case 11: Get Timeline with userId with no followed users")

response = requests.get(f"{API_URL}/tweet/")

testPass = True

if response.status_code == 404:
    print("Status Code OK")
else: 
    print("Status Code not OK: ", response.status_code)
    testPass = False

if testPass:
    print("Test case 12 passed successfully")
else:
    print("Test case 12 failed")
    testsFailed += 1

testsExecuted += 1

print("\n#######################################################\n")

print(f"Tests executed: {testsExecuted}")
print(f"Tests failed: {testsFailed}")
print(f"Tests passed: {testsExecuted - testsFailed}")
print(f"Tests success rate: {((testsExecuted - testsFailed) / testsExecuted) * 100}%")