package main

// Pmd - Source code analyzer for Java, JavaScript, and more
// Homepage: https://pmd.github.io

import (
	"fmt"
	
	"os/exec"
)

func installPmd() {
	// Método 1: Descargar y extraer .tar.gz
	pmd_tar_url := "https://github.com/pmd/pmd/releases/download/pmd_releases%2F7.5.0/pmd-dist-7.5.0-bin.zip"
	pmd_cmd_tar := exec.Command("curl", "-L", pmd_tar_url, "-o", "package.tar.gz")
	err := pmd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pmd_zip_url := "https://github.com/pmd/pmd/releases/download/pmd_releases%2F7.5.0/pmd-dist-7.5.0-bin.zip"
	pmd_cmd_zip := exec.Command("curl", "-L", pmd_zip_url, "-o", "package.zip")
	err = pmd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pmd_bin_url := "https://github.com/pmd/pmd/releases/download/pmd_releases%2F7.5.0/pmd-dist-7.5.0-bin.zip"
	pmd_cmd_bin := exec.Command("curl", "-L", pmd_bin_url, "-o", "binary.bin")
	err = pmd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pmd_src_url := "https://github.com/pmd/pmd/releases/download/pmd_releases%2F7.5.0/pmd-dist-7.5.0-bin.zip"
	pmd_cmd_src := exec.Command("curl", "-L", pmd_src_url, "-o", "source.tar.gz")
	err = pmd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pmd_cmd_direct := exec.Command("./binary")
	err = pmd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
