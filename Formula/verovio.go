package main

// Verovio - Command-line MEI music notation engraver
// Homepage: https://www.verovio.org

import (
	"fmt"
	
	"os/exec"
)

func installVerovio() {
	// Método 1: Descargar y extraer .tar.gz
	verovio_tar_url := "https://github.com/rism-digital/verovio/archive/refs/tags/version-4.2.1.tar.gz"
	verovio_cmd_tar := exec.Command("curl", "-L", verovio_tar_url, "-o", "package.tar.gz")
	err := verovio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	verovio_zip_url := "https://github.com/rism-digital/verovio/archive/refs/tags/version-4.2.1.zip"
	verovio_cmd_zip := exec.Command("curl", "-L", verovio_zip_url, "-o", "package.zip")
	err = verovio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	verovio_bin_url := "https://github.com/rism-digital/verovio/archive/refs/tags/version-4.2.1.bin"
	verovio_cmd_bin := exec.Command("curl", "-L", verovio_bin_url, "-o", "binary.bin")
	err = verovio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	verovio_src_url := "https://github.com/rism-digital/verovio/archive/refs/tags/version-4.2.1.src.tar.gz"
	verovio_cmd_src := exec.Command("curl", "-L", verovio_src_url, "-o", "source.tar.gz")
	err = verovio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	verovio_cmd_direct := exec.Command("./binary")
	err = verovio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
