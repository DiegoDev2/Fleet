package main

// Ilmbase - OpenEXR ILM Base libraries (high dynamic-range image file format)
// Homepage: https://www.openexr.com/

import (
	"fmt"
	
	"os/exec"
)

func installIlmbase() {
	// Método 1: Descargar y extraer .tar.gz
	ilmbase_tar_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v2.5.8.tar.gz"
	ilmbase_cmd_tar := exec.Command("curl", "-L", ilmbase_tar_url, "-o", "package.tar.gz")
	err := ilmbase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ilmbase_zip_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v2.5.8.zip"
	ilmbase_cmd_zip := exec.Command("curl", "-L", ilmbase_zip_url, "-o", "package.zip")
	err = ilmbase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ilmbase_bin_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v2.5.8.bin"
	ilmbase_cmd_bin := exec.Command("curl", "-L", ilmbase_bin_url, "-o", "binary.bin")
	err = ilmbase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ilmbase_src_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v2.5.8.src.tar.gz"
	ilmbase_cmd_src := exec.Command("curl", "-L", ilmbase_src_url, "-o", "source.tar.gz")
	err = ilmbase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ilmbase_cmd_direct := exec.Command("./binary")
	err = ilmbase_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
