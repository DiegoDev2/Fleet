package main

// Pidcat - Colored logcat script to show entries only for specified app
// Homepage: https://github.com/JakeWharton/pidcat

import (
	"fmt"
	
	"os/exec"
)

func installPidcat() {
	// Método 1: Descargar y extraer .tar.gz
	pidcat_tar_url := "https://github.com/JakeWharton/pidcat/archive/refs/tags/2.1.0.tar.gz"
	pidcat_cmd_tar := exec.Command("curl", "-L", pidcat_tar_url, "-o", "package.tar.gz")
	err := pidcat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pidcat_zip_url := "https://github.com/JakeWharton/pidcat/archive/refs/tags/2.1.0.zip"
	pidcat_cmd_zip := exec.Command("curl", "-L", pidcat_zip_url, "-o", "package.zip")
	err = pidcat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pidcat_bin_url := "https://github.com/JakeWharton/pidcat/archive/refs/tags/2.1.0.bin"
	pidcat_cmd_bin := exec.Command("curl", "-L", pidcat_bin_url, "-o", "binary.bin")
	err = pidcat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pidcat_src_url := "https://github.com/JakeWharton/pidcat/archive/refs/tags/2.1.0.src.tar.gz"
	pidcat_cmd_src := exec.Command("curl", "-L", pidcat_src_url, "-o", "source.tar.gz")
	err = pidcat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pidcat_cmd_direct := exec.Command("./binary")
	err = pidcat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
