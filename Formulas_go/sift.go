package main

// Sift - Fast and powerful open source alternative to grep
// Homepage: https://sift-tool.org

import (
	"fmt"
	
	"os/exec"
)

func installSift() {
	// Método 1: Descargar y extraer .tar.gz
	sift_tar_url := "https://github.com/svent/sift/archive/refs/tags/v0.9.0.tar.gz"
	sift_cmd_tar := exec.Command("curl", "-L", sift_tar_url, "-o", "package.tar.gz")
	err := sift_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sift_zip_url := "https://github.com/svent/sift/archive/refs/tags/v0.9.0.zip"
	sift_cmd_zip := exec.Command("curl", "-L", sift_zip_url, "-o", "package.zip")
	err = sift_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sift_bin_url := "https://github.com/svent/sift/archive/refs/tags/v0.9.0.bin"
	sift_cmd_bin := exec.Command("curl", "-L", sift_bin_url, "-o", "binary.bin")
	err = sift_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sift_src_url := "https://github.com/svent/sift/archive/refs/tags/v0.9.0.src.tar.gz"
	sift_cmd_src := exec.Command("curl", "-L", sift_src_url, "-o", "source.tar.gz")
	err = sift_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sift_cmd_direct := exec.Command("./binary")
	err = sift_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
