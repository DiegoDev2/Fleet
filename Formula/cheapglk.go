package main

// Cheapglk - Extremely minimal Glk library
// Homepage: https://www.eblong.com/zarf/glk/

import (
	"fmt"
	
	"os/exec"
)

func installCheapglk() {
	// Método 1: Descargar y extraer .tar.gz
	cheapglk_tar_url := "https://www.eblong.com/zarf/glk/cheapglk-106.tar.gz"
	cheapglk_cmd_tar := exec.Command("curl", "-L", cheapglk_tar_url, "-o", "package.tar.gz")
	err := cheapglk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cheapglk_zip_url := "https://www.eblong.com/zarf/glk/cheapglk-106.zip"
	cheapglk_cmd_zip := exec.Command("curl", "-L", cheapglk_zip_url, "-o", "package.zip")
	err = cheapglk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cheapglk_bin_url := "https://www.eblong.com/zarf/glk/cheapglk-106.bin"
	cheapglk_cmd_bin := exec.Command("curl", "-L", cheapglk_bin_url, "-o", "binary.bin")
	err = cheapglk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cheapglk_src_url := "https://www.eblong.com/zarf/glk/cheapglk-106.src.tar.gz"
	cheapglk_cmd_src := exec.Command("curl", "-L", cheapglk_src_url, "-o", "source.tar.gz")
	err = cheapglk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cheapglk_cmd_direct := exec.Command("./binary")
	err = cheapglk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
