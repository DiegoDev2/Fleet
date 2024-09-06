package main

// Jd - JSON diff and patch
// Homepage: https://github.com/josephburnett/jd

import (
	"fmt"
	
	"os/exec"
)

func installJd() {
	// Método 1: Descargar y extraer .tar.gz
	jd_tar_url := "https://github.com/josephburnett/jd/archive/refs/tags/v1.9.1.tar.gz"
	jd_cmd_tar := exec.Command("curl", "-L", jd_tar_url, "-o", "package.tar.gz")
	err := jd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jd_zip_url := "https://github.com/josephburnett/jd/archive/refs/tags/v1.9.1.zip"
	jd_cmd_zip := exec.Command("curl", "-L", jd_zip_url, "-o", "package.zip")
	err = jd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jd_bin_url := "https://github.com/josephburnett/jd/archive/refs/tags/v1.9.1.bin"
	jd_cmd_bin := exec.Command("curl", "-L", jd_bin_url, "-o", "binary.bin")
	err = jd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jd_src_url := "https://github.com/josephburnett/jd/archive/refs/tags/v1.9.1.src.tar.gz"
	jd_cmd_src := exec.Command("curl", "-L", jd_src_url, "-o", "source.tar.gz")
	err = jd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jd_cmd_direct := exec.Command("./binary")
	err = jd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
