package main

// Atmos - Universal Tool for DevOps and Cloud Automation
// Homepage: https://github.com/cloudposse/atmos

import (
	"fmt"
	
	"os/exec"
)

func installAtmos() {
	// Método 1: Descargar y extraer .tar.gz
	atmos_tar_url := "https://github.com/cloudposse/atmos/archive/refs/tags/v1.88.1.tar.gz"
	atmos_cmd_tar := exec.Command("curl", "-L", atmos_tar_url, "-o", "package.tar.gz")
	err := atmos_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atmos_zip_url := "https://github.com/cloudposse/atmos/archive/refs/tags/v1.88.1.zip"
	atmos_cmd_zip := exec.Command("curl", "-L", atmos_zip_url, "-o", "package.zip")
	err = atmos_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atmos_bin_url := "https://github.com/cloudposse/atmos/archive/refs/tags/v1.88.1.bin"
	atmos_cmd_bin := exec.Command("curl", "-L", atmos_bin_url, "-o", "binary.bin")
	err = atmos_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atmos_src_url := "https://github.com/cloudposse/atmos/archive/refs/tags/v1.88.1.src.tar.gz"
	atmos_cmd_src := exec.Command("curl", "-L", atmos_src_url, "-o", "source.tar.gz")
	err = atmos_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atmos_cmd_direct := exec.Command("./binary")
	err = atmos_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
