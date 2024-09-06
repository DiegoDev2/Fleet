package main

// Tcptraceroute - Traceroute implementation using TCP packets
// Homepage: https://github.com/mct/tcptraceroute

import (
	"fmt"
	
	"os/exec"
)

func installTcptraceroute() {
	// Método 1: Descargar y extraer .tar.gz
	tcptraceroute_tar_url := "https://github.com/mct/tcptraceroute/archive/refs/tags/tcptraceroute-1.5beta7.tar.gz"
	tcptraceroute_cmd_tar := exec.Command("curl", "-L", tcptraceroute_tar_url, "-o", "package.tar.gz")
	err := tcptraceroute_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcptraceroute_zip_url := "https://github.com/mct/tcptraceroute/archive/refs/tags/tcptraceroute-1.5beta7.zip"
	tcptraceroute_cmd_zip := exec.Command("curl", "-L", tcptraceroute_zip_url, "-o", "package.zip")
	err = tcptraceroute_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcptraceroute_bin_url := "https://github.com/mct/tcptraceroute/archive/refs/tags/tcptraceroute-1.5beta7.bin"
	tcptraceroute_cmd_bin := exec.Command("curl", "-L", tcptraceroute_bin_url, "-o", "binary.bin")
	err = tcptraceroute_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcptraceroute_src_url := "https://github.com/mct/tcptraceroute/archive/refs/tags/tcptraceroute-1.5beta7.src.tar.gz"
	tcptraceroute_cmd_src := exec.Command("curl", "-L", tcptraceroute_src_url, "-o", "source.tar.gz")
	err = tcptraceroute_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcptraceroute_cmd_direct := exec.Command("./binary")
	err = tcptraceroute_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libnet")
exec.Command("latte", "install", "libnet")
}
