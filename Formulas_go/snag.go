package main

// Snag - Automatic build tool for all your needs
// Homepage: https://github.com/Tonkpils/snag

import (
	"fmt"
	
	"os/exec"
)

func installSnag() {
	// Método 1: Descargar y extraer .tar.gz
	snag_tar_url := "https://github.com/Tonkpils/snag/archive/refs/tags/v1.2.0.tar.gz"
	snag_cmd_tar := exec.Command("curl", "-L", snag_tar_url, "-o", "package.tar.gz")
	err := snag_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snag_zip_url := "https://github.com/Tonkpils/snag/archive/refs/tags/v1.2.0.zip"
	snag_cmd_zip := exec.Command("curl", "-L", snag_zip_url, "-o", "package.zip")
	err = snag_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snag_bin_url := "https://github.com/Tonkpils/snag/archive/refs/tags/v1.2.0.bin"
	snag_cmd_bin := exec.Command("curl", "-L", snag_bin_url, "-o", "binary.bin")
	err = snag_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snag_src_url := "https://github.com/Tonkpils/snag/archive/refs/tags/v1.2.0.src.tar.gz"
	snag_cmd_src := exec.Command("curl", "-L", snag_src_url, "-o", "source.tar.gz")
	err = snag_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snag_cmd_direct := exec.Command("./binary")
	err = snag_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
