package main

// Ren - Rename multiple files in a directory
// Homepage: https://pdb.finkproject.org/pdb/package.php/ren

import (
	"fmt"
	
	"os/exec"
)

func installRen() {
	// Método 1: Descargar y extraer .tar.gz
	ren_tar_url := "https://www.ibiblio.org/pub/Linux/utils/file/ren-1.0.tar.gz"
	ren_cmd_tar := exec.Command("curl", "-L", ren_tar_url, "-o", "package.tar.gz")
	err := ren_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ren_zip_url := "https://www.ibiblio.org/pub/Linux/utils/file/ren-1.0.zip"
	ren_cmd_zip := exec.Command("curl", "-L", ren_zip_url, "-o", "package.zip")
	err = ren_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ren_bin_url := "https://www.ibiblio.org/pub/Linux/utils/file/ren-1.0.bin"
	ren_cmd_bin := exec.Command("curl", "-L", ren_bin_url, "-o", "binary.bin")
	err = ren_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ren_src_url := "https://www.ibiblio.org/pub/Linux/utils/file/ren-1.0.src.tar.gz"
	ren_cmd_src := exec.Command("curl", "-L", ren_src_url, "-o", "source.tar.gz")
	err = ren_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ren_cmd_direct := exec.Command("./binary")
	err = ren_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
