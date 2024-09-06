package main

// AircrackNg - Next-generation aircrack with lots of new features
// Homepage: https://aircrack-ng.org/

import (
	"fmt"
	
	"os/exec"
)

func installAircrackNg() {
	// Método 1: Descargar y extraer .tar.gz
	aircrackng_tar_url := "https://download.aircrack-ng.org/aircrack-ng-1.7.tar.gz"
	aircrackng_cmd_tar := exec.Command("curl", "-L", aircrackng_tar_url, "-o", "package.tar.gz")
	err := aircrackng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aircrackng_zip_url := "https://download.aircrack-ng.org/aircrack-ng-1.7.zip"
	aircrackng_cmd_zip := exec.Command("curl", "-L", aircrackng_zip_url, "-o", "package.zip")
	err = aircrackng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aircrackng_bin_url := "https://download.aircrack-ng.org/aircrack-ng-1.7.bin"
	aircrackng_cmd_bin := exec.Command("curl", "-L", aircrackng_bin_url, "-o", "binary.bin")
	err = aircrackng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aircrackng_src_url := "https://download.aircrack-ng.org/aircrack-ng-1.7.src.tar.gz"
	aircrackng_cmd_src := exec.Command("curl", "-L", aircrackng_src_url, "-o", "source.tar.gz")
	err = aircrackng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aircrackng_cmd_direct := exec.Command("./binary")
	err = aircrackng_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
