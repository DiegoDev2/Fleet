package main

// Certigo - Utility to examine and validate certificates in a variety of formats
// Homepage: https://github.com/square/certigo

import (
	"fmt"
	
	"os/exec"
)

func installCertigo() {
	// Método 1: Descargar y extraer .tar.gz
	certigo_tar_url := "https://github.com/square/certigo/archive/refs/tags/v1.16.0.tar.gz"
	certigo_cmd_tar := exec.Command("curl", "-L", certigo_tar_url, "-o", "package.tar.gz")
	err := certigo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	certigo_zip_url := "https://github.com/square/certigo/archive/refs/tags/v1.16.0.zip"
	certigo_cmd_zip := exec.Command("curl", "-L", certigo_zip_url, "-o", "package.zip")
	err = certigo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	certigo_bin_url := "https://github.com/square/certigo/archive/refs/tags/v1.16.0.bin"
	certigo_cmd_bin := exec.Command("curl", "-L", certigo_bin_url, "-o", "binary.bin")
	err = certigo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	certigo_src_url := "https://github.com/square/certigo/archive/refs/tags/v1.16.0.src.tar.gz"
	certigo_cmd_src := exec.Command("curl", "-L", certigo_src_url, "-o", "source.tar.gz")
	err = certigo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	certigo_cmd_direct := exec.Command("./binary")
	err = certigo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
