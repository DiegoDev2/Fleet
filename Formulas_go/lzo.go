package main

// Lzo - Real-time data compression library
// Homepage: https://www.oberhumer.com/opensource/lzo/

import (
	"fmt"
	
	"os/exec"
)

func installLzo() {
	// Método 1: Descargar y extraer .tar.gz
	lzo_tar_url := "https://www.oberhumer.com/opensource/lzo/download/lzo-2.10.tar.gz"
	lzo_cmd_tar := exec.Command("curl", "-L", lzo_tar_url, "-o", "package.tar.gz")
	err := lzo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lzo_zip_url := "https://www.oberhumer.com/opensource/lzo/download/lzo-2.10.zip"
	lzo_cmd_zip := exec.Command("curl", "-L", lzo_zip_url, "-o", "package.zip")
	err = lzo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lzo_bin_url := "https://www.oberhumer.com/opensource/lzo/download/lzo-2.10.bin"
	lzo_cmd_bin := exec.Command("curl", "-L", lzo_bin_url, "-o", "binary.bin")
	err = lzo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lzo_src_url := "https://www.oberhumer.com/opensource/lzo/download/lzo-2.10.src.tar.gz"
	lzo_cmd_src := exec.Command("curl", "-L", lzo_src_url, "-o", "source.tar.gz")
	err = lzo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lzo_cmd_direct := exec.Command("./binary")
	err = lzo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
