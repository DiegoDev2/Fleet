package main

// Openlibm - High quality, portable, open source libm implementation
// Homepage: https://openlibm.org

import (
	"fmt"
	
	"os/exec"
)

func installOpenlibm() {
	// Método 1: Descargar y extraer .tar.gz
	openlibm_tar_url := "https://github.com/JuliaMath/openlibm/archive/refs/tags/v0.8.3.tar.gz"
	openlibm_cmd_tar := exec.Command("curl", "-L", openlibm_tar_url, "-o", "package.tar.gz")
	err := openlibm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openlibm_zip_url := "https://github.com/JuliaMath/openlibm/archive/refs/tags/v0.8.3.zip"
	openlibm_cmd_zip := exec.Command("curl", "-L", openlibm_zip_url, "-o", "package.zip")
	err = openlibm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openlibm_bin_url := "https://github.com/JuliaMath/openlibm/archive/refs/tags/v0.8.3.bin"
	openlibm_cmd_bin := exec.Command("curl", "-L", openlibm_bin_url, "-o", "binary.bin")
	err = openlibm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openlibm_src_url := "https://github.com/JuliaMath/openlibm/archive/refs/tags/v0.8.3.src.tar.gz"
	openlibm_cmd_src := exec.Command("curl", "-L", openlibm_src_url, "-o", "source.tar.gz")
	err = openlibm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openlibm_cmd_direct := exec.Command("./binary")
	err = openlibm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
