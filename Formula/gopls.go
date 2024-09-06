package main

// Gopls - Language server for the Go language
// Homepage: https://github.com/golang/tools/tree/master/gopls

import (
	"fmt"
	
	"os/exec"
)

func installGopls() {
	// Método 1: Descargar y extraer .tar.gz
	gopls_tar_url := "https://github.com/golang/tools/archive/refs/tags/gopls/v0.16.2.tar.gz"
	gopls_cmd_tar := exec.Command("curl", "-L", gopls_tar_url, "-o", "package.tar.gz")
	err := gopls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gopls_zip_url := "https://github.com/golang/tools/archive/refs/tags/gopls/v0.16.2.zip"
	gopls_cmd_zip := exec.Command("curl", "-L", gopls_zip_url, "-o", "package.zip")
	err = gopls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gopls_bin_url := "https://github.com/golang/tools/archive/refs/tags/gopls/v0.16.2.bin"
	gopls_cmd_bin := exec.Command("curl", "-L", gopls_bin_url, "-o", "binary.bin")
	err = gopls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gopls_src_url := "https://github.com/golang/tools/archive/refs/tags/gopls/v0.16.2.src.tar.gz"
	gopls_cmd_src := exec.Command("curl", "-L", gopls_src_url, "-o", "source.tar.gz")
	err = gopls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gopls_cmd_direct := exec.Command("./binary")
	err = gopls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
