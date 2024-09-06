package main

// Tfupdate - Update version constraints in your Terraform configurations
// Homepage: https://github.com/minamijoyo/tfupdate

import (
	"fmt"
	
	"os/exec"
)

func installTfupdate() {
	// Método 1: Descargar y extraer .tar.gz
	tfupdate_tar_url := "https://github.com/minamijoyo/tfupdate/archive/refs/tags/v0.8.5.tar.gz"
	tfupdate_cmd_tar := exec.Command("curl", "-L", tfupdate_tar_url, "-o", "package.tar.gz")
	err := tfupdate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tfupdate_zip_url := "https://github.com/minamijoyo/tfupdate/archive/refs/tags/v0.8.5.zip"
	tfupdate_cmd_zip := exec.Command("curl", "-L", tfupdate_zip_url, "-o", "package.zip")
	err = tfupdate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tfupdate_bin_url := "https://github.com/minamijoyo/tfupdate/archive/refs/tags/v0.8.5.bin"
	tfupdate_cmd_bin := exec.Command("curl", "-L", tfupdate_bin_url, "-o", "binary.bin")
	err = tfupdate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tfupdate_src_url := "https://github.com/minamijoyo/tfupdate/archive/refs/tags/v0.8.5.src.tar.gz"
	tfupdate_cmd_src := exec.Command("curl", "-L", tfupdate_src_url, "-o", "source.tar.gz")
	err = tfupdate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tfupdate_cmd_direct := exec.Command("./binary")
	err = tfupdate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
