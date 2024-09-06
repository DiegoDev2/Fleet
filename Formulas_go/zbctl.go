package main

// Zbctl - Zeebe CLI client
// Homepage: https://docs.camunda.io/docs/apis-clients/cli-client/index/

import (
	"fmt"
	
	"os/exec"
)

func installZbctl() {
	// Método 1: Descargar y extraer .tar.gz
	zbctl_tar_url := "https://github.com/camunda/zeebe/archive/refs/tags/8.5.6.tar.gz"
	zbctl_cmd_tar := exec.Command("curl", "-L", zbctl_tar_url, "-o", "package.tar.gz")
	err := zbctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zbctl_zip_url := "https://github.com/camunda/zeebe/archive/refs/tags/8.5.6.zip"
	zbctl_cmd_zip := exec.Command("curl", "-L", zbctl_zip_url, "-o", "package.zip")
	err = zbctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zbctl_bin_url := "https://github.com/camunda/zeebe/archive/refs/tags/8.5.6.bin"
	zbctl_cmd_bin := exec.Command("curl", "-L", zbctl_bin_url, "-o", "binary.bin")
	err = zbctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zbctl_src_url := "https://github.com/camunda/zeebe/archive/refs/tags/8.5.6.src.tar.gz"
	zbctl_cmd_src := exec.Command("curl", "-L", zbctl_src_url, "-o", "source.tar.gz")
	err = zbctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zbctl_cmd_direct := exec.Command("./binary")
	err = zbctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
