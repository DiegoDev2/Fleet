package main

// Openexr - High dynamic-range image file format
// Homepage: https://www.openexr.com/

import (
	"fmt"
	
	"os/exec"
)

func installOpenexr() {
	// Método 1: Descargar y extraer .tar.gz
	openexr_tar_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v3.2.4.tar.gz"
	openexr_cmd_tar := exec.Command("curl", "-L", openexr_tar_url, "-o", "package.tar.gz")
	err := openexr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openexr_zip_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v3.2.4.zip"
	openexr_cmd_zip := exec.Command("curl", "-L", openexr_zip_url, "-o", "package.zip")
	err = openexr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openexr_bin_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v3.2.4.bin"
	openexr_cmd_bin := exec.Command("curl", "-L", openexr_bin_url, "-o", "binary.bin")
	err = openexr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openexr_src_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v3.2.4.src.tar.gz"
	openexr_cmd_src := exec.Command("curl", "-L", openexr_src_url, "-o", "source.tar.gz")
	err = openexr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openexr_cmd_direct := exec.Command("./binary")
	err = openexr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: imath")
	exec.Command("latte", "install", "imath").Run()
}
