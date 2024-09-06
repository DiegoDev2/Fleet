package main

// Imath - Library of 2D and 3D vector, matrix, and math operations
// Homepage: https://www.openexr.com/

import (
	"fmt"
	
	"os/exec"
)

func installImath() {
	// Método 1: Descargar y extraer .tar.gz
	imath_tar_url := "https://github.com/AcademySoftwareFoundation/Imath/archive/refs/tags/v3.1.11.tar.gz"
	imath_cmd_tar := exec.Command("curl", "-L", imath_tar_url, "-o", "package.tar.gz")
	err := imath_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imath_zip_url := "https://github.com/AcademySoftwareFoundation/Imath/archive/refs/tags/v3.1.11.zip"
	imath_cmd_zip := exec.Command("curl", "-L", imath_zip_url, "-o", "package.zip")
	err = imath_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imath_bin_url := "https://github.com/AcademySoftwareFoundation/Imath/archive/refs/tags/v3.1.11.bin"
	imath_cmd_bin := exec.Command("curl", "-L", imath_bin_url, "-o", "binary.bin")
	err = imath_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imath_src_url := "https://github.com/AcademySoftwareFoundation/Imath/archive/refs/tags/v3.1.11.src.tar.gz"
	imath_cmd_src := exec.Command("curl", "-L", imath_src_url, "-o", "source.tar.gz")
	err = imath_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imath_cmd_direct := exec.Command("./binary")
	err = imath_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
