package main

// Tta - Lossless audio codec
// Homepage: https://sourceforge.net/projects/tta/

import (
	"fmt"
	
	"os/exec"
)

func installTta() {
	// Método 1: Descargar y extraer .tar.gz
	tta_tar_url := "https://downloads.sourceforge.net/project/tta/tta/libtta/libtta-2.2.tar.gz"
	tta_cmd_tar := exec.Command("curl", "-L", tta_tar_url, "-o", "package.tar.gz")
	err := tta_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tta_zip_url := "https://downloads.sourceforge.net/project/tta/tta/libtta/libtta-2.2.zip"
	tta_cmd_zip := exec.Command("curl", "-L", tta_zip_url, "-o", "package.zip")
	err = tta_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tta_bin_url := "https://downloads.sourceforge.net/project/tta/tta/libtta/libtta-2.2.bin"
	tta_cmd_bin := exec.Command("curl", "-L", tta_bin_url, "-o", "binary.bin")
	err = tta_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tta_src_url := "https://downloads.sourceforge.net/project/tta/tta/libtta/libtta-2.2.src.tar.gz"
	tta_cmd_src := exec.Command("curl", "-L", tta_src_url, "-o", "source.tar.gz")
	err = tta_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tta_cmd_direct := exec.Command("./binary")
	err = tta_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
