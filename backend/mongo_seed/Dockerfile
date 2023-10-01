FROM mongo:4
 # Will be set through Environment Files
 ARG MONGOURI

 ENV MONGOURI ${MONGOURI}

 COPY todos.json /todos.json

 CMD mongoimport --uri ${MONGOURI} --collection todos --drop --jsonArray --file /todos.json


