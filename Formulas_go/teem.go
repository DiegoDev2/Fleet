package main

// Teem - Libraries for scientific raster data
// Homepage: https://teem.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installTeem() {
	// Método 1: Descargar y extraer .tar.gz
	teem_tar_url := "https://downloads.sourceforge.net/project/teem/teem/1.11.0/teem-1.11.0-src.tar.gz"
	teem_cmd_tar := exec.Command("curl", "-L", teem_tar_url, "-o", "package.tar.gz")
	err := teem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	teem_zip_url := "https://downloads.sourceforge.net/project/teem/teem/1.11.0/teem-1.11.0-src.zip"
	teem_cmd_zip := exec.Command("curl", "-L", teem_zip_url, "-o", "package.zip")
	err = teem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	teem_bin_url := "https://downloads.sourceforge.net/project/teem/teem/1.11.0/teem-1.11.0-src.bin"
	teem_cmd_bin := exec.Command("curl", "-L", teem_bin_url, "-o", "binary.bin")
	err = teem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	teem_src_url := "https://downloads.sourceforge.net/project/teem/teem/1.11.0/teem-1.11.0-src.src.tar.gz"
	teem_cmd_src := exec.Command("curl", "-L", teem_src_url, "-o", "source.tar.gz")
	err = teem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	teem_cmd_direct := exec.Command("./binary")
	err = teem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
