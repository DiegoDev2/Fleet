package main

// EfmLangserver - General purpose Language Server
// Homepage: https://github.com/mattn/efm-langserver

import (
	"fmt"
	
	"os/exec"
)

func installEfmLangserver() {
	// Método 1: Descargar y extraer .tar.gz
	efmlangserver_tar_url := "https://github.com/mattn/efm-langserver/archive/refs/tags/v0.0.53.tar.gz"
	efmlangserver_cmd_tar := exec.Command("curl", "-L", efmlangserver_tar_url, "-o", "package.tar.gz")
	err := efmlangserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	efmlangserver_zip_url := "https://github.com/mattn/efm-langserver/archive/refs/tags/v0.0.53.zip"
	efmlangserver_cmd_zip := exec.Command("curl", "-L", efmlangserver_zip_url, "-o", "package.zip")
	err = efmlangserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	efmlangserver_bin_url := "https://github.com/mattn/efm-langserver/archive/refs/tags/v0.0.53.bin"
	efmlangserver_cmd_bin := exec.Command("curl", "-L", efmlangserver_bin_url, "-o", "binary.bin")
	err = efmlangserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	efmlangserver_src_url := "https://github.com/mattn/efm-langserver/archive/refs/tags/v0.0.53.src.tar.gz"
	efmlangserver_cmd_src := exec.Command("curl", "-L", efmlangserver_src_url, "-o", "source.tar.gz")
	err = efmlangserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	efmlangserver_cmd_direct := exec.Command("./binary")
	err = efmlangserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
