package main

// Direnv - Load/unload environment variables based on $PWD
// Homepage: https://direnv.net/

import (
	"fmt"
	
	"os/exec"
)

func installDirenv() {
	// Método 1: Descargar y extraer .tar.gz
	direnv_tar_url := "https://github.com/direnv/direnv/archive/refs/tags/v2.34.0.tar.gz"
	direnv_cmd_tar := exec.Command("curl", "-L", direnv_tar_url, "-o", "package.tar.gz")
	err := direnv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	direnv_zip_url := "https://github.com/direnv/direnv/archive/refs/tags/v2.34.0.zip"
	direnv_cmd_zip := exec.Command("curl", "-L", direnv_zip_url, "-o", "package.zip")
	err = direnv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	direnv_bin_url := "https://github.com/direnv/direnv/archive/refs/tags/v2.34.0.bin"
	direnv_cmd_bin := exec.Command("curl", "-L", direnv_bin_url, "-o", "binary.bin")
	err = direnv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	direnv_src_url := "https://github.com/direnv/direnv/archive/refs/tags/v2.34.0.src.tar.gz"
	direnv_cmd_src := exec.Command("curl", "-L", direnv_src_url, "-o", "source.tar.gz")
	err = direnv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	direnv_cmd_direct := exec.Command("./binary")
	err = direnv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
}
