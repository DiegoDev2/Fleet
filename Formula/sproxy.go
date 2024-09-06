package main

// Sproxy - HTTP proxy server collecting URLs in a 'siege-friendly' manner
// Homepage: https://www.joedog.org/sproxy-home/

import (
	"fmt"
	
	"os/exec"
)

func installSproxy() {
	// Método 1: Descargar y extraer .tar.gz
	sproxy_tar_url := "https://download.joedog.org/sproxy/sproxy-1.02.tar.gz"
	sproxy_cmd_tar := exec.Command("curl", "-L", sproxy_tar_url, "-o", "package.tar.gz")
	err := sproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sproxy_zip_url := "https://download.joedog.org/sproxy/sproxy-1.02.zip"
	sproxy_cmd_zip := exec.Command("curl", "-L", sproxy_zip_url, "-o", "package.zip")
	err = sproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sproxy_bin_url := "https://download.joedog.org/sproxy/sproxy-1.02.bin"
	sproxy_cmd_bin := exec.Command("curl", "-L", sproxy_bin_url, "-o", "binary.bin")
	err = sproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sproxy_src_url := "https://download.joedog.org/sproxy/sproxy-1.02.src.tar.gz"
	sproxy_cmd_src := exec.Command("curl", "-L", sproxy_src_url, "-o", "source.tar.gz")
	err = sproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sproxy_cmd_direct := exec.Command("./binary")
	err = sproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
