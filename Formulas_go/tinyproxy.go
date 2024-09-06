package main

// Tinyproxy - HTTP/HTTPS proxy for POSIX systems
// Homepage: https://tinyproxy.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installTinyproxy() {
	// Método 1: Descargar y extraer .tar.gz
	tinyproxy_tar_url := "https://github.com/tinyproxy/tinyproxy/releases/download/1.11.2/tinyproxy-1.11.2.tar.xz"
	tinyproxy_cmd_tar := exec.Command("curl", "-L", tinyproxy_tar_url, "-o", "package.tar.gz")
	err := tinyproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tinyproxy_zip_url := "https://github.com/tinyproxy/tinyproxy/releases/download/1.11.2/tinyproxy-1.11.2.tar.xz"
	tinyproxy_cmd_zip := exec.Command("curl", "-L", tinyproxy_zip_url, "-o", "package.zip")
	err = tinyproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tinyproxy_bin_url := "https://github.com/tinyproxy/tinyproxy/releases/download/1.11.2/tinyproxy-1.11.2.tar.xz"
	tinyproxy_cmd_bin := exec.Command("curl", "-L", tinyproxy_bin_url, "-o", "binary.bin")
	err = tinyproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tinyproxy_src_url := "https://github.com/tinyproxy/tinyproxy/releases/download/1.11.2/tinyproxy-1.11.2.tar.xz"
	tinyproxy_cmd_src := exec.Command("curl", "-L", tinyproxy_src_url, "-o", "source.tar.gz")
	err = tinyproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tinyproxy_cmd_direct := exec.Command("./binary")
	err = tinyproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
}
