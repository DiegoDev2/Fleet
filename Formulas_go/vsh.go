package main

// Vsh - HashiCorp Vault interactive shell
// Homepage: https://github.com/fishi0x01/vsh

import (
	"fmt"
	
	"os/exec"
)

func installVsh() {
	// Método 1: Descargar y extraer .tar.gz
	vsh_tar_url := "https://github.com/fishi0x01/vsh/archive/refs/tags/v0.13.0.tar.gz"
	vsh_cmd_tar := exec.Command("curl", "-L", vsh_tar_url, "-o", "package.tar.gz")
	err := vsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vsh_zip_url := "https://github.com/fishi0x01/vsh/archive/refs/tags/v0.13.0.zip"
	vsh_cmd_zip := exec.Command("curl", "-L", vsh_zip_url, "-o", "package.zip")
	err = vsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vsh_bin_url := "https://github.com/fishi0x01/vsh/archive/refs/tags/v0.13.0.bin"
	vsh_cmd_bin := exec.Command("curl", "-L", vsh_bin_url, "-o", "binary.bin")
	err = vsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vsh_src_url := "https://github.com/fishi0x01/vsh/archive/refs/tags/v0.13.0.src.tar.gz"
	vsh_cmd_src := exec.Command("curl", "-L", vsh_src_url, "-o", "source.tar.gz")
	err = vsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vsh_cmd_direct := exec.Command("./binary")
	err = vsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
