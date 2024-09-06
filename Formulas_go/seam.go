package main

// Seam - This utility lets you control Seam resources
// Homepage: https://github.com/seamapi/seam-cli

import (
	"fmt"
	
	"os/exec"
)

func installSeam() {
	// Método 1: Descargar y extraer .tar.gz
	seam_tar_url := "https://registry.npmjs.org/seam-cli/-/seam-cli-0.0.58.tgz"
	seam_cmd_tar := exec.Command("curl", "-L", seam_tar_url, "-o", "package.tar.gz")
	err := seam_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	seam_zip_url := "https://registry.npmjs.org/seam-cli/-/seam-cli-0.0.58.tgz"
	seam_cmd_zip := exec.Command("curl", "-L", seam_zip_url, "-o", "package.zip")
	err = seam_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	seam_bin_url := "https://registry.npmjs.org/seam-cli/-/seam-cli-0.0.58.tgz"
	seam_cmd_bin := exec.Command("curl", "-L", seam_bin_url, "-o", "binary.bin")
	err = seam_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	seam_src_url := "https://registry.npmjs.org/seam-cli/-/seam-cli-0.0.58.tgz"
	seam_cmd_src := exec.Command("curl", "-L", seam_src_url, "-o", "source.tar.gz")
	err = seam_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	seam_cmd_direct := exec.Command("./binary")
	err = seam_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
