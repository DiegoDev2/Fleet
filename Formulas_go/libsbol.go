package main

// Libsbol - Read and write files in the Synthetic Biology Open Language (SBOL)
// Homepage: https://synbiodex.github.io/libSBOL

import (
	"fmt"
	
	"os/exec"
)

func installLibsbol() {
	// Método 1: Descargar y extraer .tar.gz
	libsbol_tar_url := "https://github.com/SynBioDex/libSBOL/archive/refs/tags/v2.3.2.tar.gz"
	libsbol_cmd_tar := exec.Command("curl", "-L", libsbol_tar_url, "-o", "package.tar.gz")
	err := libsbol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsbol_zip_url := "https://github.com/SynBioDex/libSBOL/archive/refs/tags/v2.3.2.zip"
	libsbol_cmd_zip := exec.Command("curl", "-L", libsbol_zip_url, "-o", "package.zip")
	err = libsbol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsbol_bin_url := "https://github.com/SynBioDex/libSBOL/archive/refs/tags/v2.3.2.bin"
	libsbol_cmd_bin := exec.Command("curl", "-L", libsbol_bin_url, "-o", "binary.bin")
	err = libsbol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsbol_src_url := "https://github.com/SynBioDex/libSBOL/archive/refs/tags/v2.3.2.src.tar.gz"
	libsbol_cmd_src := exec.Command("curl", "-L", libsbol_src_url, "-o", "source.tar.gz")
	err = libsbol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsbol_cmd_direct := exec.Command("./binary")
	err = libsbol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: jsoncpp")
exec.Command("latte", "install", "jsoncpp")
	fmt.Println("Instalando dependencia: raptor")
exec.Command("latte", "install", "raptor")
	fmt.Println("Instalando dependencia: rasqal")
exec.Command("latte", "install", "rasqal")
}
