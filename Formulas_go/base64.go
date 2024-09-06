package main

// Base64 - Encode and decode base64 files
// Homepage: https://www.fourmilab.ch/webtools/base64/

import (
	"fmt"
	
	"os/exec"
)

func installBase64() {
	// Método 1: Descargar y extraer .tar.gz
	base64_tar_url := "https://www.fourmilab.ch/webtools/base64/base64-1.5.tar.gz"
	base64_cmd_tar := exec.Command("curl", "-L", base64_tar_url, "-o", "package.tar.gz")
	err := base64_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	base64_zip_url := "https://www.fourmilab.ch/webtools/base64/base64-1.5.zip"
	base64_cmd_zip := exec.Command("curl", "-L", base64_zip_url, "-o", "package.zip")
	err = base64_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	base64_bin_url := "https://www.fourmilab.ch/webtools/base64/base64-1.5.bin"
	base64_cmd_bin := exec.Command("curl", "-L", base64_bin_url, "-o", "binary.bin")
	err = base64_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	base64_src_url := "https://www.fourmilab.ch/webtools/base64/base64-1.5.src.tar.gz"
	base64_cmd_src := exec.Command("curl", "-L", base64_src_url, "-o", "source.tar.gz")
	err = base64_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	base64_cmd_direct := exec.Command("./binary")
	err = base64_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
