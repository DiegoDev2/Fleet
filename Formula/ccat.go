package main

// Ccat - Like cat but displays content with syntax highlighting
// Homepage: https://github.com/owenthereal/ccat

import (
	"fmt"
	
	"os/exec"
)

func installCcat() {
	// Método 1: Descargar y extraer .tar.gz
	ccat_tar_url := "https://github.com/owenthereal/ccat/archive/refs/tags/v1.1.0.tar.gz"
	ccat_cmd_tar := exec.Command("curl", "-L", ccat_tar_url, "-o", "package.tar.gz")
	err := ccat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ccat_zip_url := "https://github.com/owenthereal/ccat/archive/refs/tags/v1.1.0.zip"
	ccat_cmd_zip := exec.Command("curl", "-L", ccat_zip_url, "-o", "package.zip")
	err = ccat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ccat_bin_url := "https://github.com/owenthereal/ccat/archive/refs/tags/v1.1.0.bin"
	ccat_cmd_bin := exec.Command("curl", "-L", ccat_bin_url, "-o", "binary.bin")
	err = ccat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ccat_src_url := "https://github.com/owenthereal/ccat/archive/refs/tags/v1.1.0.src.tar.gz"
	ccat_cmd_src := exec.Command("curl", "-L", ccat_src_url, "-o", "source.tar.gz")
	err = ccat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ccat_cmd_direct := exec.Command("./binary")
	err = ccat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
