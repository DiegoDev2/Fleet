package main

// Ctemplate - Template language for C++
// Homepage: https://github.com/olafvdspek/ctemplate

import (
	"fmt"
	
	"os/exec"
)

func installCtemplate() {
	// Método 1: Descargar y extraer .tar.gz
	ctemplate_tar_url := "https://github.com/OlafvdSpek/ctemplate/archive/refs/tags/ctemplate-2.4.tar.gz"
	ctemplate_cmd_tar := exec.Command("curl", "-L", ctemplate_tar_url, "-o", "package.tar.gz")
	err := ctemplate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ctemplate_zip_url := "https://github.com/OlafvdSpek/ctemplate/archive/refs/tags/ctemplate-2.4.zip"
	ctemplate_cmd_zip := exec.Command("curl", "-L", ctemplate_zip_url, "-o", "package.zip")
	err = ctemplate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ctemplate_bin_url := "https://github.com/OlafvdSpek/ctemplate/archive/refs/tags/ctemplate-2.4.bin"
	ctemplate_cmd_bin := exec.Command("curl", "-L", ctemplate_bin_url, "-o", "binary.bin")
	err = ctemplate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ctemplate_src_url := "https://github.com/OlafvdSpek/ctemplate/archive/refs/tags/ctemplate-2.4.src.tar.gz"
	ctemplate_cmd_src := exec.Command("curl", "-L", ctemplate_src_url, "-o", "source.tar.gz")
	err = ctemplate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ctemplate_cmd_direct := exec.Command("./binary")
	err = ctemplate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
