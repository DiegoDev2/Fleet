package main

// Dfc - Display graphs and colors of file system space/usage
// Homepage: https://github.com/Rolinh/dfc

import (
	"fmt"
	
	"os/exec"
)

func installDfc() {
	// Método 1: Descargar y extraer .tar.gz
	dfc_tar_url := "https://github.com/Rolinh/dfc/releases/download/v3.1.1/dfc-3.1.1.tar.gz"
	dfc_cmd_tar := exec.Command("curl", "-L", dfc_tar_url, "-o", "package.tar.gz")
	err := dfc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dfc_zip_url := "https://github.com/Rolinh/dfc/releases/download/v3.1.1/dfc-3.1.1.zip"
	dfc_cmd_zip := exec.Command("curl", "-L", dfc_zip_url, "-o", "package.zip")
	err = dfc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dfc_bin_url := "https://github.com/Rolinh/dfc/releases/download/v3.1.1/dfc-3.1.1.bin"
	dfc_cmd_bin := exec.Command("curl", "-L", dfc_bin_url, "-o", "binary.bin")
	err = dfc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dfc_src_url := "https://github.com/Rolinh/dfc/releases/download/v3.1.1/dfc-3.1.1.src.tar.gz"
	dfc_cmd_src := exec.Command("curl", "-L", dfc_src_url, "-o", "source.tar.gz")
	err = dfc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dfc_cmd_direct := exec.Command("./binary")
	err = dfc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
