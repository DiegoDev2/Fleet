package main

// Mkcert - Simple tool to make locally trusted development certificates
// Homepage: https://github.com/FiloSottile/mkcert

import (
	"fmt"
	
	"os/exec"
)

func installMkcert() {
	// Método 1: Descargar y extraer .tar.gz
	mkcert_tar_url := "https://github.com/FiloSottile/mkcert/archive/refs/tags/v1.4.4.tar.gz"
	mkcert_cmd_tar := exec.Command("curl", "-L", mkcert_tar_url, "-o", "package.tar.gz")
	err := mkcert_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mkcert_zip_url := "https://github.com/FiloSottile/mkcert/archive/refs/tags/v1.4.4.zip"
	mkcert_cmd_zip := exec.Command("curl", "-L", mkcert_zip_url, "-o", "package.zip")
	err = mkcert_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mkcert_bin_url := "https://github.com/FiloSottile/mkcert/archive/refs/tags/v1.4.4.bin"
	mkcert_cmd_bin := exec.Command("curl", "-L", mkcert_bin_url, "-o", "binary.bin")
	err = mkcert_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mkcert_src_url := "https://github.com/FiloSottile/mkcert/archive/refs/tags/v1.4.4.src.tar.gz"
	mkcert_cmd_src := exec.Command("curl", "-L", mkcert_src_url, "-o", "source.tar.gz")
	err = mkcert_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mkcert_cmd_direct := exec.Command("./binary")
	err = mkcert_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
