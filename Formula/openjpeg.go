package main

// Openjpeg - Library for JPEG-2000 image manipulation
// Homepage: https://www.openjpeg.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenjpeg() {
	// Método 1: Descargar y extraer .tar.gz
	openjpeg_tar_url := "https://github.com/uclouvain/openjpeg/archive/refs/tags/v2.5.2.tar.gz"
	openjpeg_cmd_tar := exec.Command("curl", "-L", openjpeg_tar_url, "-o", "package.tar.gz")
	err := openjpeg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openjpeg_zip_url := "https://github.com/uclouvain/openjpeg/archive/refs/tags/v2.5.2.zip"
	openjpeg_cmd_zip := exec.Command("curl", "-L", openjpeg_zip_url, "-o", "package.zip")
	err = openjpeg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openjpeg_bin_url := "https://github.com/uclouvain/openjpeg/archive/refs/tags/v2.5.2.bin"
	openjpeg_cmd_bin := exec.Command("curl", "-L", openjpeg_bin_url, "-o", "binary.bin")
	err = openjpeg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openjpeg_src_url := "https://github.com/uclouvain/openjpeg/archive/refs/tags/v2.5.2.src.tar.gz"
	openjpeg_cmd_src := exec.Command("curl", "-L", openjpeg_src_url, "-o", "source.tar.gz")
	err = openjpeg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openjpeg_cmd_direct := exec.Command("./binary")
	err = openjpeg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
}
