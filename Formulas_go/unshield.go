package main

// Unshield - Extract files from InstallShield cabinet files
// Homepage: https://github.com/twogood/unshield

import (
	"fmt"
	
	"os/exec"
)

func installUnshield() {
	// Método 1: Descargar y extraer .tar.gz
	unshield_tar_url := "https://github.com/twogood/unshield/archive/refs/tags/1.5.1.tar.gz"
	unshield_cmd_tar := exec.Command("curl", "-L", unshield_tar_url, "-o", "package.tar.gz")
	err := unshield_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unshield_zip_url := "https://github.com/twogood/unshield/archive/refs/tags/1.5.1.zip"
	unshield_cmd_zip := exec.Command("curl", "-L", unshield_zip_url, "-o", "package.zip")
	err = unshield_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unshield_bin_url := "https://github.com/twogood/unshield/archive/refs/tags/1.5.1.bin"
	unshield_cmd_bin := exec.Command("curl", "-L", unshield_bin_url, "-o", "binary.bin")
	err = unshield_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unshield_src_url := "https://github.com/twogood/unshield/archive/refs/tags/1.5.1.src.tar.gz"
	unshield_cmd_src := exec.Command("curl", "-L", unshield_src_url, "-o", "source.tar.gz")
	err = unshield_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unshield_cmd_direct := exec.Command("./binary")
	err = unshield_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
