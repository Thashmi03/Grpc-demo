# Grpc-demo

1. Data formats are defined in the form of messages.
2. Every member in the message will be having an index associated with it.
3. Messages can take a variety of datatypes -stings, integers,booleans,arrays and other message types
arrays are represented with keyword called repeated
4. keywords --> service,message,repeated,syntax,package
5. Protobuff files will never implement any logic they will only talk about message exchanging formats and supported methods.

Steps
1. We will create a server 
2. Implement to the rpc method 
3. Deploy it on a port 
4. We will create a client and we will call the server method from the client.