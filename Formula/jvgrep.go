package main

// Jvgrep - Grep for Japanese users of Vim
// Homepage: https://github.com/mattn/jvgrep

import (
	"fmt"
	
	"os/exec"
)

func installJvgrep() {
	// Método 1: Descargar y extraer .tar.gz
	jvgrep_tar_url := "https://github.com/mattn/jvgrep/archive/refs/tags/v5.8.12.tar.gz"
	jvgrep_cmd_tar := exec.Command("curl", "-L", jvgrep_tar_url, "-o", "package.tar.gz")
	err := jvgrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jvgrep_zip_url := "https://github.com/mattn/jvgrep/archive/refs/tags/v5.8.12.zip"
	jvgrep_cmd_zip := exec.Command("curl", "-L", jvgrep_zip_url, "-o", "package.zip")
	err = jvgrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jvgrep_bin_url := "https://github.com/mattn/jvgrep/archive/refs/tags/v5.8.12.bin"
	jvgrep_cmd_bin := exec.Command("curl", "-L", jvgrep_bin_url, "-o", "binary.bin")
	err = jvgrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jvgrep_src_url := "https://github.com/mattn/jvgrep/archive/refs/tags/v5.8.12.src.tar.gz"
	jvgrep_cmd_src := exec.Command("curl", "-L", jvgrep_src_url, "-o", "source.tar.gz")
	err = jvgrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jvgrep_cmd_direct := exec.Command("./binary")
	err = jvgrep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
