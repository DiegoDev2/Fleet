package main

// Bgpq4 - BGP filtering automation for Cisco, Juniper, BIRD and OpenBGPD routers
// Homepage: https://github.com/bgp/bgpq4

import (
	"fmt"
	
	"os/exec"
)

func installBgpq4() {
	// Método 1: Descargar y extraer .tar.gz
	bgpq4_tar_url := "https://github.com/bgp/bgpq4/archive/refs/tags/1.15.tar.gz"
	bgpq4_cmd_tar := exec.Command("curl", "-L", bgpq4_tar_url, "-o", "package.tar.gz")
	err := bgpq4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bgpq4_zip_url := "https://github.com/bgp/bgpq4/archive/refs/tags/1.15.zip"
	bgpq4_cmd_zip := exec.Command("curl", "-L", bgpq4_zip_url, "-o", "package.zip")
	err = bgpq4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bgpq4_bin_url := "https://github.com/bgp/bgpq4/archive/refs/tags/1.15.bin"
	bgpq4_cmd_bin := exec.Command("curl", "-L", bgpq4_bin_url, "-o", "binary.bin")
	err = bgpq4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bgpq4_src_url := "https://github.com/bgp/bgpq4/archive/refs/tags/1.15.src.tar.gz"
	bgpq4_cmd_src := exec.Command("curl", "-L", bgpq4_src_url, "-o", "source.tar.gz")
	err = bgpq4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bgpq4_cmd_direct := exec.Command("./binary")
	err = bgpq4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
