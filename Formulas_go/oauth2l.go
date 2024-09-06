package main

// Oauth2l - Simple CLI for interacting with Google oauth tokens
// Homepage: https://github.com/google/oauth2l

import (
	"fmt"
	
	"os/exec"
)

func installOauth2l() {
	// Método 1: Descargar y extraer .tar.gz
	oauth2l_tar_url := "https://github.com/google/oauth2l/archive/refs/tags/v1.3.0.tar.gz"
	oauth2l_cmd_tar := exec.Command("curl", "-L", oauth2l_tar_url, "-o", "package.tar.gz")
	err := oauth2l_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oauth2l_zip_url := "https://github.com/google/oauth2l/archive/refs/tags/v1.3.0.zip"
	oauth2l_cmd_zip := exec.Command("curl", "-L", oauth2l_zip_url, "-o", "package.zip")
	err = oauth2l_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oauth2l_bin_url := "https://github.com/google/oauth2l/archive/refs/tags/v1.3.0.bin"
	oauth2l_cmd_bin := exec.Command("curl", "-L", oauth2l_bin_url, "-o", "binary.bin")
	err = oauth2l_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oauth2l_src_url := "https://github.com/google/oauth2l/archive/refs/tags/v1.3.0.src.tar.gz"
	oauth2l_cmd_src := exec.Command("curl", "-L", oauth2l_src_url, "-o", "source.tar.gz")
	err = oauth2l_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oauth2l_cmd_direct := exec.Command("./binary")
	err = oauth2l_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
