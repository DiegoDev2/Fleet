package main

// Sword - Cross-platform tools to write Bible software
// Homepage: https://www.crosswire.org/sword/index.jsp

import (
	"fmt"
	
	"os/exec"
)

func installSword() {
	// Método 1: Descargar y extraer .tar.gz
	sword_tar_url := "https://www.crosswire.org/ftpmirror/pub/sword/source/v1.9/sword-1.9.0.tar.gz"
	sword_cmd_tar := exec.Command("curl", "-L", sword_tar_url, "-o", "package.tar.gz")
	err := sword_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sword_zip_url := "https://www.crosswire.org/ftpmirror/pub/sword/source/v1.9/sword-1.9.0.zip"
	sword_cmd_zip := exec.Command("curl", "-L", sword_zip_url, "-o", "package.zip")
	err = sword_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sword_bin_url := "https://www.crosswire.org/ftpmirror/pub/sword/source/v1.9/sword-1.9.0.bin"
	sword_cmd_bin := exec.Command("curl", "-L", sword_bin_url, "-o", "binary.bin")
	err = sword_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sword_src_url := "https://www.crosswire.org/ftpmirror/pub/sword/source/v1.9/sword-1.9.0.src.tar.gz"
	sword_cmd_src := exec.Command("curl", "-L", sword_src_url, "-o", "source.tar.gz")
	err = sword_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sword_cmd_direct := exec.Command("./binary")
	err = sword_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
