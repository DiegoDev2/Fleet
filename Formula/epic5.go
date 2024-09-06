package main

// Epic5 - Enhanced, programmable IRC client
// Homepage: https://www.epicsol.org/

import (
	"fmt"
	
	"os/exec"
)

func installEpic5() {
	// Método 1: Descargar y extraer .tar.gz
	epic5_tar_url := "https://ftp.epicsol.org/pub/epic/EPIC5-PRODUCTION/epic5-2.6.tar.xz"
	epic5_cmd_tar := exec.Command("curl", "-L", epic5_tar_url, "-o", "package.tar.gz")
	err := epic5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	epic5_zip_url := "https://ftp.epicsol.org/pub/epic/EPIC5-PRODUCTION/epic5-2.6.tar.xz"
	epic5_cmd_zip := exec.Command("curl", "-L", epic5_zip_url, "-o", "package.zip")
	err = epic5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	epic5_bin_url := "https://ftp.epicsol.org/pub/epic/EPIC5-PRODUCTION/epic5-2.6.tar.xz"
	epic5_cmd_bin := exec.Command("curl", "-L", epic5_bin_url, "-o", "binary.bin")
	err = epic5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	epic5_src_url := "https://ftp.epicsol.org/pub/epic/EPIC5-PRODUCTION/epic5-2.6.tar.xz"
	epic5_cmd_src := exec.Command("curl", "-L", epic5_src_url, "-o", "source.tar.gz")
	err = epic5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	epic5_cmd_direct := exec.Command("./binary")
	err = epic5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
