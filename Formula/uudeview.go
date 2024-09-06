package main

// Uudeview - Smart multi-file multi-part decoder
// Homepage: http://www.fpx.de/fp/Software/UUDeview/

import (
	"fmt"
	
	"os/exec"
)

func installUudeview() {
	// Método 1: Descargar y extraer .tar.gz
	uudeview_tar_url := "http://www.fpx.de/fp/Software/UUDeview/download/uudeview-0.5.20.tar.gz"
	uudeview_cmd_tar := exec.Command("curl", "-L", uudeview_tar_url, "-o", "package.tar.gz")
	err := uudeview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uudeview_zip_url := "http://www.fpx.de/fp/Software/UUDeview/download/uudeview-0.5.20.zip"
	uudeview_cmd_zip := exec.Command("curl", "-L", uudeview_zip_url, "-o", "package.zip")
	err = uudeview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uudeview_bin_url := "http://www.fpx.de/fp/Software/UUDeview/download/uudeview-0.5.20.bin"
	uudeview_cmd_bin := exec.Command("curl", "-L", uudeview_bin_url, "-o", "binary.bin")
	err = uudeview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uudeview_src_url := "http://www.fpx.de/fp/Software/UUDeview/download/uudeview-0.5.20.src.tar.gz"
	uudeview_cmd_src := exec.Command("curl", "-L", uudeview_src_url, "-o", "source.tar.gz")
	err = uudeview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uudeview_cmd_direct := exec.Command("./binary")
	err = uudeview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
