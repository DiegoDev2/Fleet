package main

// Livestreamer - Pipes video from streaming services into a player such as VLC
// Homepage: https://livestreamer.io/

import (
	"fmt"
	
	"os/exec"
)

func installLivestreamer() {
	// Método 1: Descargar y extraer .tar.gz
	livestreamer_tar_url := "https://files.pythonhosted.org/packages/ee/d6/efbe3456160a2c62e3dd841c5d9504d071c94449a819148bb038b50d862a/livestreamer-1.12.2.tar.gz"
	livestreamer_cmd_tar := exec.Command("curl", "-L", livestreamer_tar_url, "-o", "package.tar.gz")
	err := livestreamer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	livestreamer_zip_url := "https://files.pythonhosted.org/packages/ee/d6/efbe3456160a2c62e3dd841c5d9504d071c94449a819148bb038b50d862a/livestreamer-1.12.2.zip"
	livestreamer_cmd_zip := exec.Command("curl", "-L", livestreamer_zip_url, "-o", "package.zip")
	err = livestreamer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	livestreamer_bin_url := "https://files.pythonhosted.org/packages/ee/d6/efbe3456160a2c62e3dd841c5d9504d071c94449a819148bb038b50d862a/livestreamer-1.12.2.bin"
	livestreamer_cmd_bin := exec.Command("curl", "-L", livestreamer_bin_url, "-o", "binary.bin")
	err = livestreamer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	livestreamer_src_url := "https://files.pythonhosted.org/packages/ee/d6/efbe3456160a2c62e3dd841c5d9504d071c94449a819148bb038b50d862a/livestreamer-1.12.2.src.tar.gz"
	livestreamer_cmd_src := exec.Command("curl", "-L", livestreamer_src_url, "-o", "source.tar.gz")
	err = livestreamer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	livestreamer_cmd_direct := exec.Command("./binary")
	err = livestreamer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.10")
	exec.Command("latte", "install", "python@3.10").Run()
}
