package main

// Libb64 - Base64 encoding/decoding library
// Homepage: https://libb64.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibb64() {
	// Método 1: Descargar y extraer .tar.gz
	libb64_tar_url := "https://downloads.sourceforge.net/project/libb64/libb64/libb64/libb64-1.2.1.zip"
	libb64_cmd_tar := exec.Command("curl", "-L", libb64_tar_url, "-o", "package.tar.gz")
	err := libb64_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libb64_zip_url := "https://downloads.sourceforge.net/project/libb64/libb64/libb64/libb64-1.2.1.zip"
	libb64_cmd_zip := exec.Command("curl", "-L", libb64_zip_url, "-o", "package.zip")
	err = libb64_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libb64_bin_url := "https://downloads.sourceforge.net/project/libb64/libb64/libb64/libb64-1.2.1.zip"
	libb64_cmd_bin := exec.Command("curl", "-L", libb64_bin_url, "-o", "binary.bin")
	err = libb64_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libb64_src_url := "https://downloads.sourceforge.net/project/libb64/libb64/libb64/libb64-1.2.1.zip"
	libb64_cmd_src := exec.Command("curl", "-L", libb64_src_url, "-o", "source.tar.gz")
	err = libb64_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libb64_cmd_direct := exec.Command("./binary")
	err = libb64_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
