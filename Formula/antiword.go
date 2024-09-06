package main

// Antiword - Utility to read Word (.doc) files
// Homepage: https://web.archive.org/web/20221207132720/http://www.winfield.demon.nl/

import (
	"fmt"
	
	"os/exec"
)

func installAntiword() {
	// Método 1: Descargar y extraer .tar.gz
	antiword_tar_url := "https://web.archive.org/web/20221207132720/http://www.winfield.demon.nl/linux/antiword-0.37.tar.gz"
	antiword_cmd_tar := exec.Command("curl", "-L", antiword_tar_url, "-o", "package.tar.gz")
	err := antiword_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	antiword_zip_url := "https://web.archive.org/web/20221207132720/http://www.winfield.demon.nl/linux/antiword-0.37.zip"
	antiword_cmd_zip := exec.Command("curl", "-L", antiword_zip_url, "-o", "package.zip")
	err = antiword_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	antiword_bin_url := "https://web.archive.org/web/20221207132720/http://www.winfield.demon.nl/linux/antiword-0.37.bin"
	antiword_cmd_bin := exec.Command("curl", "-L", antiword_bin_url, "-o", "binary.bin")
	err = antiword_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	antiword_src_url := "https://web.archive.org/web/20221207132720/http://www.winfield.demon.nl/linux/antiword-0.37.src.tar.gz"
	antiword_cmd_src := exec.Command("curl", "-L", antiword_src_url, "-o", "source.tar.gz")
	err = antiword_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	antiword_cmd_direct := exec.Command("./binary")
	err = antiword_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
