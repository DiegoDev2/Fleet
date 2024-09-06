package main

// Icon - General-purpose programming language
// Homepage: https://www.cs.arizona.edu/icon/

import (
	"fmt"
	
	"os/exec"
)

func installIcon() {
	// Método 1: Descargar y extraer .tar.gz
	icon_tar_url := "https://github.com/gtownsend/icon/archive/refs/tags/v9.5.24a.tar.gz"
	icon_cmd_tar := exec.Command("curl", "-L", icon_tar_url, "-o", "package.tar.gz")
	err := icon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	icon_zip_url := "https://github.com/gtownsend/icon/archive/refs/tags/v9.5.24a.zip"
	icon_cmd_zip := exec.Command("curl", "-L", icon_zip_url, "-o", "package.zip")
	err = icon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	icon_bin_url := "https://github.com/gtownsend/icon/archive/refs/tags/v9.5.24a.bin"
	icon_cmd_bin := exec.Command("curl", "-L", icon_bin_url, "-o", "binary.bin")
	err = icon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	icon_src_url := "https://github.com/gtownsend/icon/archive/refs/tags/v9.5.24a.src.tar.gz"
	icon_cmd_src := exec.Command("curl", "-L", icon_src_url, "-o", "source.tar.gz")
	err = icon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	icon_cmd_direct := exec.Command("./binary")
	err = icon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
