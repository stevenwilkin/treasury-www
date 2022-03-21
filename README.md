# treasury-www

Web UI for [Treasury](https://github.com/stevenwilkin/treasury)

## Building

	go build .

## Running

	./treasury-www

Then browse to [http://0.0.0.0.0:8080](http://0.0.0.0.0:8080)

## Systemd service

Copy the service unit file to the configuration directory:

	cp treasury-www.service /etc/systemd/system

Enable and start the service:

	systemctl enable treasury-www
	systemctl start treasury-www
