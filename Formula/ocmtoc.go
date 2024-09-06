package main

// Ocmtoc - Mach-O to PE/COFF binary converter
// Homepage: https://github.com/acidanthera/ocmtoc

import (
	"fmt"
	
	"os/exec"
)

func installOcmtoc() {
	// Método 1: Descargar y extraer .tar.gz
	ocmtoc_tar_url := "https://github.com/acidanthera/ocmtoc/archive/refs/tags/1.0.3.tar.gz"
	ocmtoc_cmd_tar := exec.Command("curl", "-L", ocmtoc_tar_url, "-o", "package.tar.gz")
	err := ocmtoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocmtoc_zip_url := "https://github.com/acidanthera/ocmtoc/archive/refs/tags/1.0.3.zip"
	ocmtoc_cmd_zip := exec.Command("curl", "-L", ocmtoc_zip_url, "-o", "package.zip")
	err = ocmtoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocmtoc_bin_url := "https://github.com/acidanthera/ocmtoc/archive/refs/tags/1.0.3.bin"
	ocmtoc_cmd_bin := exec.Command("curl", "-L", ocmtoc_bin_url, "-o", "binary.bin")
	err = ocmtoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocmtoc_src_url := "https://github.com/acidanthera/ocmtoc/archive/refs/tags/1.0.3.src.tar.gz"
	ocmtoc_cmd_src := exec.Command("curl", "-L", ocmtoc_src_url, "-o", "source.tar.gz")
	err = ocmtoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocmtoc_cmd_direct := exec.Command("./binary")
	err = ocmtoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
