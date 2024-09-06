package main

// Yamdi - Add metadata to Flash video
// Homepage: https://yamdi.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installYamdi() {
	// Método 1: Descargar y extraer .tar.gz
	yamdi_tar_url := "https://downloads.sourceforge.net/project/yamdi/yamdi/1.9/yamdi-1.9.tar.gz"
	yamdi_cmd_tar := exec.Command("curl", "-L", yamdi_tar_url, "-o", "package.tar.gz")
	err := yamdi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yamdi_zip_url := "https://downloads.sourceforge.net/project/yamdi/yamdi/1.9/yamdi-1.9.zip"
	yamdi_cmd_zip := exec.Command("curl", "-L", yamdi_zip_url, "-o", "package.zip")
	err = yamdi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yamdi_bin_url := "https://downloads.sourceforge.net/project/yamdi/yamdi/1.9/yamdi-1.9.bin"
	yamdi_cmd_bin := exec.Command("curl", "-L", yamdi_bin_url, "-o", "binary.bin")
	err = yamdi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yamdi_src_url := "https://downloads.sourceforge.net/project/yamdi/yamdi/1.9/yamdi-1.9.src.tar.gz"
	yamdi_cmd_src := exec.Command("curl", "-L", yamdi_src_url, "-o", "source.tar.gz")
	err = yamdi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yamdi_cmd_direct := exec.Command("./binary")
	err = yamdi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
