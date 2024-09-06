package main

// Stress - Tool to impose load on and stress test a computer system
// Homepage: https://github.com/resurrecting-open-source-projects/stress

import (
	"fmt"
	
	"os/exec"
)

func installStress() {
	// Método 1: Descargar y extraer .tar.gz
	stress_tar_url := "https://github.com/resurrecting-open-source-projects/stress/archive/refs/tags/1.0.7.tar.gz"
	stress_cmd_tar := exec.Command("curl", "-L", stress_tar_url, "-o", "package.tar.gz")
	err := stress_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stress_zip_url := "https://github.com/resurrecting-open-source-projects/stress/archive/refs/tags/1.0.7.zip"
	stress_cmd_zip := exec.Command("curl", "-L", stress_zip_url, "-o", "package.zip")
	err = stress_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stress_bin_url := "https://github.com/resurrecting-open-source-projects/stress/archive/refs/tags/1.0.7.bin"
	stress_cmd_bin := exec.Command("curl", "-L", stress_bin_url, "-o", "binary.bin")
	err = stress_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stress_src_url := "https://github.com/resurrecting-open-source-projects/stress/archive/refs/tags/1.0.7.src.tar.gz"
	stress_cmd_src := exec.Command("curl", "-L", stress_src_url, "-o", "source.tar.gz")
	err = stress_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stress_cmd_direct := exec.Command("./binary")
	err = stress_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
