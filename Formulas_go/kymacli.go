package main

// KymaCli - Kyma command-line interface
// Homepage: https://kyma-project.io

import (
	"fmt"
	
	"os/exec"
)

func installKymaCli() {
	// Método 1: Descargar y extraer .tar.gz
	kymacli_tar_url := "https://github.com/kyma-project/cli/archive/refs/tags/2.20.5.tar.gz"
	kymacli_cmd_tar := exec.Command("curl", "-L", kymacli_tar_url, "-o", "package.tar.gz")
	err := kymacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kymacli_zip_url := "https://github.com/kyma-project/cli/archive/refs/tags/2.20.5.zip"
	kymacli_cmd_zip := exec.Command("curl", "-L", kymacli_zip_url, "-o", "package.zip")
	err = kymacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kymacli_bin_url := "https://github.com/kyma-project/cli/archive/refs/tags/2.20.5.bin"
	kymacli_cmd_bin := exec.Command("curl", "-L", kymacli_bin_url, "-o", "binary.bin")
	err = kymacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kymacli_src_url := "https://github.com/kyma-project/cli/archive/refs/tags/2.20.5.src.tar.gz"
	kymacli_cmd_src := exec.Command("curl", "-L", kymacli_src_url, "-o", "source.tar.gz")
	err = kymacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kymacli_cmd_direct := exec.Command("./binary")
	err = kymacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
