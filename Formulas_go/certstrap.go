package main

// Certstrap - Tools to bootstrap CAs, certificate requests, and signed certificates
// Homepage: https://github.com/square/certstrap

import (
	"fmt"
	
	"os/exec"
)

func installCertstrap() {
	// Método 1: Descargar y extraer .tar.gz
	certstrap_tar_url := "https://github.com/square/certstrap/archive/refs/tags/v1.3.0.tar.gz"
	certstrap_cmd_tar := exec.Command("curl", "-L", certstrap_tar_url, "-o", "package.tar.gz")
	err := certstrap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	certstrap_zip_url := "https://github.com/square/certstrap/archive/refs/tags/v1.3.0.zip"
	certstrap_cmd_zip := exec.Command("curl", "-L", certstrap_zip_url, "-o", "package.zip")
	err = certstrap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	certstrap_bin_url := "https://github.com/square/certstrap/archive/refs/tags/v1.3.0.bin"
	certstrap_cmd_bin := exec.Command("curl", "-L", certstrap_bin_url, "-o", "binary.bin")
	err = certstrap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	certstrap_src_url := "https://github.com/square/certstrap/archive/refs/tags/v1.3.0.src.tar.gz"
	certstrap_cmd_src := exec.Command("curl", "-L", certstrap_src_url, "-o", "source.tar.gz")
	err = certstrap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	certstrap_cmd_direct := exec.Command("./binary")
	err = certstrap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
