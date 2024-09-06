package main

// Transcrypt - Configure transparent encryption of files in a Git repo
// Homepage: https://github.com/elasticdog/transcrypt

import (
	"fmt"
	
	"os/exec"
)

func installTranscrypt() {
	// Método 1: Descargar y extraer .tar.gz
	transcrypt_tar_url := "https://github.com/elasticdog/transcrypt/archive/refs/tags/v2.2.3.tar.gz"
	transcrypt_cmd_tar := exec.Command("curl", "-L", transcrypt_tar_url, "-o", "package.tar.gz")
	err := transcrypt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	transcrypt_zip_url := "https://github.com/elasticdog/transcrypt/archive/refs/tags/v2.2.3.zip"
	transcrypt_cmd_zip := exec.Command("curl", "-L", transcrypt_zip_url, "-o", "package.zip")
	err = transcrypt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	transcrypt_bin_url := "https://github.com/elasticdog/transcrypt/archive/refs/tags/v2.2.3.bin"
	transcrypt_cmd_bin := exec.Command("curl", "-L", transcrypt_bin_url, "-o", "binary.bin")
	err = transcrypt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	transcrypt_src_url := "https://github.com/elasticdog/transcrypt/archive/refs/tags/v2.2.3.src.tar.gz"
	transcrypt_cmd_src := exec.Command("curl", "-L", transcrypt_src_url, "-o", "source.tar.gz")
	err = transcrypt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	transcrypt_cmd_direct := exec.Command("./binary")
	err = transcrypt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
	fmt.Println("Instalando dependencia: vim")
	exec.Command("latte", "install", "vim").Run()
}
