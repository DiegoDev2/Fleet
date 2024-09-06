package main

// Passenger - Server for Ruby, Python, and Node.js apps via Apache/NGINX
// Homepage: https://www.phusionpassenger.com/

import (
	"fmt"
	
	"os/exec"
)

func installPassenger() {
	// Método 1: Descargar y extraer .tar.gz
	passenger_tar_url := "https://github.com/phusion/passenger/releases/download/release-6.0.23/passenger-6.0.23.tar.gz"
	passenger_cmd_tar := exec.Command("curl", "-L", passenger_tar_url, "-o", "package.tar.gz")
	err := passenger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	passenger_zip_url := "https://github.com/phusion/passenger/releases/download/release-6.0.23/passenger-6.0.23.zip"
	passenger_cmd_zip := exec.Command("curl", "-L", passenger_zip_url, "-o", "package.zip")
	err = passenger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	passenger_bin_url := "https://github.com/phusion/passenger/releases/download/release-6.0.23/passenger-6.0.23.bin"
	passenger_cmd_bin := exec.Command("curl", "-L", passenger_bin_url, "-o", "binary.bin")
	err = passenger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	passenger_src_url := "https://github.com/phusion/passenger/releases/download/release-6.0.23/passenger-6.0.23.src.tar.gz"
	passenger_cmd_src := exec.Command("curl", "-L", passenger_src_url, "-o", "source.tar.gz")
	err = passenger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	passenger_cmd_direct := exec.Command("./binary")
	err = passenger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: httpd")
exec.Command("latte", "install", "httpd")
	fmt.Println("Instalando dependencia: nginx")
exec.Command("latte", "install", "nginx")
	fmt.Println("Instalando dependencia: apr")
exec.Command("latte", "install", "apr")
	fmt.Println("Instalando dependencia: apr-util")
exec.Command("latte", "install", "apr-util")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
}
