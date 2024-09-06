package main

// CrushTools - Command-line tools for processing delimited text data
// Homepage: https://github.com/google/crush-tools

import (
	"fmt"
	
	"os/exec"
)

func installCrushTools() {
	// Método 1: Descargar y extraer .tar.gz
	crushtools_tar_url := "https://github.com/google/crush-tools/releases/download/20150716/crush-tools-20150716.tar.gz"
	crushtools_cmd_tar := exec.Command("curl", "-L", crushtools_tar_url, "-o", "package.tar.gz")
	err := crushtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crushtools_zip_url := "https://github.com/google/crush-tools/releases/download/20150716/crush-tools-20150716.zip"
	crushtools_cmd_zip := exec.Command("curl", "-L", crushtools_zip_url, "-o", "package.zip")
	err = crushtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crushtools_bin_url := "https://github.com/google/crush-tools/releases/download/20150716/crush-tools-20150716.bin"
	crushtools_cmd_bin := exec.Command("curl", "-L", crushtools_bin_url, "-o", "binary.bin")
	err = crushtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crushtools_src_url := "https://github.com/google/crush-tools/releases/download/20150716/crush-tools-20150716.src.tar.gz"
	crushtools_cmd_src := exec.Command("curl", "-L", crushtools_src_url, "-o", "source.tar.gz")
	err = crushtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crushtools_cmd_direct := exec.Command("./binary")
	err = crushtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
