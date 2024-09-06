package main

// Ucon64 - ROM backup tool and emulator's Swiss Army knife program
// Homepage: https://ucon64.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installUcon64() {
	// Método 1: Descargar y extraer .tar.gz
	ucon64_tar_url := "https://downloads.sourceforge.net/project/ucon64/ucon64/ucon64-2.2.2/ucon64-2.2.2-src.tar.gz"
	ucon64_cmd_tar := exec.Command("curl", "-L", ucon64_tar_url, "-o", "package.tar.gz")
	err := ucon64_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ucon64_zip_url := "https://downloads.sourceforge.net/project/ucon64/ucon64/ucon64-2.2.2/ucon64-2.2.2-src.zip"
	ucon64_cmd_zip := exec.Command("curl", "-L", ucon64_zip_url, "-o", "package.zip")
	err = ucon64_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ucon64_bin_url := "https://downloads.sourceforge.net/project/ucon64/ucon64/ucon64-2.2.2/ucon64-2.2.2-src.bin"
	ucon64_cmd_bin := exec.Command("curl", "-L", ucon64_bin_url, "-o", "binary.bin")
	err = ucon64_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ucon64_src_url := "https://downloads.sourceforge.net/project/ucon64/ucon64/ucon64-2.2.2/ucon64-2.2.2-src.src.tar.gz"
	ucon64_cmd_src := exec.Command("curl", "-L", ucon64_src_url, "-o", "source.tar.gz")
	err = ucon64_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ucon64_cmd_direct := exec.Command("./binary")
	err = ucon64_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
