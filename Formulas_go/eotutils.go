package main

// EotUtils - Tools to convert fonts from OTF/TTF to EOT format
// Homepage: https://www.w3.org/Tools/eot-utils/

import (
	"fmt"
	
	"os/exec"
)

func installEotUtils() {
	// Método 1: Descargar y extraer .tar.gz
	eotutils_tar_url := "https://www.w3.org/Tools/eot-utils/eot-utilities-1.1.tar.gz"
	eotutils_cmd_tar := exec.Command("curl", "-L", eotutils_tar_url, "-o", "package.tar.gz")
	err := eotutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eotutils_zip_url := "https://www.w3.org/Tools/eot-utils/eot-utilities-1.1.zip"
	eotutils_cmd_zip := exec.Command("curl", "-L", eotutils_zip_url, "-o", "package.zip")
	err = eotutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eotutils_bin_url := "https://www.w3.org/Tools/eot-utils/eot-utilities-1.1.bin"
	eotutils_cmd_bin := exec.Command("curl", "-L", eotutils_bin_url, "-o", "binary.bin")
	err = eotutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eotutils_src_url := "https://www.w3.org/Tools/eot-utils/eot-utilities-1.1.src.tar.gz"
	eotutils_cmd_src := exec.Command("curl", "-L", eotutils_src_url, "-o", "source.tar.gz")
	err = eotutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eotutils_cmd_direct := exec.Command("./binary")
	err = eotutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
