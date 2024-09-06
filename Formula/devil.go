package main

// Devil - Cross-platform image library
// Homepage: https://sourceforge.net/projects/openil/

import (
	"fmt"
	
	"os/exec"
)

func installDevil() {
	// Método 1: Descargar y extraer .tar.gz
	devil_tar_url := "https://downloads.sourceforge.net/project/openil/DevIL/1.8.0/DevIL-1.8.0.tar.gz"
	devil_cmd_tar := exec.Command("curl", "-L", devil_tar_url, "-o", "package.tar.gz")
	err := devil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	devil_zip_url := "https://downloads.sourceforge.net/project/openil/DevIL/1.8.0/DevIL-1.8.0.zip"
	devil_cmd_zip := exec.Command("curl", "-L", devil_zip_url, "-o", "package.zip")
	err = devil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	devil_bin_url := "https://downloads.sourceforge.net/project/openil/DevIL/1.8.0/DevIL-1.8.0.bin"
	devil_cmd_bin := exec.Command("curl", "-L", devil_bin_url, "-o", "binary.bin")
	err = devil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	devil_src_url := "https://downloads.sourceforge.net/project/openil/DevIL/1.8.0/DevIL-1.8.0.src.tar.gz"
	devil_cmd_src := exec.Command("curl", "-L", devil_src_url, "-o", "source.tar.gz")
	err = devil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	devil_cmd_direct := exec.Command("./binary")
	err = devil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: jasper")
	exec.Command("latte", "install", "jasper").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
}
