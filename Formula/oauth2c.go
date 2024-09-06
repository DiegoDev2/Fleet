package main

// Oauth2c - User-friendly CLI for OAuth2
// Homepage: https://github.com/cloudentity/oauth2c

import (
	"fmt"
	
	"os/exec"
)

func installOauth2c() {
	// Método 1: Descargar y extraer .tar.gz
	oauth2c_tar_url := "https://github.com/cloudentity/oauth2c/archive/refs/tags/v1.16.0.tar.gz"
	oauth2c_cmd_tar := exec.Command("curl", "-L", oauth2c_tar_url, "-o", "package.tar.gz")
	err := oauth2c_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oauth2c_zip_url := "https://github.com/cloudentity/oauth2c/archive/refs/tags/v1.16.0.zip"
	oauth2c_cmd_zip := exec.Command("curl", "-L", oauth2c_zip_url, "-o", "package.zip")
	err = oauth2c_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oauth2c_bin_url := "https://github.com/cloudentity/oauth2c/archive/refs/tags/v1.16.0.bin"
	oauth2c_cmd_bin := exec.Command("curl", "-L", oauth2c_bin_url, "-o", "binary.bin")
	err = oauth2c_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oauth2c_src_url := "https://github.com/cloudentity/oauth2c/archive/refs/tags/v1.16.0.src.tar.gz"
	oauth2c_cmd_src := exec.Command("curl", "-L", oauth2c_src_url, "-o", "source.tar.gz")
	err = oauth2c_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oauth2c_cmd_direct := exec.Command("./binary")
	err = oauth2c_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
