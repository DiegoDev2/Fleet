package main

// CmakeLanguageServer - Language Server for CMake
// Homepage: https://github.com/regen100/cmake-language-server

import (
	"fmt"
	
	"os/exec"
)

func installCmakeLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	cmakelanguageserver_tar_url := "https://files.pythonhosted.org/packages/cc/ce/4b14dcaac4359fc9bdcb823763c7984b72e16ff2bf1c709bbc963cc0e0bc/cmake_language_server-0.1.10.tar.gz"
	cmakelanguageserver_cmd_tar := exec.Command("curl", "-L", cmakelanguageserver_tar_url, "-o", "package.tar.gz")
	err := cmakelanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmakelanguageserver_zip_url := "https://files.pythonhosted.org/packages/cc/ce/4b14dcaac4359fc9bdcb823763c7984b72e16ff2bf1c709bbc963cc0e0bc/cmake_language_server-0.1.10.zip"
	cmakelanguageserver_cmd_zip := exec.Command("curl", "-L", cmakelanguageserver_zip_url, "-o", "package.zip")
	err = cmakelanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmakelanguageserver_bin_url := "https://files.pythonhosted.org/packages/cc/ce/4b14dcaac4359fc9bdcb823763c7984b72e16ff2bf1c709bbc963cc0e0bc/cmake_language_server-0.1.10.bin"
	cmakelanguageserver_cmd_bin := exec.Command("curl", "-L", cmakelanguageserver_bin_url, "-o", "binary.bin")
	err = cmakelanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmakelanguageserver_src_url := "https://files.pythonhosted.org/packages/cc/ce/4b14dcaac4359fc9bdcb823763c7984b72e16ff2bf1c709bbc963cc0e0bc/cmake_language_server-0.1.10.src.tar.gz"
	cmakelanguageserver_cmd_src := exec.Command("curl", "-L", cmakelanguageserver_src_url, "-o", "source.tar.gz")
	err = cmakelanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmakelanguageserver_cmd_direct := exec.Command("./binary")
	err = cmakelanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
