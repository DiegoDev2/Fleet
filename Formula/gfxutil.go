package main

// Gfxutil - Device Properties conversion tool
// Homepage: https://github.com/acidanthera/gfxutil

import (
	"fmt"
	
	"os/exec"
)

func installGfxutil() {
	// Método 1: Descargar y extraer .tar.gz
	gfxutil_tar_url := "https://github.com/acidanthera/gfxutil/archive/refs/tags/1.84b.tar.gz"
	gfxutil_cmd_tar := exec.Command("curl", "-L", gfxutil_tar_url, "-o", "package.tar.gz")
	err := gfxutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gfxutil_zip_url := "https://github.com/acidanthera/gfxutil/archive/refs/tags/1.84b.zip"
	gfxutil_cmd_zip := exec.Command("curl", "-L", gfxutil_zip_url, "-o", "package.zip")
	err = gfxutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gfxutil_bin_url := "https://github.com/acidanthera/gfxutil/archive/refs/tags/1.84b.bin"
	gfxutil_cmd_bin := exec.Command("curl", "-L", gfxutil_bin_url, "-o", "binary.bin")
	err = gfxutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gfxutil_src_url := "https://github.com/acidanthera/gfxutil/archive/refs/tags/1.84b.src.tar.gz"
	gfxutil_cmd_src := exec.Command("curl", "-L", gfxutil_src_url, "-o", "source.tar.gz")
	err = gfxutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gfxutil_cmd_direct := exec.Command("./binary")
	err = gfxutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
