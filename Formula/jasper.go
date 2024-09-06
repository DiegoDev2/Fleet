package main

// Jasper - Library for manipulating JPEG-2000 images
// Homepage: https://ece.engr.uvic.ca/~frodo/jasper/

import (
	"fmt"
	
	"os/exec"
)

func installJasper() {
	// Método 1: Descargar y extraer .tar.gz
	jasper_tar_url := "https://github.com/jasper-software/jasper/releases/download/version-4.2.4/jasper-4.2.4.tar.gz"
	jasper_cmd_tar := exec.Command("curl", "-L", jasper_tar_url, "-o", "package.tar.gz")
	err := jasper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jasper_zip_url := "https://github.com/jasper-software/jasper/releases/download/version-4.2.4/jasper-4.2.4.zip"
	jasper_cmd_zip := exec.Command("curl", "-L", jasper_zip_url, "-o", "package.zip")
	err = jasper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jasper_bin_url := "https://github.com/jasper-software/jasper/releases/download/version-4.2.4/jasper-4.2.4.bin"
	jasper_cmd_bin := exec.Command("curl", "-L", jasper_bin_url, "-o", "binary.bin")
	err = jasper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jasper_src_url := "https://github.com/jasper-software/jasper/releases/download/version-4.2.4/jasper-4.2.4.src.tar.gz"
	jasper_cmd_src := exec.Command("curl", "-L", jasper_src_url, "-o", "source.tar.gz")
	err = jasper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jasper_cmd_direct := exec.Command("./binary")
	err = jasper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
}
