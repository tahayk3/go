# Api simple RESTful con GO



# Synchronization

Esta cuenta con cinco funciones:
- homePage: funcion para manejar todas las solicitudes, endpoint de la pagina principal
- handleRequest: Manejo de solicitudes
- Funcion principal: punto de entreda del programa(servidor), donde se llama a la funcion que escucha solicitudes
- Las dos funciones restantes, se utilizan para simular una solicitud get y post
	> Se utiliza mux

# De tener problemas con "github.com/gorilla/mux"
Es necesario crear un entorno con go por lo general, este se crea el C:\Users\nombre-de-usuario\go\src\github.com\

Se deben utilizar los siguientes comandos:

mkdir -p $env:GOPATH\src\github.com\nombreusuario\nombre-del-proyecto

go mod init github.com/tu-usuario/nombre-del-proyecto

go get -u github.com/gorilla/mux

De esta manera, se crean los archivos go.mod y go.sum

