import wfuzz

url = "https://reqres.in/api/users/FUZZ"
not_founds = []
i = 1
j = 20
session = wfuzz.get_payload(range(i, j))
for result in session.fuzz(url=url):
    print(f"Response {result.description}- status code {result.code}")
    if result.code == 404:
        not_founds.append(result.description)
not_founds.sort()
print(f"404 results {not_founds}")
