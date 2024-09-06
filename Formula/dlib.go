package main

// Dlib - C++ library for machine learning
// Homepage: http://dlib.net/

import (
	"fmt"
	
	"os/exec"
)

func installDlib() {
	// Método 1: Descargar y extraer .tar.gz
	dlib_tar_url := "https://github.com/davisking/dlib/archive/refs/tags/v19.24.6.tar.gz"
	dlib_cmd_tar := exec.Command("curl", "-L", dlib_tar_url, "-o", "package.tar.gz")
	err := dlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dlib_zip_url := "https://github.com/davisking/dlib/archive/refs/tags/v19.24.6.zip"
	dlib_cmd_zip := exec.Command("curl", "-L", dlib_zip_url, "-o", "package.zip")
	err = dlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dlib_bin_url := "https://github.com/davisking/dlib/archive/refs/tags/v19.24.6.bin"
	dlib_cmd_bin := exec.Command("curl", "-L", dlib_bin_url, "-o", "binary.bin")
	err = dlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dlib_src_url := "https://github.com/davisking/dlib/archive/refs/tags/v19.24.6.src.tar.gz"
	dlib_cmd_src := exec.Command("curl", "-L", dlib_src_url, "-o", "source.tar.gz")
	err = dlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dlib_cmd_direct := exec.Command("./binary")
	err = dlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
