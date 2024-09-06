package main

// Ratchet - Tool for securing CI/CD workflows with version pinning
// Homepage: https://github.com/sethvargo/ratchet

import (
	"fmt"
	
	"os/exec"
)

func installRatchet() {
	// Método 1: Descargar y extraer .tar.gz
	ratchet_tar_url := "https://github.com/sethvargo/ratchet/archive/refs/tags/v0.10.0.tar.gz"
	ratchet_cmd_tar := exec.Command("curl", "-L", ratchet_tar_url, "-o", "package.tar.gz")
	err := ratchet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ratchet_zip_url := "https://github.com/sethvargo/ratchet/archive/refs/tags/v0.10.0.zip"
	ratchet_cmd_zip := exec.Command("curl", "-L", ratchet_zip_url, "-o", "package.zip")
	err = ratchet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ratchet_bin_url := "https://github.com/sethvargo/ratchet/archive/refs/tags/v0.10.0.bin"
	ratchet_cmd_bin := exec.Command("curl", "-L", ratchet_bin_url, "-o", "binary.bin")
	err = ratchet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ratchet_src_url := "https://github.com/sethvargo/ratchet/archive/refs/tags/v0.10.0.src.tar.gz"
	ratchet_cmd_src := exec.Command("curl", "-L", ratchet_src_url, "-o", "source.tar.gz")
	err = ratchet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ratchet_cmd_direct := exec.Command("./binary")
	err = ratchet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
