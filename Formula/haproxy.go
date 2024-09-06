package main

// Haproxy - Reliable, high performance TCP/HTTP load balancer
// Homepage: https://www.haproxy.org/

import (
	"fmt"
	
	"os/exec"
)

func installHaproxy() {
	// Método 1: Descargar y extraer .tar.gz
	haproxy_tar_url := "https://www.haproxy.org/download/3.0/src/haproxy-3.0.4.tar.gz"
	haproxy_cmd_tar := exec.Command("curl", "-L", haproxy_tar_url, "-o", "package.tar.gz")
	err := haproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	haproxy_zip_url := "https://www.haproxy.org/download/3.0/src/haproxy-3.0.4.zip"
	haproxy_cmd_zip := exec.Command("curl", "-L", haproxy_zip_url, "-o", "package.zip")
	err = haproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	haproxy_bin_url := "https://www.haproxy.org/download/3.0/src/haproxy-3.0.4.bin"
	haproxy_cmd_bin := exec.Command("curl", "-L", haproxy_bin_url, "-o", "binary.bin")
	err = haproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	haproxy_src_url := "https://www.haproxy.org/download/3.0/src/haproxy-3.0.4.src.tar.gz"
	haproxy_cmd_src := exec.Command("curl", "-L", haproxy_src_url, "-o", "source.tar.gz")
	err = haproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	haproxy_cmd_direct := exec.Command("./binary")
	err = haproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
