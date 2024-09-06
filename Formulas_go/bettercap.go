package main

// Bettercap - Swiss army knife for network attacks and monitoring
// Homepage: https://www.bettercap.org/

import (
	"fmt"
	
	"os/exec"
)

func installBettercap() {
	// Método 1: Descargar y extraer .tar.gz
	bettercap_tar_url := "https://github.com/bettercap/bettercap/archive/refs/tags/v2.33.0.tar.gz"
	bettercap_cmd_tar := exec.Command("curl", "-L", bettercap_tar_url, "-o", "package.tar.gz")
	err := bettercap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bettercap_zip_url := "https://github.com/bettercap/bettercap/archive/refs/tags/v2.33.0.zip"
	bettercap_cmd_zip := exec.Command("curl", "-L", bettercap_zip_url, "-o", "package.zip")
	err = bettercap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bettercap_bin_url := "https://github.com/bettercap/bettercap/archive/refs/tags/v2.33.0.bin"
	bettercap_cmd_bin := exec.Command("curl", "-L", bettercap_bin_url, "-o", "binary.bin")
	err = bettercap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bettercap_src_url := "https://github.com/bettercap/bettercap/archive/refs/tags/v2.33.0.src.tar.gz"
	bettercap_cmd_src := exec.Command("curl", "-L", bettercap_src_url, "-o", "source.tar.gz")
	err = bettercap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bettercap_cmd_direct := exec.Command("./binary")
	err = bettercap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: libnetfilter-queue")
exec.Command("latte", "install", "libnetfilter-queue")
}
