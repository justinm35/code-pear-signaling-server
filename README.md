# code-pear-signaling-server
<img width="200" alt="IMG_0007" src="https://github.com/user-attachments/assets/c232c040-c83f-44a6-b96b-4f06de3db6b1" />

### Project Overview
### Features
### Getting Started
### Usage
To run this SDP without it's code-pear-client counterpart you can simply run the go server and curl the endpoints following the provided sequence diagram

To run server, while in the root dir run the following
`go run .`
To step through the expected usage of this signaling server preform the following terminal curl commands in order from another terminal window.
Offerer - Send the sdp to the server:
`curl -X POST -d "sdp=<insert sdp here>" "http://localhost:8080/offer_provide_sdp"`
Acceptor - Fetch offer SDP from the server using access key provided in first call:
`curl "http://localhost:8080/accept_get_sdp?access_key=<insert randomly generated acccess key>"`
Acceptor - Send SDP generated using the offer's sdp to server with access key:
`curl -X POST -d "sdp=<insert sdp here>" "http://localhost:8080/accept_provide_sdp?access_key=<insert access key>"`
Offerer - Fetch accept SDP from the server using access key (normally this endpoint would get polled from the start):
`curl "http://localhost:8080/offer_poll_for_sdp?access_key=<insert randomly generated acccess key>"`

### Architecture
<img width="500" src="https://github.com/user-attachments/assets/5201a0a1-3260-4940-9999-f0d8e792d154" />

### API Endpoints
| Action | Endpoint                | Query Params | Form Params | OK response        |
| ------ | ----------------------- | -----------  | ----------- | ------------------ |
| POST   | /offer_provide_sdp      | N/A          | {sdp: ""}   | <access_key>       |
| GET    | /offer_poll_for_sdp     | access_key=  |             | <acceptors sdp>    |
| GET    | /accept_get_sdp         | access_key=  |             | <offerers sdp>     |
| POST   | /accept_provide_sdp     | access_key=  |  {sdp: ""}  |  OK                |
