package main

// Bvi - Vi-like binary file (hex) editor
// Homepage: https://bvi.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBvi() {
	// Método 1: Descargar y extraer .tar.gz
	bvi_tar_url := "https://downloads.sourceforge.net/project/bvi/bvi/1.4.2/bvi-1.4.2.src.tar.gz"
	bvi_cmd_tar := exec.Command("curl", "-L", bvi_tar_url, "-o", "package.tar.gz")
	err := bvi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bvi_zip_url := "https://downloads.sourceforge.net/project/bvi/bvi/1.4.2/bvi-1.4.2.src.zip"
	bvi_cmd_zip := exec.Command("curl", "-L", bvi_zip_url, "-o", "package.zip")
	err = bvi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bvi_bin_url := "https://downloads.sourceforge.net/project/bvi/bvi/1.4.2/bvi-1.4.2.src.bin"
	bvi_cmd_bin := exec.Command("curl", "-L", bvi_bin_url, "-o", "binary.bin")
	err = bvi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bvi_src_url := "https://downloads.sourceforge.net/project/bvi/bvi/1.4.2/bvi-1.4.2.src.src.tar.gz"
	bvi_cmd_src := exec.Command("curl", "-L", bvi_src_url, "-o", "source.tar.gz")
	err = bvi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bvi_cmd_direct := exec.Command("./binary")
	err = bvi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
