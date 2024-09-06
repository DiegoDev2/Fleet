package main

// T1lib - C library to generate/rasterize bitmaps from Type 1 fonts
// Homepage: https://www.t1lib.org/

import (
	"fmt"
	
	"os/exec"
)

func installT1lib() {
	// Método 1: Descargar y extraer .tar.gz
	t1lib_tar_url := "https://www.ibiblio.org/pub/linux/libs/graphics/t1lib-5.1.2.tar.gz"
	t1lib_cmd_tar := exec.Command("curl", "-L", t1lib_tar_url, "-o", "package.tar.gz")
	err := t1lib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	t1lib_zip_url := "https://www.ibiblio.org/pub/linux/libs/graphics/t1lib-5.1.2.zip"
	t1lib_cmd_zip := exec.Command("curl", "-L", t1lib_zip_url, "-o", "package.zip")
	err = t1lib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	t1lib_bin_url := "https://www.ibiblio.org/pub/linux/libs/graphics/t1lib-5.1.2.bin"
	t1lib_cmd_bin := exec.Command("curl", "-L", t1lib_bin_url, "-o", "binary.bin")
	err = t1lib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	t1lib_src_url := "https://www.ibiblio.org/pub/linux/libs/graphics/t1lib-5.1.2.src.tar.gz"
	t1lib_cmd_src := exec.Command("curl", "-L", t1lib_src_url, "-o", "source.tar.gz")
	err = t1lib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	t1lib_cmd_direct := exec.Command("./binary")
	err = t1lib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
