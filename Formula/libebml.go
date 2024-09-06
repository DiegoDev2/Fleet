package main

// Libebml - Sort of a sbinary version of XML
// Homepage: https://www.matroska.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibebml() {
	// Método 1: Descargar y extraer .tar.gz
	libebml_tar_url := "https://dl.matroska.org/downloads/libebml/libebml-1.4.5.tar.xz"
	libebml_cmd_tar := exec.Command("curl", "-L", libebml_tar_url, "-o", "package.tar.gz")
	err := libebml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libebml_zip_url := "https://dl.matroska.org/downloads/libebml/libebml-1.4.5.tar.xz"
	libebml_cmd_zip := exec.Command("curl", "-L", libebml_zip_url, "-o", "package.zip")
	err = libebml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libebml_bin_url := "https://dl.matroska.org/downloads/libebml/libebml-1.4.5.tar.xz"
	libebml_cmd_bin := exec.Command("curl", "-L", libebml_bin_url, "-o", "binary.bin")
	err = libebml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libebml_src_url := "https://dl.matroska.org/downloads/libebml/libebml-1.4.5.tar.xz"
	libebml_cmd_src := exec.Command("curl", "-L", libebml_src_url, "-o", "source.tar.gz")
	err = libebml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libebml_cmd_direct := exec.Command("./binary")
	err = libebml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
