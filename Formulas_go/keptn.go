package main

// Keptn - CLI for keptn.sh, a message-driven control-plane for application delivery
// Homepage: https://keptn.sh

import (
	"fmt"
	
	"os/exec"
)

func installKeptn() {
	// Método 1: Descargar y extraer .tar.gz
	keptn_tar_url := "https://github.com/keptn/keptn/archive/refs/tags/1.4.5.tar.gz"
	keptn_cmd_tar := exec.Command("curl", "-L", keptn_tar_url, "-o", "package.tar.gz")
	err := keptn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	keptn_zip_url := "https://github.com/keptn/keptn/archive/refs/tags/1.4.5.zip"
	keptn_cmd_zip := exec.Command("curl", "-L", keptn_zip_url, "-o", "package.zip")
	err = keptn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	keptn_bin_url := "https://github.com/keptn/keptn/archive/refs/tags/1.4.5.bin"
	keptn_cmd_bin := exec.Command("curl", "-L", keptn_bin_url, "-o", "binary.bin")
	err = keptn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	keptn_src_url := "https://github.com/keptn/keptn/archive/refs/tags/1.4.5.src.tar.gz"
	keptn_cmd_src := exec.Command("curl", "-L", keptn_src_url, "-o", "source.tar.gz")
	err = keptn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	keptn_cmd_direct := exec.Command("./binary")
	err = keptn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
