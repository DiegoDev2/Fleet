package main

// Signmykey - Automated SSH Certificate Authority
// Homepage: https://signmykey.io

import (
	"fmt"
	
	"os/exec"
)

func installSignmykey() {
	// Método 1: Descargar y extraer .tar.gz
	signmykey_tar_url := "https://github.com/signmykeyio/signmykey/archive/refs/tags/v0.8.7.tar.gz"
	signmykey_cmd_tar := exec.Command("curl", "-L", signmykey_tar_url, "-o", "package.tar.gz")
	err := signmykey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	signmykey_zip_url := "https://github.com/signmykeyio/signmykey/archive/refs/tags/v0.8.7.zip"
	signmykey_cmd_zip := exec.Command("curl", "-L", signmykey_zip_url, "-o", "package.zip")
	err = signmykey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	signmykey_bin_url := "https://github.com/signmykeyio/signmykey/archive/refs/tags/v0.8.7.bin"
	signmykey_cmd_bin := exec.Command("curl", "-L", signmykey_bin_url, "-o", "binary.bin")
	err = signmykey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	signmykey_src_url := "https://github.com/signmykeyio/signmykey/archive/refs/tags/v0.8.7.src.tar.gz"
	signmykey_cmd_src := exec.Command("curl", "-L", signmykey_src_url, "-o", "source.tar.gz")
	err = signmykey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	signmykey_cmd_direct := exec.Command("./binary")
	err = signmykey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
