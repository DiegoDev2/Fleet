package main

// Modules - Dynamic modification of a user's environment via modulefiles
// Homepage: https://modules.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installModules() {
	// Método 1: Descargar y extraer .tar.gz
	modules_tar_url := "https://downloads.sourceforge.net/project/modules/Modules/modules-5.4.0/modules-5.4.0.tar.bz2"
	modules_cmd_tar := exec.Command("curl", "-L", modules_tar_url, "-o", "package.tar.gz")
	err := modules_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	modules_zip_url := "https://downloads.sourceforge.net/project/modules/Modules/modules-5.4.0/modules-5.4.0.tar.bz2"
	modules_cmd_zip := exec.Command("curl", "-L", modules_zip_url, "-o", "package.zip")
	err = modules_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	modules_bin_url := "https://downloads.sourceforge.net/project/modules/Modules/modules-5.4.0/modules-5.4.0.tar.bz2"
	modules_cmd_bin := exec.Command("curl", "-L", modules_bin_url, "-o", "binary.bin")
	err = modules_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	modules_src_url := "https://downloads.sourceforge.net/project/modules/Modules/modules-5.4.0/modules-5.4.0.tar.bz2"
	modules_cmd_src := exec.Command("curl", "-L", modules_src_url, "-o", "source.tar.gz")
	err = modules_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	modules_cmd_direct := exec.Command("./binary")
	err = modules_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tcl-tk")
	exec.Command("latte", "install", "tcl-tk").Run()
}
