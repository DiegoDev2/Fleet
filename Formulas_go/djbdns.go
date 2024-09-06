package main

// Djbdns - D.J. Bernstein's DNS tools
// Homepage: https://cr.yp.to/djbdns.html

import (
	"fmt"
	
	"os/exec"
)

func installDjbdns() {
	// Método 1: Descargar y extraer .tar.gz
	djbdns_tar_url := "https://cr.yp.to/djbdns/djbdns-1.05.tar.gz"
	djbdns_cmd_tar := exec.Command("curl", "-L", djbdns_tar_url, "-o", "package.tar.gz")
	err := djbdns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	djbdns_zip_url := "https://cr.yp.to/djbdns/djbdns-1.05.zip"
	djbdns_cmd_zip := exec.Command("curl", "-L", djbdns_zip_url, "-o", "package.zip")
	err = djbdns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	djbdns_bin_url := "https://cr.yp.to/djbdns/djbdns-1.05.bin"
	djbdns_cmd_bin := exec.Command("curl", "-L", djbdns_bin_url, "-o", "binary.bin")
	err = djbdns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	djbdns_src_url := "https://cr.yp.to/djbdns/djbdns-1.05.src.tar.gz"
	djbdns_cmd_src := exec.Command("curl", "-L", djbdns_src_url, "-o", "source.tar.gz")
	err = djbdns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	djbdns_cmd_direct := exec.Command("./binary")
	err = djbdns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: daemontools")
exec.Command("latte", "install", "daemontools")
	fmt.Println("Instalando dependencia: ucspi-tcp")
exec.Command("latte", "install", "ucspi-tcp")
	fmt.Println("Instalando dependencia: fakeroot")
exec.Command("latte", "install", "fakeroot")
}
