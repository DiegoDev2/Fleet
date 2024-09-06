package main

// Zork - Dungeon modified from FORTRAN to C
// Homepage: https://github.com/devshane/zork

import (
	"fmt"
	
	"os/exec"
)

func installZork() {
	// Método 1: Descargar y extraer .tar.gz
	zork_tar_url := "https://github.com/devshane/zork/archive/refs/tags/v1.0.3.tar.gz"
	zork_cmd_tar := exec.Command("curl", "-L", zork_tar_url, "-o", "package.tar.gz")
	err := zork_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zork_zip_url := "https://github.com/devshane/zork/archive/refs/tags/v1.0.3.zip"
	zork_cmd_zip := exec.Command("curl", "-L", zork_zip_url, "-o", "package.zip")
	err = zork_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zork_bin_url := "https://github.com/devshane/zork/archive/refs/tags/v1.0.3.bin"
	zork_cmd_bin := exec.Command("curl", "-L", zork_bin_url, "-o", "binary.bin")
	err = zork_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zork_src_url := "https://github.com/devshane/zork/archive/refs/tags/v1.0.3.src.tar.gz"
	zork_cmd_src := exec.Command("curl", "-L", zork_src_url, "-o", "source.tar.gz")
	err = zork_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zork_cmd_direct := exec.Command("./binary")
	err = zork_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
