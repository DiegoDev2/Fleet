package main

// SagittariusScheme - Free Scheme implementation supporting R6RS and R7RS
// Homepage: https://bitbucket.org/ktakashi/sagittarius-scheme/wiki/Home

import (
	"fmt"
	
	"os/exec"
)

func installSagittariusScheme() {
	// Método 1: Descargar y extraer .tar.gz
	sagittariusscheme_tar_url := "https://bitbucket.org/ktakashi/sagittarius-scheme/downloads/sagittarius-0.9.11.tar.gz"
	sagittariusscheme_cmd_tar := exec.Command("curl", "-L", sagittariusscheme_tar_url, "-o", "package.tar.gz")
	err := sagittariusscheme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sagittariusscheme_zip_url := "https://bitbucket.org/ktakashi/sagittarius-scheme/downloads/sagittarius-0.9.11.zip"
	sagittariusscheme_cmd_zip := exec.Command("curl", "-L", sagittariusscheme_zip_url, "-o", "package.zip")
	err = sagittariusscheme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sagittariusscheme_bin_url := "https://bitbucket.org/ktakashi/sagittarius-scheme/downloads/sagittarius-0.9.11.bin"
	sagittariusscheme_cmd_bin := exec.Command("curl", "-L", sagittariusscheme_bin_url, "-o", "binary.bin")
	err = sagittariusscheme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sagittariusscheme_src_url := "https://bitbucket.org/ktakashi/sagittarius-scheme/downloads/sagittarius-0.9.11.src.tar.gz"
	sagittariusscheme_cmd_src := exec.Command("curl", "-L", sagittariusscheme_src_url, "-o", "source.tar.gz")
	err = sagittariusscheme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sagittariusscheme_cmd_direct := exec.Command("./binary")
	err = sagittariusscheme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: unixodbc")
	exec.Command("latte", "install", "unixodbc").Run()
}
