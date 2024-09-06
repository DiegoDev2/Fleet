package main

// Regipy - Offline registry hive parsing tool
// Homepage: https://github.com/mkorman90/regipy

import (
	"fmt"
	
	"os/exec"
)

func installRegipy() {
	// Método 1: Descargar y extraer .tar.gz
	regipy_tar_url := "https://files.pythonhosted.org/packages/99/03/e54da3d86e833d728322ffcf8d13d7af8aa1bc81c9b5f072e9496897628b/regipy-5.0.0.tar.gz"
	regipy_cmd_tar := exec.Command("curl", "-L", regipy_tar_url, "-o", "package.tar.gz")
	err := regipy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	regipy_zip_url := "https://files.pythonhosted.org/packages/99/03/e54da3d86e833d728322ffcf8d13d7af8aa1bc81c9b5f072e9496897628b/regipy-5.0.0.zip"
	regipy_cmd_zip := exec.Command("curl", "-L", regipy_zip_url, "-o", "package.zip")
	err = regipy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	regipy_bin_url := "https://files.pythonhosted.org/packages/99/03/e54da3d86e833d728322ffcf8d13d7af8aa1bc81c9b5f072e9496897628b/regipy-5.0.0.bin"
	regipy_cmd_bin := exec.Command("curl", "-L", regipy_bin_url, "-o", "binary.bin")
	err = regipy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	regipy_src_url := "https://files.pythonhosted.org/packages/99/03/e54da3d86e833d728322ffcf8d13d7af8aa1bc81c9b5f072e9496897628b/regipy-5.0.0.src.tar.gz"
	regipy_cmd_src := exec.Command("curl", "-L", regipy_src_url, "-o", "source.tar.gz")
	err = regipy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	regipy_cmd_direct := exec.Command("./binary")
	err = regipy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
