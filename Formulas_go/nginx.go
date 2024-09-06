package main

// Nginx - HTTP(S) server and reverse proxy, and IMAP/POP3 proxy server
// Homepage: https://nginx.org/

import (
	"fmt"
	
	"os/exec"
)

func installNginx() {
	// Método 1: Descargar y extraer .tar.gz
	nginx_tar_url := "https://nginx.org/download/nginx-1.27.1.tar.gz"
	nginx_cmd_tar := exec.Command("curl", "-L", nginx_tar_url, "-o", "package.tar.gz")
	err := nginx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nginx_zip_url := "https://nginx.org/download/nginx-1.27.1.zip"
	nginx_cmd_zip := exec.Command("curl", "-L", nginx_zip_url, "-o", "package.zip")
	err = nginx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nginx_bin_url := "https://nginx.org/download/nginx-1.27.1.bin"
	nginx_cmd_bin := exec.Command("curl", "-L", nginx_bin_url, "-o", "binary.bin")
	err = nginx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nginx_src_url := "https://nginx.org/download/nginx-1.27.1.src.tar.gz"
	nginx_cmd_src := exec.Command("curl", "-L", nginx_src_url, "-o", "source.tar.gz")
	err = nginx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nginx_cmd_direct := exec.Command("./binary")
	err = nginx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
}
