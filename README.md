Steps for a correct deployment:

1. Install Docker Platform in your server (https://www.docker.com/)
2. Install Docker Compose in the same server as point 1 (https://docs.docker.com/compose/)
3. Clone the repository to any location you like
4. Enter the main folder of the cloned repository
5. Rename ".env-example" file to ".env"
6. Edit the file ".env" with the correct values for your HEI (Please note the after the "=" your string should not contain any spaces!!)
7. In a terminal window, execute the following command to build the container:
    a) docker-compose build ewp-rest
8. In a terminal window, execute the following command to start the container:
  a) docker-compose up ewp-rest (if you put -d option it will detach from the terminal)
9. Should be up and running. Try to access the https://domain_you_chose:port_you_chose/rest/manifest on your browser
