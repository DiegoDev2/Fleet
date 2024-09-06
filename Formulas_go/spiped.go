package main

// Spiped - Secure pipe daemon
// Homepage: https://www.tarsnap.com/spiped.html

import (
	"fmt"
	
	"os/exec"
)

func installSpiped() {
	// Método 1: Descargar y extraer .tar.gz
	spiped_tar_url := "https://www.tarsnap.com/spiped/spiped-1.6.2.tgz"
	spiped_cmd_tar := exec.Command("curl", "-L", spiped_tar_url, "-o", "package.tar.gz")
	err := spiped_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spiped_zip_url := "https://www.tarsnap.com/spiped/spiped-1.6.2.tgz"
	spiped_cmd_zip := exec.Command("curl", "-L", spiped_zip_url, "-o", "package.zip")
	err = spiped_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spiped_bin_url := "https://www.tarsnap.com/spiped/spiped-1.6.2.tgz"
	spiped_cmd_bin := exec.Command("curl", "-L", spiped_bin_url, "-o", "binary.bin")
	err = spiped_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spiped_src_url := "https://www.tarsnap.com/spiped/spiped-1.6.2.tgz"
	spiped_cmd_src := exec.Command("curl", "-L", spiped_src_url, "-o", "source.tar.gz")
	err = spiped_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spiped_cmd_direct := exec.Command("./binary")
	err = spiped_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: bsdmake")
exec.Command("latte", "install", "bsdmake")
}
