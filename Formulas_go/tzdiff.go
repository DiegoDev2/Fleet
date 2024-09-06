package main

// Tzdiff - Displays Timezone differences with localtime in CLI (shell script)
// Homepage: https://github.com/belgianbeer/tzdiff

import (
	"fmt"
	
	"os/exec"
)

func installTzdiff() {
	// Método 1: Descargar y extraer .tar.gz
	tzdiff_tar_url := "https://github.com/belgianbeer/tzdiff/archive/refs/tags/1.2.1.tar.gz"
	tzdiff_cmd_tar := exec.Command("curl", "-L", tzdiff_tar_url, "-o", "package.tar.gz")
	err := tzdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tzdiff_zip_url := "https://github.com/belgianbeer/tzdiff/archive/refs/tags/1.2.1.zip"
	tzdiff_cmd_zip := exec.Command("curl", "-L", tzdiff_zip_url, "-o", "package.zip")
	err = tzdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tzdiff_bin_url := "https://github.com/belgianbeer/tzdiff/archive/refs/tags/1.2.1.bin"
	tzdiff_cmd_bin := exec.Command("curl", "-L", tzdiff_bin_url, "-o", "binary.bin")
	err = tzdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tzdiff_src_url := "https://github.com/belgianbeer/tzdiff/archive/refs/tags/1.2.1.src.tar.gz"
	tzdiff_cmd_src := exec.Command("curl", "-L", tzdiff_src_url, "-o", "source.tar.gz")
	err = tzdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tzdiff_cmd_direct := exec.Command("./binary")
	err = tzdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
