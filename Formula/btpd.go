package main

// Btpd - BitTorrent Protocol Daemon
// Homepage: https://github.com/btpd/btpd

import (
	"fmt"
	
	"os/exec"
)

func installBtpd() {
	// Método 1: Descargar y extraer .tar.gz
	btpd_tar_url := "https://github.com/btpd/btpd/archive/refs/tags/v0.16.tar.gz"
	btpd_cmd_tar := exec.Command("curl", "-L", btpd_tar_url, "-o", "package.tar.gz")
	err := btpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	btpd_zip_url := "https://github.com/btpd/btpd/archive/refs/tags/v0.16.zip"
	btpd_cmd_zip := exec.Command("curl", "-L", btpd_zip_url, "-o", "package.zip")
	err = btpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	btpd_bin_url := "https://github.com/btpd/btpd/archive/refs/tags/v0.16.bin"
	btpd_cmd_bin := exec.Command("curl", "-L", btpd_bin_url, "-o", "binary.bin")
	err = btpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	btpd_src_url := "https://github.com/btpd/btpd/archive/refs/tags/v0.16.src.tar.gz"
	btpd_cmd_src := exec.Command("curl", "-L", btpd_src_url, "-o", "source.tar.gz")
	err = btpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	btpd_cmd_direct := exec.Command("./binary")
	err = btpd_cmd_direct.Run()
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
