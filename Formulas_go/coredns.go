package main

// Coredns - DNS server that chains plugins
// Homepage: https://coredns.io/

import (
	"fmt"
	
	"os/exec"
)

func installCoredns() {
	// Método 1: Descargar y extraer .tar.gz
	coredns_tar_url := "https://github.com/coredns/coredns/archive/refs/tags/v1.11.3.tar.gz"
	coredns_cmd_tar := exec.Command("curl", "-L", coredns_tar_url, "-o", "package.tar.gz")
	err := coredns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	coredns_zip_url := "https://github.com/coredns/coredns/archive/refs/tags/v1.11.3.zip"
	coredns_cmd_zip := exec.Command("curl", "-L", coredns_zip_url, "-o", "package.zip")
	err = coredns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	coredns_bin_url := "https://github.com/coredns/coredns/archive/refs/tags/v1.11.3.bin"
	coredns_cmd_bin := exec.Command("curl", "-L", coredns_bin_url, "-o", "binary.bin")
	err = coredns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	coredns_src_url := "https://github.com/coredns/coredns/archive/refs/tags/v1.11.3.src.tar.gz"
	coredns_cmd_src := exec.Command("curl", "-L", coredns_src_url, "-o", "source.tar.gz")
	err = coredns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	coredns_cmd_direct := exec.Command("./binary")
	err = coredns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: bind")
exec.Command("latte", "install", "bind")
}
