package main

// Sonic - Fast, lightweight & schema-less search backend
// Homepage: https://github.com/valeriansaliou/sonic

import (
	"fmt"
	
	"os/exec"
)

func installSonic() {
	// Método 1: Descargar y extraer .tar.gz
	sonic_tar_url := "https://github.com/valeriansaliou/sonic/archive/refs/tags/v1.4.9.tar.gz"
	sonic_cmd_tar := exec.Command("curl", "-L", sonic_tar_url, "-o", "package.tar.gz")
	err := sonic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sonic_zip_url := "https://github.com/valeriansaliou/sonic/archive/refs/tags/v1.4.9.zip"
	sonic_cmd_zip := exec.Command("curl", "-L", sonic_zip_url, "-o", "package.zip")
	err = sonic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sonic_bin_url := "https://github.com/valeriansaliou/sonic/archive/refs/tags/v1.4.9.bin"
	sonic_cmd_bin := exec.Command("curl", "-L", sonic_bin_url, "-o", "binary.bin")
	err = sonic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sonic_src_url := "https://github.com/valeriansaliou/sonic/archive/refs/tags/v1.4.9.src.tar.gz"
	sonic_cmd_src := exec.Command("curl", "-L", sonic_src_url, "-o", "source.tar.gz")
	err = sonic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sonic_cmd_direct := exec.Command("./binary")
	err = sonic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
