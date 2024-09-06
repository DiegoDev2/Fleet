package main

// Udptunnel - Tunnel UDP packets over a TCP connection
// Homepage: http://www1.cs.columbia.edu/~lennox/udptunnel/

import (
	"fmt"
	
	"os/exec"
)

func installUdptunnel() {
	// Método 1: Descargar y extraer .tar.gz
	udptunnel_tar_url := "https://pkg.freebsd.org/ports-distfiles/udptunnel-1.1.tar.gz"
	udptunnel_cmd_tar := exec.Command("curl", "-L", udptunnel_tar_url, "-o", "package.tar.gz")
	err := udptunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	udptunnel_zip_url := "https://pkg.freebsd.org/ports-distfiles/udptunnel-1.1.zip"
	udptunnel_cmd_zip := exec.Command("curl", "-L", udptunnel_zip_url, "-o", "package.zip")
	err = udptunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	udptunnel_bin_url := "https://pkg.freebsd.org/ports-distfiles/udptunnel-1.1.bin"
	udptunnel_cmd_bin := exec.Command("curl", "-L", udptunnel_bin_url, "-o", "binary.bin")
	err = udptunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	udptunnel_src_url := "https://pkg.freebsd.org/ports-distfiles/udptunnel-1.1.src.tar.gz"
	udptunnel_cmd_src := exec.Command("curl", "-L", udptunnel_src_url, "-o", "source.tar.gz")
	err = udptunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	udptunnel_cmd_direct := exec.Command("./binary")
	err = udptunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libnsl")
exec.Command("latte", "install", "libnsl")
}
