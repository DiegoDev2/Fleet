package main

// LcdfTypetools - Manipulate OpenType and multiple-master fonts
// Homepage: https://www.lcdf.org/type/

import (
	"fmt"
	
	"os/exec"
)

func installLcdfTypetools() {
	// Método 1: Descargar y extraer .tar.gz
	lcdftypetools_tar_url := "https://www.lcdf.org/type/lcdf-typetools-2.110.tar.gz"
	lcdftypetools_cmd_tar := exec.Command("curl", "-L", lcdftypetools_tar_url, "-o", "package.tar.gz")
	err := lcdftypetools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lcdftypetools_zip_url := "https://www.lcdf.org/type/lcdf-typetools-2.110.zip"
	lcdftypetools_cmd_zip := exec.Command("curl", "-L", lcdftypetools_zip_url, "-o", "package.zip")
	err = lcdftypetools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lcdftypetools_bin_url := "https://www.lcdf.org/type/lcdf-typetools-2.110.bin"
	lcdftypetools_cmd_bin := exec.Command("curl", "-L", lcdftypetools_bin_url, "-o", "binary.bin")
	err = lcdftypetools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lcdftypetools_src_url := "https://www.lcdf.org/type/lcdf-typetools-2.110.src.tar.gz"
	lcdftypetools_cmd_src := exec.Command("curl", "-L", lcdftypetools_src_url, "-o", "source.tar.gz")
	err = lcdftypetools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lcdftypetools_cmd_direct := exec.Command("./binary")
	err = lcdftypetools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
