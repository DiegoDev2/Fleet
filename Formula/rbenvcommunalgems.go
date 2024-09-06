package main

// RbenvCommunalGems - Share gems across multiple rbenv Ruby installs
// Homepage: https://github.com/tpope/rbenv-communal-gems

import (
	"fmt"
	
	"os/exec"
)

func installRbenvCommunalGems() {
	// Método 1: Descargar y extraer .tar.gz
	rbenvcommunalgems_tar_url := "https://github.com/tpope/rbenv-communal-gems/archive/refs/tags/v1.0.1.tar.gz"
	rbenvcommunalgems_cmd_tar := exec.Command("curl", "-L", rbenvcommunalgems_tar_url, "-o", "package.tar.gz")
	err := rbenvcommunalgems_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbenvcommunalgems_zip_url := "https://github.com/tpope/rbenv-communal-gems/archive/refs/tags/v1.0.1.zip"
	rbenvcommunalgems_cmd_zip := exec.Command("curl", "-L", rbenvcommunalgems_zip_url, "-o", "package.zip")
	err = rbenvcommunalgems_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbenvcommunalgems_bin_url := "https://github.com/tpope/rbenv-communal-gems/archive/refs/tags/v1.0.1.bin"
	rbenvcommunalgems_cmd_bin := exec.Command("curl", "-L", rbenvcommunalgems_bin_url, "-o", "binary.bin")
	err = rbenvcommunalgems_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbenvcommunalgems_src_url := "https://github.com/tpope/rbenv-communal-gems/archive/refs/tags/v1.0.1.src.tar.gz"
	rbenvcommunalgems_cmd_src := exec.Command("curl", "-L", rbenvcommunalgems_src_url, "-o", "source.tar.gz")
	err = rbenvcommunalgems_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbenvcommunalgems_cmd_direct := exec.Command("./binary")
	err = rbenvcommunalgems_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rbenv")
	exec.Command("latte", "install", "rbenv").Run()
}
