package main

// UBootTools - Universal boot loader
// Homepage: https://www.denx.de/wiki/U-Boot/

import (
	"fmt"
	
	"os/exec"
)

func installUBootTools() {
	// Método 1: Descargar y extraer .tar.gz
	uboottools_tar_url := "https://ftp.denx.de/pub/u-boot/u-boot-2024.07.tar.bz2"
	uboottools_cmd_tar := exec.Command("curl", "-L", uboottools_tar_url, "-o", "package.tar.gz")
	err := uboottools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uboottools_zip_url := "https://ftp.denx.de/pub/u-boot/u-boot-2024.07.tar.bz2"
	uboottools_cmd_zip := exec.Command("curl", "-L", uboottools_zip_url, "-o", "package.zip")
	err = uboottools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uboottools_bin_url := "https://ftp.denx.de/pub/u-boot/u-boot-2024.07.tar.bz2"
	uboottools_cmd_bin := exec.Command("curl", "-L", uboottools_bin_url, "-o", "binary.bin")
	err = uboottools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uboottools_src_url := "https://ftp.denx.de/pub/u-boot/u-boot-2024.07.tar.bz2"
	uboottools_cmd_src := exec.Command("curl", "-L", uboottools_src_url, "-o", "source.tar.gz")
	err = uboottools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uboottools_cmd_direct := exec.Command("./binary")
	err = uboottools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
