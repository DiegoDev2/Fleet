package main

// Xcinfo - Tool to get information about and install available Xcode versions
// Homepage: https://github.com/xcodereleases/xcinfo

import (
	"fmt"
	
	"os/exec"
)

func installXcinfo() {
	// Método 1: Descargar y extraer .tar.gz
	xcinfo_tar_url := "https://github.com/xcodereleases/xcinfo/archive/refs/tags/1.0.3.tar.gz"
	xcinfo_cmd_tar := exec.Command("curl", "-L", xcinfo_tar_url, "-o", "package.tar.gz")
	err := xcinfo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcinfo_zip_url := "https://github.com/xcodereleases/xcinfo/archive/refs/tags/1.0.3.zip"
	xcinfo_cmd_zip := exec.Command("curl", "-L", xcinfo_zip_url, "-o", "package.zip")
	err = xcinfo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcinfo_bin_url := "https://github.com/xcodereleases/xcinfo/archive/refs/tags/1.0.3.bin"
	xcinfo_cmd_bin := exec.Command("curl", "-L", xcinfo_bin_url, "-o", "binary.bin")
	err = xcinfo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcinfo_src_url := "https://github.com/xcodereleases/xcinfo/archive/refs/tags/1.0.3.src.tar.gz"
	xcinfo_cmd_src := exec.Command("curl", "-L", xcinfo_src_url, "-o", "source.tar.gz")
	err = xcinfo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcinfo_cmd_direct := exec.Command("./binary")
	err = xcinfo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
