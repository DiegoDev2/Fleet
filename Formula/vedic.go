package main

// Vedic - Simple Sanskrit programming language
// Homepage: https://vedic-lang.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installVedic() {
	// Método 1: Descargar y extraer .tar.gz
	vedic_tar_url := "https://github.com/vedic-lang/vedic/archive/refs/tags/v2.0.6.tar.gz"
	vedic_cmd_tar := exec.Command("curl", "-L", vedic_tar_url, "-o", "package.tar.gz")
	err := vedic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vedic_zip_url := "https://github.com/vedic-lang/vedic/archive/refs/tags/v2.0.6.zip"
	vedic_cmd_zip := exec.Command("curl", "-L", vedic_zip_url, "-o", "package.zip")
	err = vedic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vedic_bin_url := "https://github.com/vedic-lang/vedic/archive/refs/tags/v2.0.6.bin"
	vedic_cmd_bin := exec.Command("curl", "-L", vedic_bin_url, "-o", "binary.bin")
	err = vedic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vedic_src_url := "https://github.com/vedic-lang/vedic/archive/refs/tags/v2.0.6.src.tar.gz"
	vedic_cmd_src := exec.Command("curl", "-L", vedic_src_url, "-o", "source.tar.gz")
	err = vedic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vedic_cmd_direct := exec.Command("./binary")
	err = vedic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
