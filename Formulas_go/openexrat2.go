package main

// OpenexrAT2 - High dynamic-range image file format
// Homepage: https://www.openexr.com/

import (
	"fmt"
	
	"os/exec"
)

func installOpenexrAT2() {
	// Método 1: Descargar y extraer .tar.gz
	openexrat2_tar_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v2.5.8.tar.gz"
	openexrat2_cmd_tar := exec.Command("curl", "-L", openexrat2_tar_url, "-o", "package.tar.gz")
	err := openexrat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openexrat2_zip_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v2.5.8.zip"
	openexrat2_cmd_zip := exec.Command("curl", "-L", openexrat2_zip_url, "-o", "package.zip")
	err = openexrat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openexrat2_bin_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v2.5.8.bin"
	openexrat2_cmd_bin := exec.Command("curl", "-L", openexrat2_bin_url, "-o", "binary.bin")
	err = openexrat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openexrat2_src_url := "https://github.com/AcademySoftwareFoundation/openexr/archive/refs/tags/v2.5.8.src.tar.gz"
	openexrat2_cmd_src := exec.Command("curl", "-L", openexrat2_src_url, "-o", "source.tar.gz")
	err = openexrat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openexrat2_cmd_direct := exec.Command("./binary")
	err = openexrat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ilmbase")
exec.Command("latte", "install", "ilmbase")
}
