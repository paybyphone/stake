# Stake - A Claims Management API

From https://en.wikipedia.org/wiki/Claims-based_identity:

> Claims are not what the subject can and cannot do. They are what the subject is or is not. It is up to the application receiving the incoming claim to map the is/is not claims to the may/may not rules of the application.

__For example:__

| Claimant | Subject | Claim |
|:
| jmiller | email | jmiller@somewhere.com |
| jmiller | BackofficeAdmin | true |
| gsmith | BackofficeAdmin | true |

__Every time a claim is made, the claimant is also added to a list of all claims made about a subject:__

| Claimant | Subject | Claim |
|:
| email | jmiller | jmiller@somewhere.com |
| BackofficeAdmin | jmiller | true |
| BackofficeAdmin | gsmith | true |


"Claimant jmiller claims to have jmiller@somewhere.com for subject email."

"Claimant email claims to have jmiller for subject jmiller@somewhere.com."

"Claimant BackofficeAdmin claims to have true for subject jmiller."

## Usage

Clone this repo and start the API:

```
cd stake
go get
cd claimsapi
go build
./claimsapi
```

POST a claim:

```
curl -X POST -H "Content-Type: application/json" -H "Cache-Control: no-cache" -d '{ "subject" : "email", "claim" : "alfred@someemail.com" }' 'http://localhost:8090/claimants/alfred/claims'
```

GET a claim:

```
curl -X GET -H "Content-Type: application/json" -H "Cache-Control: no-cache"  'http://localhost:8090/claimants/alfred/claims'
```

GET everybody who made a claim about (e.g.) 'email':

```
curl -X GET -H "Content-Type: application/json" -H "Cache-Control: no-cache"  'http://localhost:8090/claimants/email/claims'
```

DELETE a claim:

```
curl -X DELETE -H "Content-Type: application/json" -H "Cache-Control: no-cache" -d '{ "subject" : "email", "claim" : "alfred@someemail.com" }' 'http://localhost:8090/claimants/alfred/claims'
```
