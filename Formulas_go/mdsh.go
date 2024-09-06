package main

// Mdsh - Markdown shell pre-processor
// Homepage: https://zimbatm.github.io/mdsh/

import (
	"fmt"
	
	"os/exec"
)

func installMdsh() {
	// Método 1: Descargar y extraer .tar.gz
	mdsh_tar_url := "https://github.com/zimbatm/mdsh/archive/refs/tags/v0.9.0.tar.gz"
	mdsh_cmd_tar := exec.Command("curl", "-L", mdsh_tar_url, "-o", "package.tar.gz")
	err := mdsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdsh_zip_url := "https://github.com/zimbatm/mdsh/archive/refs/tags/v0.9.0.zip"
	mdsh_cmd_zip := exec.Command("curl", "-L", mdsh_zip_url, "-o", "package.zip")
	err = mdsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdsh_bin_url := "https://github.com/zimbatm/mdsh/archive/refs/tags/v0.9.0.bin"
	mdsh_cmd_bin := exec.Command("curl", "-L", mdsh_bin_url, "-o", "binary.bin")
	err = mdsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdsh_src_url := "https://github.com/zimbatm/mdsh/archive/refs/tags/v0.9.0.src.tar.gz"
	mdsh_cmd_src := exec.Command("curl", "-L", mdsh_src_url, "-o", "source.tar.gz")
	err = mdsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdsh_cmd_direct := exec.Command("./binary")
	err = mdsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
