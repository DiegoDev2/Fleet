package main

// Web100clt - Command-line version of NDT diagnostic client
// Homepage: https://software.internet2.edu/ndt/

import (
	"fmt"
	
	"os/exec"
)

func installWeb100clt() {
	// Método 1: Descargar y extraer .tar.gz
	web100clt_tar_url := "https://software.internet2.edu/sources/ndt/ndt-3.7.0.2.tar.gz"
	web100clt_cmd_tar := exec.Command("curl", "-L", web100clt_tar_url, "-o", "package.tar.gz")
	err := web100clt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	web100clt_zip_url := "https://software.internet2.edu/sources/ndt/ndt-3.7.0.2.zip"
	web100clt_cmd_zip := exec.Command("curl", "-L", web100clt_zip_url, "-o", "package.zip")
	err = web100clt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	web100clt_bin_url := "https://software.internet2.edu/sources/ndt/ndt-3.7.0.2.bin"
	web100clt_cmd_bin := exec.Command("curl", "-L", web100clt_bin_url, "-o", "binary.bin")
	err = web100clt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	web100clt_src_url := "https://software.internet2.edu/sources/ndt/ndt-3.7.0.2.src.tar.gz"
	web100clt_cmd_src := exec.Command("curl", "-L", web100clt_src_url, "-o", "source.tar.gz")
	err = web100clt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	web100clt_cmd_direct := exec.Command("./binary")
	err = web100clt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: i2util")
	exec.Command("latte", "install", "i2util").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
