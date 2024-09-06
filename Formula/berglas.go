package main

// Berglas - Tool for managing secrets on Google Cloud
// Homepage: https://github.com/GoogleCloudPlatform/berglas

import (
	"fmt"
	
	"os/exec"
)

func installBerglas() {
	// Método 1: Descargar y extraer .tar.gz
	berglas_tar_url := "https://github.com/GoogleCloudPlatform/berglas/archive/refs/tags/v2.0.2.tar.gz"
	berglas_cmd_tar := exec.Command("curl", "-L", berglas_tar_url, "-o", "package.tar.gz")
	err := berglas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	berglas_zip_url := "https://github.com/GoogleCloudPlatform/berglas/archive/refs/tags/v2.0.2.zip"
	berglas_cmd_zip := exec.Command("curl", "-L", berglas_zip_url, "-o", "package.zip")
	err = berglas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	berglas_bin_url := "https://github.com/GoogleCloudPlatform/berglas/archive/refs/tags/v2.0.2.bin"
	berglas_cmd_bin := exec.Command("curl", "-L", berglas_bin_url, "-o", "binary.bin")
	err = berglas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	berglas_src_url := "https://github.com/GoogleCloudPlatform/berglas/archive/refs/tags/v2.0.2.src.tar.gz"
	berglas_cmd_src := exec.Command("curl", "-L", berglas_src_url, "-o", "source.tar.gz")
	err = berglas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	berglas_cmd_direct := exec.Command("./binary")
	err = berglas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
