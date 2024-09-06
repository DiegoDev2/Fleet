package main

// Bfg - Remove large files or passwords from Git history like git-filter-branch
// Homepage: https://rtyley.github.io/bfg-repo-cleaner/

import (
	"fmt"
	
	"os/exec"
)

func installBfg() {
	// Método 1: Descargar y extraer .tar.gz
	bfg_tar_url := "https://search.maven.org/remotecontent?filepath=com/madgag/bfg/1.14.0/bfg-1.14.0.jar"
	bfg_cmd_tar := exec.Command("curl", "-L", bfg_tar_url, "-o", "package.tar.gz")
	err := bfg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bfg_zip_url := "https://search.maven.org/remotecontent?filepath=com/madgag/bfg/1.14.0/bfg-1.14.0.jar"
	bfg_cmd_zip := exec.Command("curl", "-L", bfg_zip_url, "-o", "package.zip")
	err = bfg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bfg_bin_url := "https://search.maven.org/remotecontent?filepath=com/madgag/bfg/1.14.0/bfg-1.14.0.jar"
	bfg_cmd_bin := exec.Command("curl", "-L", bfg_bin_url, "-o", "binary.bin")
	err = bfg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bfg_src_url := "https://search.maven.org/remotecontent?filepath=com/madgag/bfg/1.14.0/bfg-1.14.0.jar"
	bfg_cmd_src := exec.Command("curl", "-L", bfg_src_url, "-o", "source.tar.gz")
	err = bfg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bfg_cmd_direct := exec.Command("./binary")
	err = bfg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
