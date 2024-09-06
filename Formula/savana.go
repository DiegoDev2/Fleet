package main

// Savana - Transactional workspaces for SVN
// Homepage: https://github.com/codehaus/savana

import (
	"fmt"
	
	"os/exec"
)

func installSavana() {
	// Método 1: Descargar y extraer .tar.gz
	savana_tar_url := "https://search.maven.org/remotecontent?filepath=org/codehaus/savana/1.2/savana-1.2-install.tar.gz"
	savana_cmd_tar := exec.Command("curl", "-L", savana_tar_url, "-o", "package.tar.gz")
	err := savana_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	savana_zip_url := "https://search.maven.org/remotecontent?filepath=org/codehaus/savana/1.2/savana-1.2-install.zip"
	savana_cmd_zip := exec.Command("curl", "-L", savana_zip_url, "-o", "package.zip")
	err = savana_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	savana_bin_url := "https://search.maven.org/remotecontent?filepath=org/codehaus/savana/1.2/savana-1.2-install.bin"
	savana_cmd_bin := exec.Command("curl", "-L", savana_bin_url, "-o", "binary.bin")
	err = savana_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	savana_src_url := "https://search.maven.org/remotecontent?filepath=org/codehaus/savana/1.2/savana-1.2-install.src.tar.gz"
	savana_cmd_src := exec.Command("curl", "-L", savana_src_url, "-o", "source.tar.gz")
	err = savana_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	savana_cmd_direct := exec.Command("./binary")
	err = savana_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
