import mmh3
import requests
import codecs
 
response = requests.get('http://203.245.0.121/favicon.ico')
favicon = codecs.encode(response.content,"base64")

print(f"Content size: {len(response.content)} bytes")
print(f"Base64 size: {len(favicon)} chars")  
print(f"Base64 first 100 chars: {favicon[:100]}")
if len(favicon) > 100:
    print(f"Base64 last 50 chars: {favicon[-50:]}")

hash = mmh3.hash(favicon)
print(f"Hash (decimal): {hash}")