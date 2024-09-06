package main

// Freexl - Library to extract data from Excel .xls files
// Homepage: https://www.gaia-gis.it/fossil/freexl/index

import (
	"fmt"
	
	"os/exec"
)

func installFreexl() {
	// Método 1: Descargar y extraer .tar.gz
	freexl_tar_url := "https://www.gaia-gis.it/gaia-sins/freexl-sources/freexl-2.0.0.tar.gz"
	freexl_cmd_tar := exec.Command("curl", "-L", freexl_tar_url, "-o", "package.tar.gz")
	err := freexl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freexl_zip_url := "https://www.gaia-gis.it/gaia-sins/freexl-sources/freexl-2.0.0.zip"
	freexl_cmd_zip := exec.Command("curl", "-L", freexl_zip_url, "-o", "package.zip")
	err = freexl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freexl_bin_url := "https://www.gaia-gis.it/gaia-sins/freexl-sources/freexl-2.0.0.bin"
	freexl_cmd_bin := exec.Command("curl", "-L", freexl_bin_url, "-o", "binary.bin")
	err = freexl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freexl_src_url := "https://www.gaia-gis.it/gaia-sins/freexl-sources/freexl-2.0.0.src.tar.gz"
	freexl_cmd_src := exec.Command("curl", "-L", freexl_src_url, "-o", "source.tar.gz")
	err = freexl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freexl_cmd_direct := exec.Command("./binary")
	err = freexl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: minizip")
exec.Command("latte", "install", "minizip")
}
