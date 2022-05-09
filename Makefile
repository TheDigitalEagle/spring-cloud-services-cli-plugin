build:
	rm spring-cloud-services-cli-plugin
	go build -ldflags="-X main.pluginVersion=1.0.27"

install:
	cf uninstall-plugin spring-cloud-services
	cf install-plugin -f spring-cloud-services-cli-plugin