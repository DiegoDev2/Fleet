package main

// Nettle - Low-level cryptographic library
// Homepage: https://www.lysator.liu.se/~nisse/nettle/

import (
	"fmt"
	
	"os/exec"
)

func installNettle() {
	// Método 1: Descargar y extraer .tar.gz
	nettle_tar_url := "https://ftp.gnu.org/gnu/nettle/nettle-3.10.tar.gz"
	nettle_cmd_tar := exec.Command("curl", "-L", nettle_tar_url, "-o", "package.tar.gz")
	err := nettle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nettle_zip_url := "https://ftp.gnu.org/gnu/nettle/nettle-3.10.zip"
	nettle_cmd_zip := exec.Command("curl", "-L", nettle_zip_url, "-o", "package.zip")
	err = nettle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nettle_bin_url := "https://ftp.gnu.org/gnu/nettle/nettle-3.10.bin"
	nettle_cmd_bin := exec.Command("curl", "-L", nettle_bin_url, "-o", "binary.bin")
	err = nettle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nettle_src_url := "https://ftp.gnu.org/gnu/nettle/nettle-3.10.src.tar.gz"
	nettle_cmd_src := exec.Command("curl", "-L", nettle_src_url, "-o", "source.tar.gz")
	err = nettle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nettle_cmd_direct := exec.Command("./binary")
	err = nettle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
