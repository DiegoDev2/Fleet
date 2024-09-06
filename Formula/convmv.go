package main

// Convmv - Filename encoding conversion tool
// Homepage: https://www.j3e.de/linux/convmv/

import (
	"fmt"
	
	"os/exec"
)

func installConvmv() {
	// Método 1: Descargar y extraer .tar.gz
	convmv_tar_url := "https://www.j3e.de/linux/convmv/convmv-2.05.tar.gz"
	convmv_cmd_tar := exec.Command("curl", "-L", convmv_tar_url, "-o", "package.tar.gz")
	err := convmv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	convmv_zip_url := "https://www.j3e.de/linux/convmv/convmv-2.05.zip"
	convmv_cmd_zip := exec.Command("curl", "-L", convmv_zip_url, "-o", "package.zip")
	err = convmv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	convmv_bin_url := "https://www.j3e.de/linux/convmv/convmv-2.05.bin"
	convmv_cmd_bin := exec.Command("curl", "-L", convmv_bin_url, "-o", "binary.bin")
	err = convmv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	convmv_src_url := "https://www.j3e.de/linux/convmv/convmv-2.05.src.tar.gz"
	convmv_cmd_src := exec.Command("curl", "-L", convmv_src_url, "-o", "source.tar.gz")
	err = convmv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	convmv_cmd_direct := exec.Command("./binary")
	err = convmv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
