package main

// Wiggle - Program for applying patches with conflicting changes
// Homepage: https://github.com/neilbrown/wiggle

import (
	"fmt"
	
	"os/exec"
)

func installWiggle() {
	// Método 1: Descargar y extraer .tar.gz
	wiggle_tar_url := "https://github.com/neilbrown/wiggle/archive/refs/tags/v1.3.tar.gz"
	wiggle_cmd_tar := exec.Command("curl", "-L", wiggle_tar_url, "-o", "package.tar.gz")
	err := wiggle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wiggle_zip_url := "https://github.com/neilbrown/wiggle/archive/refs/tags/v1.3.zip"
	wiggle_cmd_zip := exec.Command("curl", "-L", wiggle_zip_url, "-o", "package.zip")
	err = wiggle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wiggle_bin_url := "https://github.com/neilbrown/wiggle/archive/refs/tags/v1.3.bin"
	wiggle_cmd_bin := exec.Command("curl", "-L", wiggle_bin_url, "-o", "binary.bin")
	err = wiggle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wiggle_src_url := "https://github.com/neilbrown/wiggle/archive/refs/tags/v1.3.src.tar.gz"
	wiggle_cmd_src := exec.Command("curl", "-L", wiggle_src_url, "-o", "source.tar.gz")
	err = wiggle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wiggle_cmd_direct := exec.Command("./binary")
	err = wiggle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
